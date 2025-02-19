package handlers

import "net/http"

// Generic data for any key-value pair
type T map[string]any

type response struct {
	status int
	others T
}

func new(statusCode int, others T) *response {
	return &response{
		status: statusCode,
		others: others,
	}
}

func (r *response) send() (int, T) {
	return r.status, r.others
}

func SuccessResponse(message string, data any) (int, T) {
	return new(http.StatusOK, T{
		"success": true,
		"message": message,
		"data":    data,
	}).send()
}

func SuccessMessageResponse(message string) (int, T) {
	return new(http.StatusOK, T{
		"success": true,
		"message": message,
	}).send()
}

func FailureMessageResponse(message string) (int, T) {
	return new(http.StatusOK, T{
		"success": false,
		"message": message,
	}).send()
}

func ServerErrorResponse(message string) (int, T) {
	return new(http.StatusInternalServerError, T{
		"success": false,
		"message": message,
	}).send()
}

func AuthenticationErrorResponse(message string) (int, T) {
	return new(http.StatusUnauthorized, T{
		"success": false,
		"message": message,
	}).send()
}
