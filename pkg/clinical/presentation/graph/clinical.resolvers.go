package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"

	"github.com/savannahghi/clinical/pkg/clinical/application/dto"
	"github.com/savannahghi/clinical/pkg/clinical/presentation/graph/generated"
)

// CreateEpisodeOfCare is the resolver for the createEpisodeOfCare field.
func (r *mutationResolver) CreateEpisodeOfCare(ctx context.Context, episodeOfCare dto.EpisodeOfCareInput) (*dto.EpisodeOfCare, error) {
	r.CheckDependencies()

	return r.usecases.CreateEpisodeOfCare(ctx, episodeOfCare)
}

// EndEpisodeOfCare is the resolver for the endEpisodeOfCare field.
func (r *mutationResolver) EndEpisodeOfCare(ctx context.Context, id string) (*dto.EpisodeOfCare, error) {
	r.CheckDependencies()
	return r.usecases.EndEpisodeOfCare(ctx, id)
}

// StartEncounter is the resolver for the startEncounter field.
func (r *mutationResolver) StartEncounter(ctx context.Context, episodeID string) (string, error) {
	r.CheckDependencies()
	return r.usecases.Clinical.StartEncounter(ctx, episodeID)
}

// EndEncounter is the resolver for the endEncounter field.
func (r *mutationResolver) EndEncounter(ctx context.Context, encounterID string) (bool, error) {
	r.CheckDependencies()
	return r.usecases.Clinical.EndEncounter(ctx, encounterID)
}

// RecordTemperature is the resolver for the recordTemperature field.
func (r *mutationResolver) RecordTemperature(ctx context.Context, input dto.ObservationInput) (*dto.Observation, error) {
	r.CheckDependencies()
	return r.usecases.Clinical.RecordTemperature(ctx, input)
}

// RecordHeight is the resolver for the recordHeight field.
func (r *mutationResolver) RecordHeight(ctx context.Context, input dto.ObservationInput) (*dto.Observation, error) {
	r.CheckDependencies()
	return r.usecases.Clinical.RecordHeight(ctx, input)
}

// RecordWeight is the resolver for the recordWeight field.
func (r *mutationResolver) RecordWeight(ctx context.Context, input dto.ObservationInput) (*dto.Observation, error) {
	r.CheckDependencies()
	return r.usecases.Clinical.RecordWeight(ctx, input)
}

// RecordRespiratoryRate is the resolver for the recordRespiratoryRate field.
func (r *mutationResolver) RecordRespiratoryRate(ctx context.Context, input dto.ObservationInput) (*dto.Observation, error) {
	r.CheckDependencies()
	return r.usecases.Clinical.RecordRespiratoryRate(ctx, input)
}

// RecordPulseRate is the resolver for the recordPulseRate field.
func (r *mutationResolver) RecordPulseRate(ctx context.Context, input dto.ObservationInput) (*dto.Observation, error) {
	r.CheckDependencies()
	return r.usecases.Clinical.RecordPulseRate(ctx, input)
}

// RecordBloodPressure is the resolver for the recordBloodPressure field.
func (r *mutationResolver) RecordBloodPressure(ctx context.Context, input dto.ObservationInput) (*dto.Observation, error) {
	r.CheckDependencies()
	return r.usecases.Clinical.RecordBloodPressure(ctx, input)
}

// RecordBmi is the resolver for the recordBMI field.
func (r *mutationResolver) RecordBmi(ctx context.Context, input dto.ObservationInput) (*dto.Observation, error) {
	r.CheckDependencies()
	return r.usecases.Clinical.RecordBMI(ctx, input)
}

// CreatePatient is the resolver for the createPatient field.
func (r *mutationResolver) CreatePatient(ctx context.Context, input dto.PatientInput) (*dto.Patient, error) {
	r.CheckDependencies()

	return r.usecases.CreatePatient(ctx, input)
}

// CreateCondition is the resolver for the createCondition field.
func (r *mutationResolver) CreateCondition(ctx context.Context, input dto.ConditionInput) (*dto.Condition, error) {
	r.CheckDependencies()
	return r.usecases.CreateCondition(ctx, input)
}

// CreateAllergyIntolerance is the resolver for the createAllergyIntolerance field.
func (r *mutationResolver) CreateAllergyIntolerance(ctx context.Context, input dto.AllergyInput) (*dto.Allergy, error) {
	return r.usecases.CreateAllergyIntolerance(ctx, input)
}

