package global

import (
	"go.uber.org/zap"
)

// Config stores all configuration of the application.
type Config struct {
	DBUser     string `mapstructure:"DB_USER"`
	DBPass     string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	LogLevel   string `mapstructure:"LOG_LEVEL"`
	MMReferer  string `mapstructure:"MM_REFERER"`
}

var (
	Logger *zap.Logger
	ConfigValue *Config
)
