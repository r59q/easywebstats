package metrics

import (
	"fmt"
	"r59q.com/easywebstats/internal/datastore"
	"strconv"
	"strings"
)

func GetPrometheusExport() string {
	numericStore := datastore.GetNumberStore()
	names := numericStore.GetNames()

	allStats := []string{}
	allMeans := []string{}

	for _, name := range names {
		mean, stats := getNumericStat(name)
		allMeans = append(allMeans, mean)
		allStats = append(allStats, strings.Join(stats, "\n"))
	}

	statsExport := strings.Join(allStats, "\n")
	meansExport := strings.Join(allMeans, "\n")
	return fmt.Sprintf("%s\n\n%s", statsExport, meansExport)
}

func getNumericStat(name string) (string, []string) {
	return getNumericStatMeanExport(name), getNumericStatValues(name)
}

func getNumericStatValues(name string) []string {
	numericStore := datastore.GetNumberStore()
	labels := numericStore.GetLabels(name)

	exports := []string{}
	for key, value := range labels {
		fVal := strconv.FormatFloat(value, 'f', -1, 64)
		export := "ews_stat{name=\"" + name + "\",label=\"" + key + "\"} " + fVal
		exports = append(exports, export)
	}

	return exports
}

func getNumericStatMeanExport(name string) string {
	numericStore := datastore.GetNumberStore()
	mean := numericStore.GetMean(name)
	fMean := strconv.FormatFloat(mean, 'f', -1, 64)
	return "ews_stat_mean{name=\"" + name + "\"} " + fMean
}
