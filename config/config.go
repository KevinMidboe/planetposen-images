// Package config handles environment variables.
package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

// Config contains environment variables.
type Config struct {
	Hostname     					string  `envconfig:"HOSTNAME"`
	GCSBucket 						string  `envconfig:"GCS_BUCKET" default:"p"`
	Port                  string  `envconfig:"PORT" default:"8000"`
}

// LoadConfig reads environment variables, populates and returns Config.
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Info("No .env file found")
	}

	var c Config

	err := envconfig.Process("", &c)

	return &c, err
}