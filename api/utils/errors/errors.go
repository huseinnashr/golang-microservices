package errors

import "net/http"

type ApiError interface {
	GetStatus() int
	GetMessage() string
	GetError() string
}

type apiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func (e *apiError) GetStatus() int {
	return e.Status
}

func (e *apiError) GetMessage() string {
	return e.Message
}

func (e *apiError) GetError() string {
	return e.Error
}

func NewApiError(statusCode int, message string) ApiError {
	return &apiError{Status: statusCode, Message: message}
}

func NewNotFoundApiError(message string) ApiError {
	return &apiError{
		Status:  http.StatusNotFound,
		Message: message,
	}
}

func NewInternalServerError(message string) ApiError {
	return &apiError{
		Status:  http.StatusInternalServerError,
		Message: message,
	}
}

func NewBadRequestError(message string) ApiError {
	return &apiError{
		Status:  http.StatusBadRequest,
		Message: message,
	}
}
