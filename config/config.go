package config

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		DB             DB
		Server         Server
		BreedValidator BreedValidator
	}

	DB struct {
		Address  string `envconfig:"DB_HOST" required:"true"`
		Port     string `envconfig:"DB_PORT" required:"true"`
		Name     string `envconfig:"DB_NAME" required:"true"`
		User     string `envconfig:"DB_USER" required:"true"`
		Password string `envconfig:"DB_PASSWORD" required:"true"`
	}

	Server struct {
		Port        string        `envconfig:"SERVER_PORT" required:"true"`
		ReadTimeout time.Duration `envconfig:"SERVER_READ_TIMEOUT" required:"true"`
		IdleTimeout time.Duration `envconfig:"SERVER_IDLE_TIMEOUT" required:"true"`
	}

	BreedValidator struct {
		TheCatApiKey   string        `envconfig:"THE_CAT_API_KEY" required:"true"`
		RequestTimeout time.Duration `envconfig:"THE_CAT_API_TIMEOUT" required:"true"`
	}
)

func New() (Config, error) {
	var c Config

	err := envconfig.Process("", &c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}

func Load(envPath ...string) error {
	return godotenv.Load(envPath...)
}
