package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"r59q.com/easywebstats/internal"
)

// @BasePath /api/v0

// SetNumStat godoc
// @Summary Set numeric stat
// @Schemes
// @Param message body api.NumStatRegistration true "Stat and value"
// @Description Set a numeric stat
// @Accept json
// @Produce json
// @Success 200 {number} Current value
// @Router /register/num/set [post]
func SetNumStat(g *gin.Context) {
	var registration NumStatRegistration
	err := g.BindJSON(&registration)
	if err != nil {
		g.JSON(http.StatusInternalServerError, "Failed to capture registration object. Is it formatted correctly?")
		return
	}
	value := internal.HandleSetNum(registration.Name, registration.Label, registration.Value)
	g.JSON(http.StatusOK, value)
}

// IncreaseNumStat godoc
// @Summary Increase a numeric stat
// @Schemes
// @Param message body api.NumStatRegistration true "Stat and value"
// @Description Increase a numeric stat
// @Accept json
// @Produce json
// @Success 200 {number} Current value
// @Router /register/num/increase [post]
func IncreaseNumStat(g *gin.Context) {
	var registration NumStatRegistration
	err := g.BindJSON(&registration)
	if err != nil {
		g.JSON(http.StatusInternalServerError, "Failed to capture registration object. Is it formatted correctly?")
		return
	}
	value := internal.HandleIncreaseNum(registration.Name, registration.Label, registration.Value)
	g.JSON(http.StatusOK, value)
}

// DecreaseNumStat godoc
// @Summary Decrease a numeric stat
// @Schemes
// @Param message body api.NumStatRegistration true "Stat and value"
// @Description Decrease a numeric stat
// @Accept json
// @Produce json
// @Success 200 {number} Current value
// @Router /register/num/decrease [post]
func DecreaseNumStat(g *gin.Context) {
	var registration NumStatRegistration
	err := g.BindJSON(&registration)
	if err != nil {
		g.JSON(http.StatusInternalServerError, "Failed to capture registration object. Is it formatted correctly?")
		return
	}
	value := internal.HandleDecreaseNum(registration.Name, registration.Label, registration.Value)
	g.JSON(http.StatusOK, value)
}

// ReadStatNameLabel godoc
// @Summary Read a specific statistic
// @Schemes
// @Param name path string true "Stat name"
// @Param label path string true "Stat label"
// @Description Read a specific statistic
// @Accept json
// @Produce json
// @Success 200 {number} Current Value
// @Router /read/num/{name}/{label} [get]
func ReadStatNameLabel(g *gin.Context) {
	name := g.Params.ByName("name")
	label := g.Params.ByName("label")

	value, exists := internal.ReadNumLabel(name, label)
	if !exists {
		g.Value(0) // Default to 0
	}
	g.Value(value)
}

// ReadStatName godoc
// @Summary Read a specific statistic
// @Schemes
// @Param name path string true "Stat name"
// @Description Read a specific statistic
// @Accept json
// @Produce json
// @Success 200 {object} JSONNumReadResult "desc"
// @Router /read/num/{name} [get]
func ReadStatName(g *gin.Context) {
	name := g.Params.ByName("name")

	values := internal.ReadNumName(name)
	g.JSON(http.StatusOK, values)
}
