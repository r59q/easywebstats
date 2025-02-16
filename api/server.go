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
		register.POST("/num", RegisterNumStat)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
