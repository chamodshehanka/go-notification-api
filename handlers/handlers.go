package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func HandleSuccessWithDataResponse(c *gin.Context, message string, data interface{}, statusCode int) {
	response := Response{
		Success: true,
		Message: message,
		Data:    data,
	}

	c.JSON(statusCode, response)
}

func HandleSuccessResponse(c *gin.Context, message string, statusCode int) {
	response := Response{
		Success: true,
		Message: message,
	}

	c.JSON(statusCode, response)
}

func HandleErrorResponse(c *gin.Context, message string, error error, statusCode int) {
	logrus.Errorf("Error: %v", error)
	response := Response{
		Success: false,
		Message: message,
		Error:   error.Error(),
	}

	c.JSON(statusCode, response)
}

type UnauthorizedError struct{}

func (e *UnauthorizedError) Error() string {
	return "Unauthorized"
}

func NewUnauthorizedError() *UnauthorizedError {
	return &UnauthorizedError{}
}

func HandleUnauthorisedResponse(c *gin.Context) {
	response := Response{
		Success: false,
		Message: "Unauthorised",
		Error:   NewUnauthorizedError().Error(),
	}

	c.JSON(http.StatusUnauthorized, response)
}
