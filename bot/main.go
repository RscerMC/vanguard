package bot

import (
	"vanguard/config"

	"github.com/bwmarrin/discordgo"
)

var (
	Session *discordgo.Session
)
//Main bot operations
func init() {
	bot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		panic(err)
	}

	Session = bot

	Session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	Session.State.MaxMessageCount = 50
}
