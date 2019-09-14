package pkg

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	ListID string `envconfig:"LIST_ID" required:"true"`
}

var Config *EnvConfig

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c := &EnvConfig{}
	err = envconfig.Process("", c)
	if err != nil {
		log.Fatal(err.Error())
	}

	Config = c
}
