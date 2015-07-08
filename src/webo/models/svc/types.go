package svc

type SvcParams map[string]interface{}

type SvcResults struct {
	Status  int         `json:"status"`
	Results interface{} `json:"results"`
}
