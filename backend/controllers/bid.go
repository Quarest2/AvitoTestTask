package controllers

import (
	"backend/db"
	b "backend/db/bids"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Creating new bid
// @Description Creating new bid
// @ID new-bid
// @Tags bids
// @Accept json
// @Produce json
// @Param input body b.NewBidParams true "New bid params"
// @Success 200 {object} db.Bid
// @Failure 400 {object} FailureResponse
// @Router /bids/new [post]
func BidsNewHandler(c *gin.Context) {
	var newBid b.NewBidParams
	var bid db.Bid
	var err error

	data, _ := c.GetRawData()

	if err = json.Unmarshal(data, &newBid); err != nil {
		c.JSON(http.StatusInternalServerError, FailureResponse{Error: err.Error()})
	}

	if bid, err = b.New(newBid); err != nil {
		c.JSON(http.StatusInternalServerError, FailureResponse{Error: err.Error()})
	}

	c.JSON(http.StatusOK, bid)
}

// @Summary Get my bids
// @Description Get my bids
// @ID get-my-bids
// @Tags bids
// @Produce json
// @Param limit query string false "Limit"
// @Param offset query string false "Offset"
// @Param username query string true "Username"
// @Success 200 {object} []db.Bid
// @Failure 400 {object} FailureResponse
// @Router /bids/my [get]
func BidsMyHandler(c *gin.Context) {
	var err error
	var getMyBidsParams b.GetMyBidsParams

	limitStr := c.Query("limit")
	var limitInt int
	if limitInt, err = strconv.Atoi(limitStr); err != nil {
		getMyBidsParams.Limit = 0
	} else {
		getMyBidsParams.Limit = uint(limitInt)
	}

	offsetStr := c.Query("offset")
	var offsetInt int
	if offsetInt, err = strconv.Atoi(offsetStr); err != nil {
		getMyBidsParams.Limit = 0
	} else {
		getMyBidsParams.Limit = uint(offsetInt)
	}

	username := c.Query("username")
	getMyBidsParams.Username = username

	var bids []db.Bid
	if bids, err = b.GetMyBids(getMyBidsParams); err != nil {
		c.JSON(http.StatusInternalServerError, FailureResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, bids)
}

// @Summary Get bids by tender
// @Description Get bids by tender
// @ID get-bids-by-tender
// @Tags bids
// @Produce json
// @Param tenderId path string true "TenderId"
// @Param username query string true "Username"
// @Param limit query string false "Limit"
// @Param offset query string false "Offset"
// @Success 200 {object} []db.Bid
// @Failure 400 {object} FailureResponse
// @Router /bids/{bidId}/list [get]
func BidsListHandler(c *gin.Context) {
	var bidsListParams b.GetListOfBidsByTenderParams
	var err error

	limitStr := c.Query("limit")
	var limitInt int
	if limitInt, err = strconv.Atoi(limitStr); err != nil {
		bidsListParams.Limit = 0
	} else {
		bidsListParams.Limit = uint(limitInt)
	}

	offsetStr := c.Query("offset")
	var offsetInt int
	if offsetInt, err = strconv.Atoi(offsetStr); err != nil {
		bidsListParams.Limit = 0
	} else {
		bidsListParams.Limit = uint(offsetInt)
	}

	username := c.Query("username")
	bidsListParams.Username = username

	tenderId := c.Param("tenderId")
	bidsListParams.TenderId = tenderId

	var bids []db.Bid
	if bids, err = b.GetListOfBidsByTender(bidsListParams); err != nil {
		c.JSON(http.StatusInternalServerError, FailureResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, bids)
}

// @Summary Get bid status
// @Description Get bid status
// @ID get-bid-status
// @Tags bids
// @Produce json
// @Param username query string false "Username"
// @Param bidId path string true "BidId"
// @Success 200 {object} string
// @Failure 400 {object} FailureResponse
// @Router /bids/{bidId}/status [get]
func BidsGetStatusHandler(c *gin.Context) {
	var err error
	var getBidStatusParams b.GetBidStatusParams

	username := c.Query("username")
	getBidStatusParams.Username = username

	bidId := c.Param("bidId")
	getBidStatusParams.BidId = bidId

	var status string
	if status, err = b.GetStatus(getBidStatusParams); err != nil {
		c.JSON(http.StatusInternalServerError, FailureResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, status)
}

// @Summary Change bid status
// @Description Change bid status
// @ID change-bid-status
// @Tags bids
// @Produce json
// @Param bidId path string true "BidId"
// @Param status query string true "Status"
// @Param username query string true "Username"
// @Success 200 {object} db.Bid
// @Failure 400 {object} FailureResponse
// @Router /bids/{bidId}/status [put]
func BidsPutStatusHandler(c *gin.Context) {
	var err error
	var changeBidStatusParams b.ChangeBidStatusParams

	bidId := c.Param("bidId")
	changeBidStatusParams.BidId = bidId

	status := c.Query("status")
	changeBidStatusParams.Status = db.Status(status)

	username := c.Query("username")
	changeBidStatusParams.Username = username

	var bid db.Bid
	if bid, err = b.ChangeStatus(changeBidStatusParams); err != nil {
		c.JSON(http.StatusInternalServerError, FailureResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, bid)
}

// @Summary Edit bid params
// @Description Edit bid params
// @ID edit-bid-params
// @Tags bids
// @Accept json
// @Produce json
// @Param bidId path string true "BidId"
// @Param username query string true "Username"
// @Param input body b.UpdateBidParams true "Bid new params"
// @Success 200 {object} db.Bid
// @Failure 400 {object} FailureResponse
// @Router /bids/{bidId}/edit [patch]
func BidsEditHandler(c *gin.Context) {
	var updateBidParams b.UpdateBidParams
	var bid db.Bid
	var err error

	data, _ := c.GetRawData()

	if err = json.Unmarshal(data, &updateBidParams); err != nil {
		c.JSON(http.StatusInternalServerError, FailureResponse{Error: err.Error()})
	}

	bidId := c.Param("bidId")
	updateBidParams.BidId = bidId

	username := c.Query("username")
	updateBidParams.Username = username

	if bid, err = b.Update(updateBidParams); err != nil {
		c.JSON(http.StatusInternalServerError, FailureResponse{Error: err.Error()})
	}

	c.JSON(http.StatusOK, bid)
}

// @Summary Submit decision on bid
// @Description Submit decision on bid
// @ID submit-decision-on-bid
// @Tags bids
// @Produce json
// @Param bidId path string true "BidId"
// @Param decision query db.Decision true "Decision"
// @Param username query string true "Username"
// @Success 200 {object} db.Bid
// @Failure 400 {object} FailureResponse
// @Router /bids/{bidId}/submit_decision [put]
func BidsSubmitDecisionHandler(c *gin.Context) {
	var err error
	var submitDecisionParams b.SubmitDecisionParams

	bidId := c.Param("bidId")
	submitDecisionParams.BidId = bidId

	decision := c.Query("decision")
	submitDecisionParams.Decision = b.Decision(decision)

	username := c.Query("username")
	submitDecisionParams.Username = username

	var bid db.Bid
	if bid, err = b.SubmitDecision(submitDecisionParams); err != nil {
		c.JSON(http.StatusInternalServerError, FailureResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, bid)
}

// @Summary Send a feedback on bid
// @Description Send a feedback on bid
// @ID send-a-feedback-on-bid
// @Tags bids
// @Produce json
// @Param bidId path string true "BidId"
// @Param bidFeedback query string true "BidFeedback"
// @Param username query string true "Username"
// @Success 200 {object} db.Bid
// @Failure 400 {object} FailureResponse
// @Router /bids/{bidId}/feedback [put]
func BidsFeedbackHandler(c *gin.Context) {
	var err error
	var feedback b.FeedbackParams

	bidId := c.Param("bidId")
	feedback.BidId = bidId

	bidFeedback := c.Query("bidFeedback")
	feedback.BidFeedback = bidFeedback

	username := c.Query("username")
	feedback.Username = username

	var bid db.Bid
	if bid, err = b.Feedback(feedback); err != nil {
		c.JSON(http.StatusInternalServerError, FailureResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, bid)
}

//// @Summary Get feedbacks by tender
//// @Description Get feedbacks by tender
//// @ID get-feedbacks-by-tender
//// @Tags bids
//// @Produce json
//// @Param tenderId path string true "TenderId"
//// @Param authorUsername query string true "AuthorUsername"
//// @Param requesterUsername query string true "RequesterUsername"
//// @Param limit query string false "Limit"
//// @Param offset query string false "Offset"
//// @Success 200 {object} []db.Feedback
//// @Failure 400 {object} FailureResponse
//// @Router /bids/{tenderId}/reviews [get]
//func BidsReviewsHandler(c *gin.Context) {
//	var err error
//	var getReviewsParams b.GetReviewsParams
//
//	tenderId := c.Param("tenderId")
//	getReviewsParams.TenderId = tenderId
//
//	author := c.Query("authorUsername")
//	getReviewsParams.AuthorUsername = author
//
//	requester := c.Query("authorUsername")
//	getReviewsParams.RequesterUsername = requester
//
//	limitStr := c.Query("limit")
//	var limitInt int
//	if limitInt, err = strconv.Atoi(limitStr); err != nil {
//		getReviewsParams.Limit = 0
//	} else {
//		getReviewsParams.Limit = uint(limitInt)
//	}
//
//	offsetStr := c.Query("offset")
//	var offsetInt int
//	if offsetInt, err = strconv.Atoi(offsetStr); err != nil {
//		getReviewsParams.Limit = 0
//	} else {
//		getReviewsParams.Limit = uint(offsetInt)
//	}
//
//	var feedbacks []db.Feedback
//	if feedbacks, err = b.GetReviews(getReviewsParams); err != nil {
//		c.JSON(http.StatusInternalServerError, FailureResponse{Error: err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, feedbacks)
//}
