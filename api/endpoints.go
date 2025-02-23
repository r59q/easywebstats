package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"r59q.com/easywebstats/internal"
)

// @BasePath /api/v1

// SetNumStat godoc
// @Summary Set numeric stat
// @Schemes
// @Param message body api.NumStatRegistration true "Stat and value to increase by"
// @Description Set a numeric stat
// @Accept json
// @Produce json
// @Success 200 {object} api.ValueResponse
// @Router /register/num/set [post]
func SetNumStat(g *gin.Context) {
	var registration NumStatRegistration
	err := g.BindJSON(&registration)
	if err != nil {
		g.JSON(http.StatusInternalServerError, "Failed to capture registration object. Is it formatted correctly?")
		return
	}
	value := internal.HandleSetNum(registration.Name, registration.Label, registration.Value)
	g.JSON(http.StatusOK, gin.H{"value": value})
}

// IncreaseNumStat godoc
// @Summary Increase a numeric stat
// @Schemes
// @Param message body api.NumStatRegistration true "Stat and value"
// @Description Increase a numeric stat by some amount
// @Accept json
// @Produce json
// @Success 200 {object} api.ValueResponse
// @Router /register/num/increase [post]
func IncreaseNumStat(g *gin.Context) {
	var registration NumStatRegistration
	err := g.BindJSON(&registration)
	if err != nil {
		g.JSON(http.StatusInternalServerError, "Failed to capture registration object. Is it formatted correctly?")
		return
	}
	value := internal.HandleIncreaseNum(registration.Name, registration.Label, registration.Value)
	g.JSON(http.StatusOK, gin.H{"value": value})
}

// DecreaseNumStat godoc
// @Summary Decrease a numeric stat
// @Schemes
// @Param message body api.NumStatRegistration true "Stat and value to decrease by"
// @Description Decrease a numeric stat by some amount
// @Accept json
// @Produce json
// @Success 200 {object} api.ValueResponse
// @Router /register/num/decrease [post]
func DecreaseNumStat(g *gin.Context) {
	var registration NumStatRegistration
	err := g.BindJSON(&registration)
	if err != nil {
		g.JSON(http.StatusInternalServerError, "Failed to capture registration object. Is it formatted correctly?")
		return
	}
	value := internal.HandleDecreaseNum(registration.Name, registration.Label, registration.Value)
	g.JSON(http.StatusOK, gin.H{"value": value})
}

// ReadStatNameLabel godoc
// @Summary Read a specific statistic
// @Schemes
// @Param name path string true "Stat name"
// @Param label path string true "Stat label"
// @Description Read a specific single numeric statistic by name and label
// @Accept json
// @Produce json
// @Success 200 {object} api.ValueResponse
// @Router /read/num/value/{name}/{label} [get]
func ReadStatNameLabel(g *gin.Context) {
	name := g.Params.ByName("name")
	label := g.Params.ByName("label")

	value := internal.ReadNumLabel(name, label)
	g.JSON(http.StatusOK, gin.H{"value": value})
}

// ReadStatName godoc
// @Summary Read a specific statistic
// @Schemes
// @Param name path string true "Stat name"
// @Description Read all labels associated with a stat by name
// @Accept json
// @Produce json
// @Success 200 {object} JSONNumReadResult
// @Router /read/num/value/{name} [get]
func ReadStatName(g *gin.Context) {
	name := g.Params.ByName("name")

	values := internal.ReadNumName(name)
	g.JSON(http.StatusOK, values)
}

// ReadNumNames godoc
// @Summary Read all stat names
// @Schemes
// @Description Read list of numeric stat names
// @Accept json
// @Produce json
// @Success 200 {object} JSONNumReadResult
// @Router /read/num/names [get]
func ReadNumNames(g *gin.Context) {
	values := internal.ReadNumNames()
	g.JSON(http.StatusOK, values)
}

// ReadStatNameLabelExponentialRate godoc
// @Summary Read an exponential *estimate* of a numeric rate of change
// @Schemes
// @Param name path string true "Stat name"
// @Param label path string true "Stat label"
// @Description Read rate estimate of a specific statistic. Not suitable for accuracy, but great for quickly comparing stats.
// @Accept json
// @Produce json
// @Success 200 {object} api.ValueResponse
// @Router /read/num/exponentialrate/{name}/{label} [get]
func ReadStatNameLabelExponentialRate(g *gin.Context) {
	name := g.Params.ByName("name")
	label := g.Params.ByName("label")

	value := internal.ReadNumExponentialRate(name, label)
	g.JSON(http.StatusOK, gin.H{"value": value})
}

// ReadStatNameExponentialRates godoc
// @Summary Read all rate estimates for a stat name
// @Schemes
// @Param name path string true "Stat name"
// @Description Read rate estimates of a specific statistic. Not suitable for accuracy, but great for quickly comparing stats.
// @Accept json
// @Produce json
// @Success 200 {object} api.ValueResponse
// @Router /read/num/exponentialrate/{name} [get]
func ReadStatNameExponentialRates(g *gin.Context) {
	name := g.Params.ByName("name")

	value := internal.ReadNumExponentialRates(name)
	g.JSON(http.StatusOK, gin.H{"value": value})
}

// ReadNumMean godoc
// @Summary Read the mean value for a name
// @Schemes
// @Param name path string true "Stat name"
// @Description Returns the mean for all values across the stat name
// @Accept json
// @Produce json
// @Success 200 {object} api.ValueResponse
// @Router /read/num/mean/{name} [get]
func ReadNumMean(g *gin.Context) {
	name := g.Params.ByName("name")

	value := internal.ReadNumMean(name)
	g.JSON(http.StatusOK, gin.H{"value": value})
}

func GetMetrics(g *gin.Context) {
	g.String(http.StatusOK, internal.GetMetrics())
}
