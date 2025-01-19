package bootstrap

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Env struct {
	AppEnv    string `mapstructure:"APP_ENV"`
	AppPort   string `mapstructure:"APP_PORT"`
	DBHost    string `mapstructure:"DB_HOST"`
	DBPort    string `mapstructure:"DB_PORT"`
	DBUser    string `mapstructure:"DB_USER"`
	DBPass    string `mapstructure:"DB_PASSWORD"`
	DBName    string `mapstructure:"DB_NAME"`
	SecretKey string `mapstructure:"SECRET_KEY"`
}

func NewEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Error().Msgf("Error loading .env file: %v", err)
	}

	env := Env{
		AppEnv:    os.Getenv("APP_ENV"),
		AppPort:   os.Getenv("APP_PORT"),
		DBHost:    os.Getenv("DB_HOST"),
		DBPort:    os.Getenv("DB_PORT"),
		DBUser:    os.Getenv("DB_USER"),
		DBPass:    os.Getenv("DB_PASSWORD"),
		DBName:    os.Getenv("DB_NAME"),
		SecretKey: os.Getenv("SECRET_KEY"),
	}

	if env.AppEnv == "development" {
		log.Info().Msg("The App is running in development env")
	}

	return &env
}
