package controllers

import (
	"github.com/chamodshehanka/go-notification-api/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Healthz godoc
// @Summary Health check
// @Description Check if the service is healthy
// @Tags healthz
// @Produce json
// @Success 200 {string} string "Go Notification API is healthy"
// @Router /notification/healthz [get]
func Healthz(c *gin.Context) {
	handlers.HandleSuccessResponse(c, "Go Notification API is healthy", http.StatusOK)
}
