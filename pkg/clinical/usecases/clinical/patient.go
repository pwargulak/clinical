package clinical

import (
	"context"
	"fmt"
	"sync"

	linq "github.com/ahmetb/go-linq/v3"
	"github.com/google/uuid"
	"github.com/savannahghi/clinical/pkg/clinical/application/common"
	"github.com/savannahghi/clinical/pkg/clinical/application/common/helpers"
	"github.com/savannahghi/clinical/pkg/clinical/application/dto"
	"github.com/savannahghi/clinical/pkg/clinical/application/utils"
	"github.com/savannahghi/clinical/pkg/clinical/domain"
	"github.com/savannahghi/clinical/pkg/clinical/infrastructure"
	"github.com/savannahghi/converterandformatter"
	"github.com/savannahghi/scalarutils"
	log "github.com/sirupsen/logrus"
)

// UseCasesClinicalImpl represents the patient usecase implementation
type UseCasesClinicalImpl struct {
	infrastructure infrastructure.Infrastructure
}

// NewUseCasesClinicalImpl initializes new Clinical/Patient implementation
func NewUseCasesClinicalImpl(infra infrastructure.Infrastructure) *UseCasesClinicalImpl {
	return &UseCasesClinicalImpl{
		infrastructure: infra,
	}
}

// PatientTimeline return's the patient's historical timeline sorted in descending order i.e when it was first recorded
// The timeline consists of Allergies, Observations, Medication statement and Test results
func (c *UseCasesClinicalImpl) PatientTimeline(ctx context.Context, patientID string) ([]map[string]interface{}, error) {
	_, err := uuid.Parse(patientID)
	if err != nil {
		return nil, fmt.Errorf("invalid patient id: %s", patientID)
	}

	timeline := []map[string]interface{}{}
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	patientFilterParams := map[string]interface{}{
		"patient": fmt.Sprintf("Patient/%v", patientID),
	}

	// timelineResourceFunc is a go routine that fetches particular FHIR resource and
	// adds it to the timeline
	type timelineResourceFunc func(wg *sync.WaitGroup, mut *sync.Mutex)

	allergyIntoleranceResourceFunc := func(wg *sync.WaitGroup, mut *sync.Mutex) {
		defer wg.Done()

		conn, err := c.infrastructure.FHIR.SearchFHIRAllergyIntolerance(ctx, patientFilterParams)
		if err != nil {
			utils.ReportErrorToSentry(err)
			log.Errorf("AllergyIntolerance search error: %v", err)

			return
		}

		for _, edge := range conn.Edges {
			if edge.Node == nil {
				continue
			}

			rMap, err := converterandformatter.StructToMap(edge.Node)
			if err != nil {
				utils.ReportErrorToSentry(err)
				log.Errorf("AllergyIntolerance edge struct to map error: %v", err)

				return
			}

			if rMap == nil {
				continue
			}

			rMap["resourceType"] = "AllergyIntolerance"
			rMap["timelineDate"] = rMap["recordedDate"]

			mut.Lock()
			timeline = append(timeline, rMap)
			mut.Unlock()
		}
	}

	observationResourceFunc := func(wg *sync.WaitGroup, mut *sync.Mutex) {
		defer wg.Done()

		conn, err := c.infrastructure.FHIR.SearchFHIRObservation(ctx, patientFilterParams)
		if err != nil {
			utils.ReportErrorToSentry(err)
			log.Errorf("Observation search error: %v", err)

			return
		}

		for _, edge := range conn.Edges {
			if edge.Node == nil {
				continue
			}

			rMap, err := converterandformatter.StructToMap(edge.Node)
			if err != nil {
				utils.ReportErrorToSentry(err)
				log.Errorf("Observation edge struct to map error: %v", err)

				return
			}

			if rMap == nil {
				continue
			}

			rMap["resourceType"] = "Observation"
			rMap["timelineDate"] = rMap["effectiveInstant"]

			mut.Lock()
			timeline = append(timeline, rMap)
			mut.Unlock()
		}
	}

	medicationStatementResourceFunc := func(wg *sync.WaitGroup, mut *sync.Mutex) {
		defer wg.Done()

		conn, err := c.infrastructure.FHIR.SearchFHIRMedicationStatement(ctx, patientFilterParams)
		if err != nil {
			utils.ReportErrorToSentry(err)
			log.Errorf("MedicationStatement search error: %v", err)

			return
		}

		for _, edge := range conn.Edges {
			if edge.Node == nil {
				continue
			}

			rMap, err := converterandformatter.StructToMap(edge.Node)
			if err != nil {
				utils.ReportErrorToSentry(err)
				log.Errorf("MedicationStatement edge struct to map error: %v", err)

				return
			}

			if rMap == nil {
				continue
			}

			rMap["resourceType"] = "MedicationStatement"
			rMap["timelineDate"] = rMap["effectiveDateTime"]

			mut.Lock()
			timeline = append(timeline, rMap)
			mut.Unlock()
		}
	}

	resources := []timelineResourceFunc{
		allergyIntoleranceResourceFunc,
		observationResourceFunc,
		medicationStatementResourceFunc,
	}

	for _, resource := range resources {
		wg.Add(1)

		go resource(wg, mut)
	}

	wg.Wait()

	return timeline, nil
}

