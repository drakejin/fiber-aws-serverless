package config

import (
	"time"

	"github.com/spf13/viper"
)

type DB struct {
	Host     string
	Port     string
	Dialect  string
	Schema   string
	Username string
	Password string
}

type Server struct {
	Port    string
	Timeout time.Duration
}

type Config struct {
	App        string
	Env        string
	Debug      bool
	Release    string

	HTTPServer *Server
	ServiceDB  *DB
}

var Release = ""

func New(useDotenv bool) (*Config, error) {
	if useDotenv {
		viper.AddConfigPath(".")
		viper.SetConfigFile(".env")
		err := viper.ReadInConfig()
		if err != nil {
			return nil, err
		}
	} else {
		viper.AutomaticEnv()
	}

	return &Config{
		App:        viper.GetString("FIBER_APP_APP"),
		Env:        viper.GetString("FIBER_APP_ENV"),
		Debug:      viper.GetBool("FIBER_APP_DEBUG"),
		Release:    Release,

		HTTPServer: &Server{
			Port:    viper.GetString("FIBER_APP_HTTP_SERVER_PORT"),
			Timeout: viper.GetDuration("FIBER_APP_HTTP_SERVER_TIMEOUT"),
		},
		ServiceDB: &DB{
			Host:     viper.GetString("FIBER_APP_DB_SERVICE_HOST"),
			Port:     viper.GetString("FIBER_APP_DB_SERVICE_PORT"),
			Dialect:  viper.GetString("FIBER_APP_DB_SERVICE_DIALECT"),
			Schema:   viper.GetString("FIBER_APP_DB_SERVICE_SCHEMA"),
			Username: viper.GetString("FIBER_APP_DB_SERVICE_USERNAME"),
			Password: viper.GetString("FIBER_APP_DB_SERVICE_PASSWORD"),
		},
	}, nil
}
