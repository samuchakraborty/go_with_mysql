package model


type Data struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"msg"`
	Data       []Customers `json:"customerData"`
}