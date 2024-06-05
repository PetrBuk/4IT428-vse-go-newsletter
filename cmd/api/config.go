package main

import (
	"sync"

	envx "go.strv.io/env"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

const dotenvPath = ".env"

var (
	once sync.Once

	validate = validator.New()
)

type Config struct {
	Port        int    `env:"PORT" validate:"required"`
	AuthSecret  string `env:"AUTH_SECRET" validate:"required"`
	DatabaseURL string `env:"DATABASE_URL" validate:"required"`
	SupabaseID  string `env:"SUPABASE_ID" validate:"required"`
	SupabaseKey string `env:"SUPABASE_KEY" validate:"required"`
}

func LoadConfig() (Config, error) {
	loaddotenv(dotenvPath)

	cfg := Config{}
	if err := envx.Apply(&cfg); err != nil {
		return cfg, err
	}
	if err := validate.Struct(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

func loaddotenv(path string) {
	once.Do(func() {
		if path == "" {
			path = ".env"
		}

		_ = godotenv.Load(dotenvPath)
		_ = godotenv.Load(dotenvPath + ".common")
	})
}

func MustLoadConfig() Config {
	cfg, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
