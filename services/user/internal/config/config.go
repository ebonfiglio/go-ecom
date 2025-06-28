package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Auth0Config struct {
	Domain   string `mapstructure:"domain"`
	Audience string `mapstructure:"audience"`
	ClientID string `mapstructure:"clientId"`
}

type Config struct {
	Port  string      `mapstructure:"port"`
	Auth0 Auth0Config `mapstructure:"auth0"`
}

func Load() (*Config, error) {
	v := viper.New()

	v.AddConfigPath("./configs")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	// Uncomment to use different configs per environment
	// v.SetDefault("environment", "development")
	// v.BindEnv("environment", "APP_ENV")
	// env := v.GetString("environment")
	// // merge in environment-specific overrides
	// v.SetConfigName(fmt.Sprintf("config.%s", env))

	if err := v.MergeInConfig(); err != nil {
		return nil, fmt.Errorf("loading base config: %w", err)
	}

	v.SetDefault("app.port", "8080")
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()

	var cfg struct {
		App struct {
			Port string `mapstructure:"port"`
		} `mapstructure:"app"`
		Auth0 Auth0Config `mapstructure:"auth0"`
	}
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return &Config{
		Port:  cfg.App.Port,
		Auth0: cfg.Auth0,
	}, nil
}
