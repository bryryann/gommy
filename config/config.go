package config

import (
	"errors"
	"os"
)

type Config struct {
	Token  string
	Prefix string
}

var (
	AppConfig Config
)

func ReadConfig() (*Config, error) {
	token, ok := os.LookupEnv("TOKEN")
	if !ok {
		return nil, errors.New("Couldn't find token.")

	}

	prefix, ok := os.LookupEnv("BOT_PREFIX")
	if !ok {
		return nil, errors.New("Couldn't find appropriate bot prefix.")
	}

	AppConfig.Token = token
	AppConfig.Prefix = prefix

	return &AppConfig, nil
}
