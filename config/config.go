package config

import (
	"github.com/spf13/viper"
)

type BaseConfig struct {
	Title       string `mapstructure:"title"`
	Description string `mapstructure:"description"`
	BaseURL     string `mapstructure:"base_url"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type AppConfig struct {
	Seed            int `mapstructure:"seed"`
	RefreshInterval int `mapstructure:"refresh_interval"`
	CacheDuration   int `mapstructure:"cache_duration"`
}

type StaticConfig struct {
	MessagesDir string `mapstructure:"messages_dir"`
	FontsDir    string `mapstructure:"fonts_dir"`
	PublicDir   string `mapstructure:"public_dir"`
}

type Config struct {
	Base BaseConfig `mapstructure:"base"`

	Server ServerConfig `mapstructure:"server"`

	App AppConfig `mapstructure:"app"`

	Static StaticConfig `mapstructure:"static"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	viper.SetDefault("base.title", "WhoCares.io")
	viper.SetDefault("base.description", "Professional Silence Tracker")
	viper.SetDefault("base.base_url", "https://whocares.io")
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("app.seed", 8000000)
	viper.SetDefault("app.refresh_interval", 60)
	viper.SetDefault("app.cache_duration", 3600)
	viper.SetDefault("static.messages_dir", "assets/messages")
	viper.SetDefault("static.fonts_dir", "assets/fonts")
	viper.SetDefault("static.public_dir", "public")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("WHC")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
