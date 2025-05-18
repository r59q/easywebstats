package api

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"r59q.com/easywebstats/docs"
)

func RunGinSever() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	{ // Version 1
		v1 := r.Group("/api/v1")
		v1.GET("/serialize", Serialize)

		register := v1.Group("/register")
		register.POST("/num/set", SetNumStat)
		register.POST("/num/increase", IncreaseNumStat)
		register.POST("/num/decrease", DecreaseNumStat)

		read := v1.Group("/read")
		read.GET("/num/value/:name/:label", ReadStatNameLabel)
		read.GET("/num/value/:name", ReadStatName)
		read.GET("/num/names", ReadNumNames)
		read.GET("/num/exponentialrate/:name/:label", ReadStatNameLabelExponentialRate)
		read.GET("/num/exponentialrate/:name", ReadStatNameExponentialRates)
		read.GET("/num/mean/:name", ReadNumMean)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/metrics", GetMetrics)
	port := os.Getenv("EWS_PORT")
	if len(port) == 0 {
		port = "8080"
	}
	r.Run(":" + port)
}