// GetMedicalData returns a limited subset of specific medical data that for a specific patient
// These include: Allergies, Viral Load, Body Mass Index, Weight, CD4 Count using their respective OCL CIEL Terminology
// For each category the latest three records are fetched
func (c *UseCasesClinicalImpl) GetMedicalData(ctx context.Context, patientID string) (*domain.MedicalData, error) {
	data := &domain.MedicalData{}

	filterParams := map[string]interface{}{
		"patient": fmt.Sprintf("Patient/%v", patientID),
		"_count":  common.MedicalDataCount,
		"_sort":   "-date",
	}

	fields := []string{
		"Regimen",
		"AllergyIntolerance",
		"Weight",
		"BMI",
		"ViralLoad",
		"CD4Count",
	}

	for _, field := range fields {
		switch field {
		case "Regimen":
			conn, err := c.infrastructure.FHIR.SearchFHIRMedicationStatement(ctx, filterParams)
			if err != nil {
				utils.ReportErrorToSentry(err)
				return nil, fmt.Errorf("%s search error: %w", field, err)
			}

			for _, edge := range conn.Edges {
				if edge.Node == nil {
					continue
				}

				data.Regimen = append(data.Regimen, edge.Node)
			}
		case "AllergyIntolerance":
			conn, err := c.infrastructure.FHIR.SearchFHIRAllergyIntolerance(ctx, filterParams)
			if err != nil {
				utils.ReportErrorToSentry(err)
				return nil, fmt.Errorf("%s search error: %w", field, err)
			}

			for _, edge := range conn.Edges {
				if edge.Node == nil {
					continue
				}

				data.Allergies = append(data.Allergies, edge.Node)
			}

		case "Weight":
			filterParams["code"] = common.WeightCIELTerminologyCode

			conn, err := c.infrastructure.FHIR.SearchFHIRObservation(ctx, filterParams)
			if err != nil {
				utils.ReportErrorToSentry(err)
				return nil, fmt.Errorf("%s search error: %w", field, err)
			}

			for _, edge := range conn.Edges {
				if edge.Node == nil {
					continue
				}

				data.Weight = append(data.Weight, edge.Node)
			}

		case "BMI":
			filterParams["code"] = common.BMICIELTerminologyCode

			conn, err := c.infrastructure.FHIR.SearchFHIRObservation(ctx, filterParams)
			if err != nil {
				utils.ReportErrorToSentry(err)
				return nil, fmt.Errorf("%s search error: %w", field, err)
			}

			for _, edge := range conn.Edges {
				if edge.Node == nil {
					continue
				}

				data.BMI = append(data.BMI, edge.Node)
			}

		case "ViralLoad":
			filterParams["code"] = common.ViralLoadCIELTerminologyCode

			conn, err := c.infrastructure.FHIR.SearchFHIRObservation(ctx, filterParams)
			if err != nil {
				utils.ReportErrorToSentry(err)
				return nil, fmt.Errorf("%s search error: %w", field, err)
			}

			for _, edge := range conn.Edges {
				if edge.Node == nil {
					continue
				}

				data.ViralLoad = append(data.ViralLoad, edge.Node)
			}

		case "CD4Count":
			filterParams["code"] = common.CD4CountCIELTerminologyCode

			conn, err := c.infrastructure.FHIR.SearchFHIRObservation(ctx, filterParams)
			if err != nil {
				utils.ReportErrorToSentry(err)
				return nil, fmt.Errorf("%s search error: %w", field, err)
			}

			for _, edge := range conn.Edges {
				if edge.Node == nil {
					continue
				}

				data.CD4Count = append(data.CD4Count, edge.Node)
			}
		}
	}

	return data, nil
}

