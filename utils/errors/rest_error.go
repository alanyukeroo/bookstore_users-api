package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
		Message: message,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Status:  http.StatusNotFound,
		Error:   "bad_request",
		Message: message,
	}
}
