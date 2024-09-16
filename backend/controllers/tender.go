package controllers

import (
	"backend/db"
	ten "backend/db/tenders"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Get tenders
// @Description Get tenders
// @ID get-tenders
// @Tags tenders
// @Produce  json
// @Param limit query string false "Limit"
// @Param offset query string false "Offset"
// @Param service_type query []db.ServiceType false "Service Type"
// @Success 200 {object} []db.Tender
// @Failure 400 {object} FailureResponse
// @Router /tenders [get]
func TendersHandler(c *gin.Context) {
	var err error
	var getTendersParams ten.GetTendersParams

	limitStr := c.Query("limit")
	var limitInt int
	if limitInt, err = strconv.Atoi(limitStr); err != nil {
		getTendersParams.Limit = 0
	} else {
		getTendersParams.Limit = uint(limitInt)
	}

	offsetStr := c.Query("offset")
	var offsetInt int
	if offsetInt, err = strconv.Atoi(offsetStr); err != nil {
		getTendersParams.Limit = 0
	} else {
		getTendersParams.Limit = uint(offsetInt)
	}

	serviceTypes := c.QueryArray("service_type")
	for _, serviceType := range serviceTypes {
		getTendersParams.ServiceType = append(getTendersParams.ServiceType, db.ServiceType(serviceType))
	}

	var tenders []db.Tender
	if tenders, err = ten.Get(getTendersParams); err != nil {
		c.JSON(http.StatusInternalServerError, FailureResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tenders)
}

// @Summary Creating new tender
// @Description Creating new tender
// @ID new-tender
// @Tags tenders
// @Accept json
// @Produce json
// @Param input body ten.NewTenderParams true "New tender params"
// @Success 200 {object} db.Tender
// @Failure 400 {object} FailureResponse
// @Router /tenders/new [post]
func TendersNewHandler(c *gin.Context) {
	var newTen ten.NewTenderParams
	var tender db.Tender
	var err error

	data, _ := c.GetRawData()

	if err = json.Unmarshal(data, &newTen); err != nil {
		c.JSON(http.StatusInternalServerError, FailureResponse{Error: err.Error()})
	}

	if tender, err = ten.New(newTen); err != nil {
		c.JSON(http.StatusInternalServerError, FailureResponse{Error: err.Error()})
	}

	c.JSON(http.StatusOK, tender)
}

// @Summary Get my tenders
// @Description Get my tenders
// @ID get-my-tenders
// @Tags tenders
// @Produce  json
// @Param limit query string false "Limit"
// @Param offset query string false "Offset"
// @Param username query string true "Username"
// @Success 200 {object} []db.Tender
// @Failure 400 {object} FailureResponse
// @Router /tenders/my [get]
func TendersMyHandler(c *gin.Context) {
	var err error
	var getMyTendersParams ten.GetMyTendersParams

	limitStr := c.Query("limit")
	var limitInt int
	if limitInt, err = strconv.Atoi(limitStr); err != nil {
		getMyTendersParams.Limit = 0
	} else {
		getMyTendersParams.Limit = uint(limitInt)
	}

	offsetStr := c.Query("offset")
	var offsetInt int
	if offsetInt, err = strconv.Atoi(offsetStr); err != nil {
		getMyTendersParams.Limit = 0
	} else {
		getMyTendersParams.Limit = uint(offsetInt)
	}

	username := c.Query("username")
	getMyTendersParams.Username = username

	var tenders []db.Tender
	if tenders, err = ten.GetMyTenders(getMyTendersParams); err != nil {
		c.JSON(http.StatusInternalServerError, FailureResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tenders)
}

// @Summary Get tender status
// @Description Get tender status
// @ID get-tender-status
// @Tags tenders
// @Produce  json
// @Param username query string false "Username"
// @Param tenderId path string true "TenderId"
// @Success 200 {object} string
// @Failure 400 {object} FailureResponse
// @Router /tenders/{tenderId}/status [get]
func TendersGetStatusHandler(c *gin.Context) {
	var err error
	var getTenderStatusParams ten.GetTenderStatusParams

	username := c.Query("username")
	getTenderStatusParams.Username = username

	tenderId := c.Param("tenderId")
	getTenderStatusParams.TenderId = tenderId

	var status string
	if status, err = ten.GetStatus(getTenderStatusParams); err != nil {
		c.JSON(http.StatusInternalServerError, FailureResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, status)
}

// @Summary Change tender status
// @Description Change tender status
// @ID change-tender-status
// @Tags tenders
// @Produce json
// @Param tenderId path string true "TenderId"
// @Param status query string true "Status"
// @Param username query string true "Username"
// @Success 200 {object} db.Tender
// @Failure 400 {object} FailureResponse
// @Router /tenders/{tenderId}/status [put]
func TendersPutStatusHandler(c *gin.Context) {
	var err error
	var changeTenderStatusParams ten.ChangeTenderStatusParams

	tenderId := c.Param("tenderId")
	changeTenderStatusParams.TenderId = tenderId

	status := c.Query("status")
	changeTenderStatusParams.Status = db.Status(status)

	username := c.Query("username")
	changeTenderStatusParams.Username = username

	var tender db.Tender
	if tender, err = ten.ChangeStatus(changeTenderStatusParams); err != nil {
		c.JSON(http.StatusInternalServerError, FailureResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tender)
}

// @Summary Edit tender params
// @Description Edit tender params
// @ID edit-tender-params
// @Tags tenders
// @Accept json
// @Produce json
// @Param tenderId path string true "TenderId"
// @Param username query string true "Username"
// @Param input body ten.UpdateTenderParams true "Tender new params"
// @Success 200 {object} db.Tender
// @Failure 400 {object} FailureResponse
// @Router /tenders/{tenderId}/edit [patch]
func TendersEditHandler(c *gin.Context) {
	var updateTenderParams ten.UpdateTenderParams
	var tender db.Tender
	var err error

	data, _ := c.GetRawData()

	if err = json.Unmarshal(data, &updateTenderParams); err != nil {
		c.JSON(http.StatusInternalServerError, FailureResponse{Error: err.Error()})
	}

	tenderId := c.Param("tenderId")
	updateTenderParams.TenderId = tenderId

	username := c.Query("username")
	updateTenderParams.Username = username

	if tender, err = ten.Update(updateTenderParams); err != nil {
		c.JSON(http.StatusInternalServerError, FailureResponse{Error: err.Error()})
	}

	c.JSON(http.StatusOK, tender)
}

//func TendersRollbackHandler (c *gin.Context) {
//
//}
