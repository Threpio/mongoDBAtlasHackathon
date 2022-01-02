package types

type Collection struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Public         bool   `json:"public"`
	Active         bool   `json:"active"`
	OrganisationID string `json:"organisation_id"`
	OwnerID        string `json:"owner_id"`
}

type CollectionCreationRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Public         bool   `json:"public"`
	OrganisationID string `json:"organisation_id"`
	OwnerID        string `json:"owner_id"`
}