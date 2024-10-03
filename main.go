package main

import (
	"fmt"
	_ "github.com/chamodshehanka/go-notification-api/docs"
	"github.com/chamodshehanka/go-notification-api/internal/configs"
	"github.com/chamodshehanka/go-notification-api/middlewares"
	"github.com/chamodshehanka/go-notification-api/routes"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

// @title			Go Notification API
// @version		1.0
// @description	Go Notification API
// @BasePath	/api/v1
func main() {
	r := gin.New()
	r.Use(middlewares.Logger())

	routes.SetupRoutes(r)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		fmt.Println("Shutting down application...")
		os.Exit(0)
	}()

	err := r.Run(":" + configs.GetConfig().Port)
	if err != nil {
		fmt.Println("Error starting the server: ", err.Error())
		return
	}
}
