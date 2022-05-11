package main

import (
	"fmt"
	"strings"

	"github.com/methosi/tigerhall-kittens/storage/postgres"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	envFilename = "env/config"
)

func main() {
	logger := logrus.New()
	config, err := newConfig(envFilename)
	if err != nil {
		logger.Fatalf("error loading configuration: %v", err)
	}

	if err := setupServices(logger, config); err != nil {
		logger.WithError(err).Fatalln("failed setting up services")
	}

}

func setupServices(logger *logrus.Logger, config *viper.Viper) error {
	_, err := postgres.NewStorageFromConfig(config)
	if err != nil {
		logger.WithError(err).Fatal("failed setting up storage")
		return fmt.Errorf("storage init: %v", err)
	}

	return nil

	// Setup handlers in here
	// router := mux.NewRouter()

}

func newConfig(filename string) (*viper.Viper, error) {
	config := viper.NewWithOptions(viper.EnvKeyReplacer(strings.NewReplacer(".", "_")))
	config.SetConfigFile(filename)
	config.SetConfigType("ini")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		return nil, err
	}
	return config, nil
}
