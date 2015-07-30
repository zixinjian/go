package svc

type Params map[string]interface{}

type Results struct {
	Status  string      	`json:"status"`
	Results interface{} 	`json:"results"`
}
