package dto

import "time"

// This file will be used to hold the payload for data that has been published to
// different pubsub topics

// CreatePatientPubSubMessage models the payload that is published to the `patient.create`
// topic
type CreatePatientPubSubMessage struct {
	ID                string     `json:"id"`
	Active            bool       `json:"active"`
	ClientType        string     `json:"client_type"`
	EnrollmentDate    *time.Time `json:"enrollment_date"`
	FHIRPatientID     string     `json:"fhir_patient_id"`
	EMRHealthRecordID string     `json:"emr_health_record_id"`
	TreatmentBuddy    string     `json:"treatment_buddy"`
	Counselled        bool       `json:"counselled"`
	Organisation      string     `json:"organisation"`
	UserID            string     `json:"user"`
	CurrentFacilityID string     `json:"current_facility"`
	CHV               string     `json:"chv"`
	Caregiver         string     `json:"caregiver"`
}

// CreateFacilityPubSubMessage models the details of healthcare facilities that are on myCareHub platform.
// This will be used to create a FHIR organization
type CreateFacilityPubSubMessage struct {
	ID          *string `json:"id"`
	Name        string  `json:"name"`
	Code        int     `json:"code"`
	Phone       string  `json:"phone"`
	Active      bool    `json:"active"`
	County      string  `json:"county"`
	Description string  `json:"description"`
}

// CreateVitalSignPubSubMessage models the details that will be posted to the vitals pub/sub topic
type CreateVitalSignPubSubMessage struct {
	PatientID      string    `json:"patientID"`
	OrganizationID string    `json:"organizationID"`
	Name           string    `json:"name"`
	ConceptID      *string   `json:"conceptId"`
	Value          string    `json:"value"`
	Date           time.Time `json:"date"`
}
