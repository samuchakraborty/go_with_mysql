package model

import "time"

type Data struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"msg"`
	Count	int `json:count`
	Data       []Customers `json:"customerData"`
}
type InsertData struct {
	StatusCode int       `json:"statusCode"`
	Message    string    `json:"msg"`
	Data       Customers `json:"customer"`
}

type ErrorResponse struct {
	StatusCode int       `json:"statusCode"`
	Message    string    `json:"msg"`
	Data       time.Time `json:"timestamp"`
}
