package configs

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	DbHost       string `envconfig:"DB_HOST" default:"localhost"`
	DbPort       string `envconfig:"DB_PORT" default:"27017"`
	DbName       string `envconfig:"DB_NAME" default:"myDb"`
	DbCollection string `envconfig:"DB_COLLECTION" default:"UserInfo"`
	LogDebug     bool   `envconfig:"LOG_DEBUG" default:"true"`
}

var Conf Config

func Init() {
	if err := envconfig.Process("", &Conf); err != nil {
		log.Fatal("failed to load envconfig")
	}
}
