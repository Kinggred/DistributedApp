package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DEBUG        bool
	VERSION_MAIN int8
	VERSION_MIN  int8
	SERVER_ONLY  bool
}

var CONFIG Config

func LoadVariables() {
	file, err := os.Open("./core/config/config.yaml")
	if err != nil {
		log.Fatalf("Variables are fucked, bye")
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&CONFIG); err != nil {
		log.Fatalf("Variables are fucked (file exists, cannot be decoded tho), bye")
	}
}
