package utils

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// -----------------------------
// Standard Response Structures
// -----------------------------

type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
}

type APIError struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// ------------------------
// Helper for Error Struct
// ------------------------

func NewAPIError(err error, code string) *APIError {
	switch v := err.(type) {
	case *url.Error:
		return &APIError{
			Code:    code,
			Message: "URL Error",
			Details: v.Err.Error(),
		}
	default:
		return &APIError{
			Code:    code,
			Message: err.Error(),
		}
	}
}

// -----------------------------
// Generic JSON Response Writer
// -----------------------------

func Respond(c *gin.Context, statusCode int, message string, data interface{}, err *APIError, token string) {
	if token != "" {
		c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	response := APIResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
		Error:   err,
	}

	c.JSON(statusCode, response)
}

// -----------------------------
// Predefined Common Responses
// -----------------------------

func RespondOK(c *gin.Context, data interface{}, message, token string) {
	Respond(c, http.StatusOK, message, data, nil, token)
}

func RespondCreated(c *gin.Context, data interface{}, message, token string) {
	Respond(c, http.StatusCreated, message, data, nil, token)
}

func RespondBadRequest(c *gin.Context, err error, token string) {
	Respond(c, http.StatusBadRequest, "Bad Request", nil, NewAPIError(err, "BAD_REQUEST"), token)
}

func RespondUnauthorized(c *gin.Context, err error) {
	Respond(c, http.StatusUnauthorized, "Unauthorized", nil, NewAPIError(err, "UNAUTHORIZED"), "")
}

func RespondForbidden(c *gin.Context, err error) {
	Respond(c, http.StatusForbidden, "Forbidden", nil, NewAPIError(err, "FORBIDDEN"), "")
}

func RespondNotFound(c *gin.Context, err error, token string) {
	Respond(c, http.StatusNotFound, "Not Found", nil, NewAPIError(err, "NOT_FOUND"), token)
}

func RespondInternalError(c *gin.Context, err error, token string) {
	Respond(c, http.StatusInternalServerError, "Internal Server Error", nil, NewAPIError(err, "INTERNAL_ERROR"), token)
}

func RespondRedirect(c *gin.Context, location, token string) {
	if token != "" {
		c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	c.Redirect(http.StatusMovedPermanently, location)
}
