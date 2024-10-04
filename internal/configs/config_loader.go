package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var config Config

var requiredEnvs = [...]string{
	// Server
	"PORT",
	//"API_KEY",

	// Notification API
	"NOTIFICATION_API_CLIENT_ID",
	"NOTIFICATION_API_CLIENT_SECRET",

	// Knock API
	"KNOCK_API_KEY",
}

func init() {
	loadEnvFileIfAvailable()

	if err := ensureRequiredEnvsAreAvailable(); err != nil {
		logrus.Fatalf("Error loading environment variables: %v", err)
	}

	config = Config{
		Port:   getEnv("PORT"),
		APIKey: getEnv("API_KEY"),
		NotificationAPIConfig: NotificationAPIConfig{
			ClientID:     getEnv("NOTIFICATION_API_CLIENT_ID"),
			ClientSecret: getEnv("NOTIFICATION_API_CLIENT_SECRET"),
		},
	}

	logrus.Println("Config loaded successfully!")
}

func GetConfig() *Config {
	return &config
}

func ensureRequiredEnvsAreAvailable() error {
	for _, env := range requiredEnvs {
		if getEnv(env) == "" {
			return fmt.Errorf("fatal: required environment variable '%s' not found", env)
		}
	}
	return nil
}

func loadEnvFileIfAvailable() {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			logrus.Fatalf("Error loading .env file: %v", err)
		}
	} else if !os.IsNotExist(err) {
		logrus.Fatalf("Error checking .env file: %v", err)
	}
}

func getEnv(key string) string {
	return os.Getenv(key)
}

func getIntegerEnv(key string) int {
	envStr := getEnv(key)

	intEnv, err := strconv.Atoi(envStr)
	if err != nil {
		logrus.Errorf("Error converting environment variable '%s' to integer: %v", key, err)
		return 0
	}

	return intEnv
}