// PatientHealthTimeline return's the patient's historical timeline sorted in descending order i.e when it was first recorded
// The timeline consists of Allergies, Observations, Medication statement and Test results
func (c *UseCasesClinicalImpl) PatientHealthTimeline(ctx context.Context, input domain.HealthTimelineInput) (*domain.HealthTimeline, error) {
	records, err := c.PatientTimeline(ctx, input.PatientID)
	if err != nil {
		utils.ReportErrorToSentry(err)
		return nil, fmt.Errorf("cannot retrieve patient timeline error: %w", err)
	}

	data := &domain.HealthTimeline{}
	timeline := []map[string]interface{}{}

	sortFunc := func(i, j interface{}) bool {
		itemI := i.(map[string]interface{})

		dateStringI, ok := itemI["timelineDate"].(string)
		if !ok {
			return false
		}

		timeI := helpers.ParseDate(dateStringI)

		itemJ := j.(map[string]interface{})

		dateStringJ, ok := itemJ["timelineDate"].(string)
		if !ok {
			return false
		}

		timeJ := helpers.ParseDate(dateStringJ)

		return timeI.After(timeJ)
	}

	linq.From(records).Sort(sortFunc).Skip(input.Offset).Take(input.Limit).ToSlice(&timeline)

	data.TotalCount = len(records)
	data.Timeline = timeline

	return data, nil
}

// CreateFHIROrganization creates a FHIROrganization instance
func (c *UseCasesClinicalImpl) CreateFHIROrganization(ctx context.Context, input domain.FHIROrganizationInput) (*domain.FHIROrganizationRelayPayload, error) {
	organizationRelayPayload, err := c.infrastructure.FHIR.CreateFHIROrganization(ctx, input)
	if err != nil {
		utils.ReportErrorToSentry(err)
		return nil, err
	}

	return organizationRelayPayload, nil
}

// CreatePatient is used to create a new patient in the repository
func (c *UseCasesClinicalImpl) CreatePatient(ctx context.Context, input dto.PatientInput) (*dto.Patient, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	var identifiers []*domain.FHIRIdentifierInput
	for _, idDocument := range input.IdentificationDocuments {
		system := scalarutils.URI("http://terminology.hl7.org/CodeSystem/v2-0203")
		identifierSystem := scalarutils.URI("urn:oid:1.2.36.146.595.217.0.1")

		var codingInput []*domain.FHIRCodingInput
		codingMap := map[string]string{
			"passport":    "PPN",
			"national_id": "NI",
			"ccc_number":  "MR",
		}
		for _, identificationDocument := range input.IdentificationDocuments {
			coding, ok := codingMap[string(identificationDocument.Type)]
			if !ok {
				return nil, fmt.Errorf("key %s not found", identificationDocument.Type)
			}
			codingInput = append(codingInput, &domain.FHIRCodingInput{
				System: &system,
				Code:   scalarutils.Code(coding),
			})
		}

		identifiers = append(identifiers, &domain.FHIRIdentifierInput{
			Use: domain.IdentifierUseEnumOfficial,
			Type: domain.FHIRCodeableConceptInput{
				Coding: codingInput,
			},
			System: &identifierSystem,
			Value:  idDocument.Number,
		})
	}

	var telecoms []*domain.PhoneNumberInput
	for _, phoneNumber := range input.PhoneNumbers {
		telecoms = append(telecoms, &domain.PhoneNumberInput{
			Msisdn:             phoneNumber,
			CommunicationOptIn: true,
		})
	}

	registrationInput := domain.SimplePatientRegistrationInput{
		Names:        []*domain.NameInput{{FirstName: input.FirstName, LastName: input.LastName, OtherNames: &input.OtherNames}},
		BirthDate:    input.BirthDate,
		PhoneNumbers: telecoms,
		Gender:       input.Gender.String(),
		Active:       true,
	}

	exists, err := c.CheckPatientExistenceUsingPhoneNumber(ctx, registrationInput)
	if err != nil {
		utils.ReportErrorToSentry(err)
		return nil, fmt.Errorf("unable to check patient existence")
	}

	if exists {
		return nil, fmt.Errorf("patient with phone number already exists")
	}

	patientInput, err := c.SimplePatientRegistrationInputToPatientInput(ctx, registrationInput)
	if err != nil {
		return nil, err
	}

	patientID := uuid.New().String()
	patientInput.ID = &patientID

	patientInput.Identifier = identifiers

	response, err := c.infrastructure.FHIR.CreateFHIRPatient(ctx, *patientInput)
	if err != nil {
		return nil, err
	}

	var phones []string
	for _, phone := range response.PatientRecord.Telecom {
		phones = append(phones, *phone.Value)
	}
	patient := dto.Patient{
		ID:          *response.PatientRecord.ID,
		Active:      *response.PatientRecord.Active,
		Name:        response.PatientRecord.Names(),
		PhoneNumber: phones,
		Gender:      string(*response.PatientRecord.Gender),
		BirthDate:   *response.PatientRecord.BirthDate,
	}

	return &patient, nil
}
