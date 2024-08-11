package structs

import (
	"vanguard/config"

	"github.com/bwmarrin/discordgo"
)

type Quixbocommand struct {
	Command       *discordgo.ApplicationCommand
	Usage         string
	DeveloperOnly bool
	GuildOnly     bool

	RunCMD func(s *discordgo.Session, i *discordgo.InteractionCreate) error
}

func (c *Quixbocommand) Run(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if c.DeveloperOnly && i.Member.User.ID != config.DeveloperID {
		return errorEmbed(s, i, "You are not the developer")
	}

	if c.GuildOnly && i.GuildID == "" {
		return errorEmbed(s, i, "This command can only be used in a server")
	}

	err := c.RunCMD(s, i)
	if err != nil {
		return errorEmbed(s, i, "An error occurred while running the command!")
	}

	return nil
}

func errorEmbed(s *discordgo.Session, i *discordgo.InteractionCreate, errorMessage string) error {
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Error",
					Description: errorMessage,
					Color:       config.ErrorColor,
				},
			},
		},
	})
}
