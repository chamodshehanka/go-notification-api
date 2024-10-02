package configs

type NotificationAPIConfig struct {
	ClientID     string
	ClientSecret string
}

type Config struct {
	Port                  string
	APIKey                string
	NotificationAPIConfig NotificationAPIConfig
}