// PatientHealthTimeline is the resolver for the patientHealthTimeline field.
func (r *queryResolver) PatientHealthTimeline(ctx context.Context, input dto.HealthTimelineInput) (*dto.HealthTimeline, error) {
	r.CheckDependencies()

	return r.usecases.PatientHealthTimeline(ctx, input)
}

// GetMedicalData is the resolver for the getMedicalData field.
func (r *queryResolver) GetMedicalData(ctx context.Context, patientID string) (*dto.MedicalData, error) {
	r.CheckDependencies()

	return r.usecases.GetMedicalData(ctx, patientID)
}

// GetEpisodeOfCare is the resolver for the getEpisodeOfCare field.
func (r *queryResolver) GetEpisodeOfCare(ctx context.Context, id string) (*dto.EpisodeOfCare, error) {
	r.CheckDependencies()
	return r.usecases.GetEpisodeOfCare(ctx, id)
}

// ListPatientConditions is the resolver for the listPatientConditions field.
func (r *queryResolver) ListPatientConditions(ctx context.Context, patientID string, pagination dto.Pagination) (*dto.ConditionConnection, error) {
	r.CheckDependencies()
	return r.usecases.ListPatientConditions(ctx, patientID, pagination)
}

// ListPatientEncounters is the resolver for the listPatientEncounters field.
func (r *queryResolver) ListPatientEncounters(ctx context.Context, patientID string) ([]*dto.Encounter, error) {
	r.CheckDependencies()
	return r.usecases.Clinical.ListPatientEncounters(ctx, patientID)
}

// GetPatientTemperatureEntries is the resolver for the getPatientTemperatureEntries field.
func (r *queryResolver) GetPatientTemperatureEntries(ctx context.Context, patientID string) ([]*dto.Observation, error) {
	r.CheckDependencies()
	return r.usecases.Clinical.GetPatientTemperatureEntries(ctx, patientID)
}

// GetPatientBloodPressureEntries is the resolver for the getPatientBloodPressureEntries field.
func (r *queryResolver) GetPatientBloodPressureEntries(ctx context.Context, patientID string) ([]*dto.Observation, error) {
	r.CheckDependencies()
	return r.usecases.Clinical.GetPatientBloodPressureEntries(ctx, patientID)
}

// GetPatientHeightEntries is the resolver for the getPatientHeightEntries field.
func (r *queryResolver) GetPatientHeightEntries(ctx context.Context, patientID string) ([]*dto.Observation, error) {
	return r.usecases.Clinical.GetPatientHeightEntries(ctx, patientID)
}

// GetPatientRespiratoryRateEntries is the resolver for the getPatientRespiratoryRateEntries field.
func (r *queryResolver) GetPatientRespiratoryRateEntries(ctx context.Context, patientID string) ([]*dto.Observation, error) {
	r.CheckDependencies()
	return r.usecases.Clinical.GetPatientRespiratoryRateEntries(ctx, patientID)
}

// GetPatientPulseRateEntries is the resolver for the GetPatientPulseRateEntries field.
func (r *queryResolver) GetPatientPulseRateEntries(ctx context.Context, patientID string) ([]*dto.Observation, error) {
	return r.usecases.Clinical.GetPatientPulseRateEntries(ctx, patientID)
}

// GetPatientBMIEntries is the resolver for the getPatientBMIEntries field.
func (r *queryResolver) GetPatientBMIEntries(ctx context.Context, patientID string) ([]*dto.Observation, error) {
	r.CheckDependencies()
	return r.usecases.Clinical.GetPatientBMIEntries(ctx, patientID)
}

// GetPatientWeightEntries is the resolver for the getPatientWeightEntries field.
func (r *queryResolver) GetPatientWeightEntries(ctx context.Context, patientID string) ([]*dto.Observation, error) {
	return r.usecases.Clinical.GetPatientWeightEntries(ctx, patientID)
}

// SearchAllergy is the resolver for the searchAllergy field.
func (r *queryResolver) SearchAllergy(ctx context.Context, name string) ([]*dto.Terminology, error) {
	r.CheckDependencies()

	return r.usecases.SearchAllergy(ctx, name)
}

// GetAllergy is the resolver for the getAllergy field.
func (r *queryResolver) GetAllergy(ctx context.Context, id string) (*dto.Allergy, error) {
	r.CheckDependencies()

	return r.usecases.GetAllergyIntolerance(ctx, id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
