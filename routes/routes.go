package routes

import (
	"github.com/chamodshehanka/go-notification-api/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		notification := v1.Group("/notification")
		notification.POST("", controllers.AddNotification)
		notification.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		notification.GET("/healthz", controllers.Healthz)
	}
}
