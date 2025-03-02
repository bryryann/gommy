package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bryryann/gommy/config"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	BotId string
	Bot   *discordgo.Session
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error getting .env file.")
	}
}

func main() {
	appConfig, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	BotId, Bot = config.StartBot(appConfig.Token)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	Bot.Close()
}
