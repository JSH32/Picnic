package config

import (
	"fmt"
	"os"

	"github.com/Riku32/Picnic/stdlib/logger"
	"gopkg.in/yaml.v2"
)

// Config : config object
type Config struct {
	Token  string `yaml:"token"`
	Prefix string `yaml:"prefix"`
	ES6    bool   `yaml:"es6"`
}

// Load : load config file
func Load() Config {
	configf, err := os.Open("./config.yaml")
	if err != nil {
		logger.Panic(fmt.Sprintf("could not find config"))
	}
	defer configf.Close()

	var config Config

	err = yaml.NewDecoder(configf).Decode(&config)
	if err != nil {
		logger.Panic(fmt.Sprintf("invalid config file"))
	}

	return config
}
