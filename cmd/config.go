package cmd

import (
	"fmt"
	"net/url"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ApiNinjaAPIKey  string  `envconfig:"API_NINJA_API_KEY" required:"true"`
	ApiNinjaBaseURL url.URL `envconfig:"API_NINJA_BASE_URL" required:"true"`
	AbstractAPIKey  string  `envconfig:"ABSTRACT_API_KEY"  required:"true"`
	AbstractBaseURL url.URL `envconfig:"ABSTRACT_BASE_URL" required:"true"`
	DadJokeBaseURL  url.URL `envconfig:"DAD_JOKE_BASE_URL" required:"true"`
	MeowFactBaseURL url.URL `envconfig:"MEOW_FACT_BASE_URL" required:"true"`
}

func ReadConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, fmt.Errorf("failed to load config: %w", err)
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, fmt.Errorf("failed to parse config: %w", err)
	}

	return cfg, nil
}
