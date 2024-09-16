package db

import "time"

type User struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type OrganizationType string

const (
	IE  OrganizationType = "IE"
	LLC OrganizationType = "LLC"
	JSC OrganizationType = "JSC"
)

type Organization struct {
	Id               string           `json:"id"`
	Name             string           `json:"name"`
	Description      string           `json:"description"`
	OrganizationType OrganizationType `json:"organizationType"`
	CreatedAt        time.Time        `json:"createdAt"`
	UpdatedAt        time.Time        `json:"updatedAt"`
}

type OrganizationResponsible struct {
	Id             string `json:"id"`
	OrganizationId string `json:"organizationId"`
	UserId         string `json:"userId"`
}

type ServiceType string

const (
	Construction ServiceType = "Construction"
	Delivery     ServiceType = "Delivery"
	Manufacture  ServiceType = "Manufacture"
)

type Status string

const (
	Created   Status = "Created"
	Published Status = "Published"
	Closed    Status = "Closed"
)

type Tender struct {
	Id              string      `json:"id"`
	Name            string      `json:"name"`
	Description     string      `json:"description"`
	ServiceType     ServiceType `json:"serviceType"`
	Status          Status      `json:"status"`
	OrganizationId  string      `json:"organizationId"`
	CreatorUsername string      `json:"creatorUsername"`
	CreatedAt       time.Time   `json:"createdAt"`
	Version         uint        `json:"version"`
}

type AuthorType string

const (
	Organizatiom AuthorType = "Organization"
	UserAuthor   AuthorType = "User"
)

type Bid struct {
	Id              string     `json:"id"`
	Name            string     `json:"name"`
	CreatorUsername string     `json:"creatorUsername"`
	Description     string     `json:"description"`
	Status          Status     `json:"status"`
	TenderId        string     `json:"tenderId"`
	AuthorType      AuthorType `json:"authorType"`
	AuthorId        string     `json:"authorId"`
	CreatedAt       time.Time  `json:"createdAt"`
	Version         uint       `json:"version"`
	ApproveCount    uint       `json:"approveCount" default:"0"`
	RejectCount     uint       `json:"rejectCount" default:"0"`
}

type Feedback struct {
	Id             string    `json:"id"`
	BidId          string    `json:"bidId"`
	AuthorUsername string    `json:"authorUsername"`
	Comment        string    `json:"comment"`
	CreatedAt      time.Time `json:"createdAt"`
}
