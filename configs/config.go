package configs

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
}

var Conf Config

func Init() {
	if err := envconfig.Process("", &Conf); err != nil {
		log.Fatal("failed to load envconfig")
	}
}
