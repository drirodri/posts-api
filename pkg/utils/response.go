package utils
import (
	"encoding/json"
	"net/http"
)
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
}
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
type ValidationErrors struct {
	Errors []ValidationError `json:"validation_errors"`
}
func WriteJSONResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
func WriteSuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	response := APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
	WriteJSONResponse(w, statusCode, response)
}
func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string, errorCode string, details string) {
	response := APIResponse{
		Success: false,
		Message: message,
		Error: &APIError{
			Code:    errorCode,
			Message: message,
			Details: details,
		},
	}
	WriteJSONResponse(w, statusCode, response)
}
func WriteValidationErrorResponse(w http.ResponseWriter, validationErrors []ValidationError) {
	response := APIResponse{
		Success: false,
		Message: "Validation failed",
		Error: &APIError{
			Code:    "VALIDATION_ERROR",
			Message: "Request validation failed",
		},
		Data: ValidationErrors{Errors: validationErrors},
	}
	WriteJSONResponse(w, http.StatusBadRequest, response)
}
func WriteNotFoundResponse(w http.ResponseWriter, resource string) {
	WriteErrorResponse(w, http.StatusNotFound,
		resource+" not found",
		"NOT_FOUND",
		"The requested "+resource+" does not exist")
}
func WriteInternalErrorResponse(w http.ResponseWriter, err error) {
	WriteErrorResponse(w, http.StatusInternalServerError,
		"Internal server error",
		"INTERNAL_ERROR",
		err.Error())
}
