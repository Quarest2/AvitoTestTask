package db

import "backend/db"

type GetTendersParams struct {
	Limit       uint             `json:"limit"`
	Offset      uint             `json:"offset"`
	ServiceType []db.ServiceType `json:"service_type"`
}

type NewTenderParams struct {
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	ServiceType     db.ServiceType `json:"serviceType"`
	OrganizationId  string         `json:"organizationId"`
	CreatorUsername string         `json:"creatorUsername"`
}

type GetMyTendersParams struct {
	Limit    uint   `json:"limit"`
	Offset   uint   `json:"offset"`
	Username string `json:"username"`
}

type GetTenderStatusParams struct {
	TenderId string `json:"tenderId"`
	Username string `json:"username"`
}

type ChangeTenderStatusParams struct {
	TenderId string    `json:"tenderId"`
	Status   db.Status `json:"status"`
	Username string    `json:"username"`
}

type UpdateTenderParams struct {
	TenderId    string         `json:"tenderId"`
	Username    string         `json:"username"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	ServiceType db.ServiceType `json:"serviceType"`
}

type RollbackTenderParams struct {
	TenderId string `json:"tenderId"`
	Version  string `json:"version"`
	Username string `json:"username"`
}
