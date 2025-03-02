package config

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func StartBot(token string) (string, *discordgo.Session) {
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	err = bot.Open()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Gommy awakes...")
	return bot.State.User.ID, bot
}
