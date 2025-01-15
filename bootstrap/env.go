package bootstrap

import (
	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv  string `mapstructure:"APP_ENV"`
	AppPort string `mapstructure:"APP_PORT"`
	DBHost  string `mapstructure:"DB_HOST"`
	DBPort  string `mapstructure:"DB_PORT"`
	DBUser  string `mapstructure:"DB_USER"`
	DBPass  string `mapstructure:"DB_PASSWORD"`
	DBName  string `mapstructure:"DB_NAME"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Error().Msgf("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Error().Msgf("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Error().Msg("The App is running in development env")
	}

	return &env
}
