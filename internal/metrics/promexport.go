package metrics

import (
	"r59q.com/easywebstats/internal/datastore"
	"strconv"
	"strings"
)

func GetPrometheusExport() string {
	numericStore := datastore.GetNumberStore()
	names := numericStore.GetNames()
	exports := []string{}

	for _, name := range names {
		exports = append(exports, getNumericStat(name))
	}

	return strings.Join(exports, "\n")
}

func getNumericStat(name string) string {
	numericStore := datastore.GetNumberStore()
	labels := numericStore.GetLabels(name)
	exports := []string{}

	for key, value := range labels {
		// ews_stat{name="name",label="label"} xx
		fVal := strconv.FormatFloat(value, 'f', -1, 64)
		export := "ews_stat{name=\"" + name + "\",label=\"" + key + "\"} " + fVal
		exports = append(exports, export)
	}
	return strings.Join(exports, "\n")
}
