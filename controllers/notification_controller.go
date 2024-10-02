package controllers

import (
	"github.com/chamodshehanka/go-notification-api/handlers"
	"github.com/chamodshehanka/go-notification-api/internal/notification"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NotificationDto struct {
	Message     string `json:"message"`
	ServiceName string `json:"serviceName"`
}

// AddNotification godoc
// @Summary Add a new notification
// @Description Add a new notification
// @Tags notification
// @Accept  json
// @Produce  json
// @Param notification body NotificationDto true "Notification object"
// @Success 200 {string} string "Notification sent successfully"
// @Router /notification [post]
func AddNotification(c *gin.Context) {
	var addNotificationDto NotificationDto

	if err := c.ShouldBindJSON(&addNotificationDto); err != nil {
		handlers.HandleErrorResponse(c, "Invalid request body", err, http.StatusBadRequest)
		return
	}

	// Implement AddNotification logic here
	notifier := notification.InAppNotifier{}
	err := notifier.SendNotification("Hello, World!")
	if err != nil {
		handlers.HandleErrorResponse(c, "Error occurred while sending notification", err, http.StatusInternalServerError)
		return
	}

	handlers.HandleSuccessResponse(c, "Notification sent successfully", http.StatusOK)
}
