package messages

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Ping(m *discordgo.MessageCreate) string {
	msg := fmt.Sprintf("pong %s", time.Since(m.Timestamp).String())
	return msg
}
