package svc

type SvcParams map[string]interface{}

type SvcResults struct {
	Status  string      `json:"status"`
	Results interface{} `json:"results"`
}
