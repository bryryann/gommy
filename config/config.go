package config

import "os"

type Config struct {
	Token  string
	Prefix string
}

var (
	AppConfig Config
)

func ReadConfig() *Config {

	token, ok := os.LookupEnv("TOKEN")
	if !ok {
		panic("Token mismatch.")
	}

	prefix, ok := os.LookupEnv("BOT_PREFIX")
	if !ok {
		panic("Empty prefix.")
	}

	AppConfig.Token = token
	AppConfig.Prefix = prefix

	return &AppConfig
}
