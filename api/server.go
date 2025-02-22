package api

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"r59q.com/easywebstats/docs"
)

func RunGinSever() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v0"
	v0 := r.Group("/api/v0")
	{
		register := v0.Group("/register")
		register.POST("/num/set", SetNumStat)
		register.POST("/num/increase", IncreaseNumStat)
		register.POST("/num/decrease", DecreaseNumStat)

		read := v0.Group("/read")
		read.GET("/num/value/:name/:label", ReadStatNameLabel)
		read.GET("/num/value/:name", ReadStatName)
		read.GET("/num/names", ReadNumNames)
		read.GET("/num/exponentialrate/:name/:label", ReadStatNameLabelExponentialRate)
		read.GET("/num/exponentialrate/:name", ReadStatNameExponentialRates)
		read.GET("/num/mean/:name", ReadNumMean)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/metrics", GetMetrics)
	r.Run(":8080")
}
