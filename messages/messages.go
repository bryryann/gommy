package messages

import (
	"log"
	"unicode/utf8"

	"github.com/bryryann/gommy/config"
	"github.com/bwmarrin/discordgo"
)

type handler func(*discordgo.MessageCreate) string

var messages map[string]handler = make(map[string]handler)

func StrapMessage(message string, h handler) {
	messages[config.AppConfig.Prefix+message] = h
}

func InitHandler(sess *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == config.AppConfig.BotID {
		return
	}

	if utf8.RuneCountInString(msg.Content) < 4 || msg.Content[:4] != config.AppConfig.Prefix {
		return
	}

	handler, exists := messages[msg.Content]
	if exists {
		_, err := sess.ChannelMessageSend(msg.ChannelID, handler(msg))
		if err != nil {
			log.Fatal(err)
		}
	}

	if msg.Content == config.AppConfig.Prefix+"ping" {
	}
}
