package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type (
	// Env is list all of env.
	Env struct {
		PostgresHostname string `mapstructure:"POSTGRES_HOSTNAME"`
		PostgresSsl      string `mapstructure:"POSTGRES_SSL"`
		PostgresUser     string `mapstructure:"POSTGRES_USER"`
		PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
		PostgresDB       string `mapstructure:"POSTGRES_DB"`
		PostgresPort     int32  `mapstructure:"POSTGRES_PORT"`
	}
)

func SetupEnv() (*Env, error) {
	env := &Env{}
	envBase := "local"
	mode := os.Getenv("MODE")
	if mode != "" {
		envBase = mode
	}

	viper.AddConfigPath(".")
	viper.SetConfigName(envBase)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal(err)
	}

	return env, nil
}
