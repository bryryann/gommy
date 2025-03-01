package main

import (
	"log"

	"github.com/bryryann/gommy/config"

	_ "github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error getting .env file.")
	}
}

func main() {
	_ = config.ReadConfig()
}
