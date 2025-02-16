package api

type NumStatRegistration struct {
	Name  string  `json:"name"`
	Label string  `json:"label"`
	Value float64 `json:"value"`
}

type NumStatRetrieval struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

type JSONNumReadResult struct {
	StatName1 float64 `json:"stat_name1" `
	StatName2 float64 `json:"stat_name2" `
}
