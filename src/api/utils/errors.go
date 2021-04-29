package utils

import "net/http"

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	Astatus  int    `json:"status"`
	Amessage string `json:"message"`
	Aerror   string `json:"error,omitempty"`
}

func (e *apiError) Status() int {
	return e.Astatus
}

func (e *apiError) Message() string {
	return e.Amessage
}

func (e *apiError) Error() string {
	return e.Aerror
}

func NewInternalServerError(message string) ApiError {
	return &apiError{
		Amessage: message,
		Astatus:  http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) ApiError {
	return &apiError{
		Amessage: message,
		Astatus:  http.StatusNotFound,
	}
}

func NewBadRequest(message string) ApiError {
	return &apiError{
		Amessage: message,
		Astatus:  http.StatusBadRequest,
	}
}

func NewApiError(status int, message string) ApiError {
	return &apiError{
		Amessage: message,
		Astatus:  status,
	}
}
