package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	TGToken string `env:"TG_TOKEN" envDefault:"test123"`
}

func NewConfig() Config {
	var cfg Config

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file", err.Error())
	}

	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	return cfg
}
