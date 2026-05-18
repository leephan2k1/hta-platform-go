package initialize

import (
	"fmt"
	"hta-platform/global"
	"os"

	"github.com/spf13/viper"
)

// LoadConfig reads configuration from a specific .env file.
func LoadConfig() (err error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	// List of keys to bind for Unmarshal to work without a physical .env file
	keys := []string{
		"DB_USER", "DB_PASSWORD", "DB_HOST", "DB_PORT", "DB_NAME",
		"SERVER_PORT", "LOG_LEVEL", "MM_REFERER", "AUTH0_DOMAIN", "AUTH0_AUDIENCE",
	}
	for _, key := range keys {
		viper.BindEnv(key)
	}

	err = viper.ReadInConfig()
	if err != nil {
		// If the config file is not found, we ignore it and rely on Environment Variables
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No .env file found, relying on environment variables")
		} else if _, ok := err.(*os.PathError); ok {
			fmt.Println("No .env file found, relying on environment variables")
		} else {
			return fmt.Errorf("error reading config file: %w", err)
		}
	}

	var config global.Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return fmt.Errorf("unable to decode config into struct: %w", err)
	}

	global.ConfigValue = &config
	return
}
