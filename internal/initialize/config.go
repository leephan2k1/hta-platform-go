package initialize

import (
	"fmt"
	"hta-platform/global"

	"github.com/spf13/viper"
)

// LoadConfig reads configuration from a specific .env file.
func LoadConfig() (err error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		// If the config file is not found, return a specific error
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return fmt.Errorf("config file not found: %w", err)
		}
		return fmt.Errorf("error reading config file: %w", err)
	}

	var config global.Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return fmt.Errorf("unable to decode config into struct: %w", err)
	}

	global.ConfigValue = &config
	return
}
