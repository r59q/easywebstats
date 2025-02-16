package api

type NumStatRegistration struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Label  string  `json:"label"`
	Action string  `json:"action"`
	Value  float64 `json:"value"`
}
