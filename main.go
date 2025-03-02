package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bryryann/gommy/config"
	"github.com/bryryann/gommy/messages"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	Bot *discordgo.Session
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error getting .env file.")
	}
}

func main() {
	err := config.ReadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	config.AppConfig.BotID, Bot = config.StartBot(config.AppConfig.Token)

	Bot.AddHandler(messages.MessageHandler)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	Bot.Close()
}
