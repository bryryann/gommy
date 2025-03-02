package messages

import (
	"log"
	"time"

	"github.com/bryryann/gommy/config"
	"github.com/bwmarrin/discordgo"
)

func MessageHandler(sess *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == config.AppConfig.BotID {
		return
	}

	if msg.Content[:4] != config.AppConfig.Prefix {
		return
	}

	if msg.Content == config.AppConfig.Prefix+"ping" {
		_, err := sess.ChannelMessageSend(msg.ChannelID, "pong"+time.Since(msg.Timestamp).String())

		if err != nil {
			log.Fatal(err)
		}
	}
}
