package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"r59q.com/easywebstats/docs"
)

// gin-swagger middleware
// swagger embed files

// album represents data about a record album.
type StatRegistration struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// @BasePath /api/v1

// Register godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Register(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v0"
	v1 := r.Group("/api/v0")
	{
		eg := v1.Group("/example")
		{
			eg.POST("/helloworld", Register)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
