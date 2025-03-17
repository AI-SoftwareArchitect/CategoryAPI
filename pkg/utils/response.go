package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response represents a standard API response
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SendSuccess sends a success response
func SendSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

// SendError sends an error response
func SendError(c *gin.Context, status int, message string) {
	c.JSON(status, Response{
		Status:  status,
		Message: message,
	})
}

// SendBadRequest sends a bad request response
func SendBadRequest(c *gin.Context, message string) {
	SendError(c, http.StatusBadRequest, message)
}

// SendNotFound sends a not found response
func SendNotFound(c *gin.Context, message string) {
	SendError(c, http.StatusNotFound, message)
}

// SendInternalServerError sends an internal server error response
func SendInternalServerError(c *gin.Context, message string) {
	SendError(c, http.StatusInternalServerError, message)
}
