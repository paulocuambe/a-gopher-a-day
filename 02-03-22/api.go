package main

import "time"

type api struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Author  string `json:"author"`
}

func NewApi() *api {
	return &api{Name: "Simple Api", Version: "0.0.1", Author: "paulocuambe"}
}

type ApiError struct {
	Message   string    `json:"message"`
	Error     string    `json:"error"`
	Status    int       `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

func NewApiError(title string, err error, code int) *ApiError {
	return &ApiError{Message: err.Error(), Error: title, Status: code, Timestamp: time.Now()}
}
