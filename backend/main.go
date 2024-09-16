package main

import (
	"backend/controllers"
	"backend/db"
	"backend/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func main() {
	db.Connect()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "OPTIONS"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Authorization"},
		AllowHeaders:     []string{"Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	docs.SwaggerInfo.Title = "TENDERS API"
	docs.SwaggerInfo.Description = "This is an API  of TENDERS project"
	docs.SwaggerInfo.BasePath = "/api"

	swaggerUrl := ginSwagger.URL("./doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerUrl))

	apiGroup := r.Group("/api")

	apiGroup.GET("/ping", controllers.Ping)

	tendersGroup := apiGroup.Group("/tenders")
	tendersGroup.GET("", controllers.TendersHandler)
	tendersGroup.POST("/new", controllers.TendersNewHandler)
	tendersGroup.GET("/my", controllers.TendersMyHandler)
	tendersGroup.GET("/:tenderId/status", controllers.TendersGetStatusHandler)
	tendersGroup.PUT("/:tenderId/status", controllers.TendersPutStatusHandler)
	tendersGroup.PATCH("/:tenderId/edit", controllers.TendersEditHandler)
	//tendersGroup.PUT("/:tenderId/rollback/:version", controllers.TendersRollbackHandler)

	bidsGroup := apiGroup.Group("/bids")
	bidsGroup.POST("/new", controllers.BidsNewHandler)
	bidsGroup.GET("/my", controllers.BidsMyHandler)
	bidsGroup.GET("/:bidId/list", controllers.BidsListHandler)
	bidsGroup.GET("/:bidId/status", controllers.BidsGetStatusHandler)
	bidsGroup.PUT("/:bidId/status", controllers.BidsPutStatusHandler)
	bidsGroup.PATCH("/:bidId/edit", controllers.BidsEditHandler)
	bidsGroup.PUT("/:bidId/submit_decision", controllers.BidsSubmitDecisionHandler)
	bidsGroup.PUT("/:bidId/feedback", controllers.BidsFeedbackHandler)
	//bidsGroup.PUT("/:bidId/rollback/:version", controllers.BidsRollbackHandler)
	//bidsGroup.GET("/:tenderId/reviews", controllers.BidsReviewsHandler)

	r.Run(":8080")
}
