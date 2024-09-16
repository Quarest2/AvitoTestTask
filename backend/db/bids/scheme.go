package db

import "backend/db"

type NewBidParams struct {
	Name            string        `json:"name"`
	CreatorUsername string        `json:"creatorUsername"`
	Description     string        `json:"description"`
	TenderId        string        `json:"tenderId"`
	AuthorType      db.AuthorType `json:"authorType"`
	AuthorId        string        `json:"authorId"`
}

type GetMyBidsParams struct {
	Limit    uint   `json:"limit"`
	Offset   uint   `json:"offset"`
	Username string `json:"username"`
}

type GetListOfBidsByTenderParams struct {
	TenderId string `json:"bidId"`
	Username string `json:"username"`
	Limit    uint   `json:"limit"`
	Offset   uint   `json:"offset"`
}

type GetBidStatusParams struct {
	BidId    string `json:"bidId"`
	Username string `json:"username"`
}

type ChangeBidStatusParams struct {
	BidId    string    `json:"bidId"`
	Status   db.Status `json:"status"`
	Username string    `json:"username"`
}

type UpdateBidParams struct {
	BidId       string `json:"bidId"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Decision string

const (
	Approved Decision = "Approved"
	Rejected Decision = "Rejected"
)

type SubmitDecisionParams struct {
	BidId    string   `json:"bidId"`
	Decision Decision `json:"decision"`
	Username string   `json:"username"`
}

type FeedbackParams struct {
	BidId       string `json:"bidId"`
	BidFeedback string `json:"bidFeedback"`
	Username    string `json:"username"`
}
type RollbackBidParams struct {
	BidId    string `json:"bidId"`
	Version  string `json:"version"`
	Username string `json:"username"`
}

type GetReviewsParams struct {
	TenderId          string `json:"tenderId"`
	AuthorUsername    string `json:"authorUsername"`
	RequesterUsername string `json:"requesterUsername"`
	Limit             uint   `json:"limit"`
	Offset            uint   `json:"offset"`
}
