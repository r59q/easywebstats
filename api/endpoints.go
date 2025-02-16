package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v0

// RegisterNumStat godoc
// @Summary Register numeric stat
// @Schemes
// @Param message body api.NumStatRegistration true "test"
// @Description Register a numeric stat
// @Accept json
// @Produce json
// @Success 200 {number} Current value
// @Router /register/num [post]
func RegisterNumStat(g *gin.Context) {
	var registration NumStatRegistration
	err := g.BindJSON(&registration)
	if err != nil {
		g.JSON(http.StatusInternalServerError, "Failed to capture registration object. Is it formatted correctly?")
		return
	}
	println(registration.Name)
	g.JSON(http.StatusOK, "helloworld")
}
