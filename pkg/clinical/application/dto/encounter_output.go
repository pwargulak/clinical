package dto

// Encounter definition: an interaction between a patient and healthcare provider(s) for the purpose of providing healthcare service(s) or assessing the health status of a patient.
type Encounter struct {
	ID              string              `json:"id,omitempty"`
	Status          EncounterStatusEnum `json:"status,omitempty"`
	Class           EncounterClass      `json:"class,omitempty"`
	PatientID       string              `json:"patientID,omitempty"`
	EpisodeOfCareID string              `json:"episodeOfCareID,omitempty"`
}

// EncounterConnection is the encounter connection type
type EncounterConnection struct {
	TotalCount int
	Edges      []EncounterEdge
	PageInfo   PageInfo
}

// EncounterEdge is an patient encounter edge
type EncounterEdge struct {
	Node   Encounter
	Cursor string
}

// CreateEncounterConnection creates a connection that follows the GraphQl Cursor Connection Specification
func CreateEncounterConnection(encounters []*Encounter, pageInfo PageInfo, total int) EncounterConnection {
	connection := EncounterConnection{
		TotalCount: total,
		Edges:      []EncounterEdge{},
		PageInfo:   pageInfo,
	}

	for _, encounter := range encounters {
		edge := EncounterEdge{
			Node:   *encounter,
			Cursor: encounter.ID,
		}

		connection.Edges = append(connection.Edges, edge)
	}

	return connection
}
