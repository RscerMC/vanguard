package structs

import (
	"strings"

	"github.com/RscerMC/vanguard/config"
	"github.com/bwmarrin/discordgo"
)

type Vanguardcommand struct {
	Command         *discordgo.ApplicationCommand
	Usage           string
	DeveloperOnly   bool
	GuildOnly       bool
	UserPermissions []int64
	BotPermissions  []int64

	RunCMD func(s *discordgo.Session, i *discordgo.InteractionCreate) error
}

func (c *Vanguardcommand) Run(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if c.DeveloperOnly && i.Member.User.ID != config.DeveloperID {
		return errorEmbed(s, i, "You are not the developer")
	}

	if c.GuildOnly && i.GuildID == "" {
		return errorEmbed(s, i, "This command can only be used in a server")
	}

	if !c.HasPermission(i.Member) {
		missingPerms := FindMissingPermissions(i.Member.Permissions, c.UserPermissions)
		return errorEmbed(s, i, "You are missing permissions: "+PermissionsToName(missingPerms))
	}

	botUser, err := s.State.Member(i.GuildID, s.State.User.ID)
	if err != nil {
		return errorEmbed(s, i, "Failed to retrieve bot permissions")
	}

	if !c.HasPermission(botUser) {
		missingPerms := FindMissingPermissions(botUser.Permissions, c.BotPermissions)
		return errorEmbed(s, i, "I am missing permissions: "+PermissionsToName(missingPerms))
	}

	err = c.RunCMD(s, i)
	if err != nil {
		return errorEmbed(s, i, "An error occurred while running the command!")
	}

	return nil
}

type Permission struct {
	Name    string
	Bitmask int64
}

var Permissions = []Permission{
	{"Create Instant Invite", 0x0000000000000001},
	{"Kick Members", 0x0000000000000002},
	{"Ban Members", 0x0000000000000004},
	{"Administrator", 0x0000000000000008},
	{"Manage Channels", 0x0000000000000010},
	{"Manage Guild", 0x0000000000000020},
	{"Add Reactions", 0x0000000000000040},
	{"View Audit Log", 0x0000000000000080},
	{"Priority Speaker", 0x0000000000000100},
	{"Stream", 0x0000000000000200},
	{"View Channel", 0x0000000000000400},
	{"Send Messages", 0x0000000000000800},
	{"Send TTS Messages", 0x0000000000001000},
	{"Manage Messages", 0x0000000000002000},
	{"Embed Links", 0x0000000000004000},
	{"Attach Files", 0x0000000000008000},
	{"Read Message History", 0x0000000000010000},
	{"Mention Everyone", 0x0000000000020000},
	{"Use External Emojis", 0x0000000000040000},
	{"View Guild Insights", 0x0000000000080000},
	{"Connect", 0x0000000000100000},
	{"Speak", 0x0000000000200000},
	{"Mute Members", 0x0000000000400000},
	{"Deafen Members", 0x0000000000800000},
	{"Move Members", 0x0000000001000000},
	{"Use VAD", 0x0000000002000000},
	{"Change Nickname", 0x0000000004000000},
	{"Manage Nicknames", 0x0000000008000000},
	{"Manage Roles", 0x0000000010000000},
	{"Manage Webhooks", 0x0000000020000000},
	{"Manage Guild Expressions", 0x0000000040000000},
	{"Use Application Commands", 0x0000000080000000},
	{"Request To Speak", 0x0000000100000000},
	{"Manage Events", 0x0000000200000000},
	{"Manage Threads", 0x0000000400000000},
	{"Create Public Threads", 0x0000000800000000},
	{"Create Private Threads", 0x0000001000000000},
	{"Use External Stickers", 0x0000002000000000},
	{"Send Messages In Threads", 0x0000004000000000},
}

func HasPermission(user *discordgo.Member, permissions []int64) bool {
	for _, perm := range permissions {
		if !hasSinglePermission(user.Permissions, perm) {
			return false
		}
	}
	return true
}

func hasSinglePermission(userPermissions int64, permission int64) bool {
	return userPermissions&permission != 0
}

func PermissionsToName(perms []int64) string {
	var names []string
	for _, perm := range perms {
		name := findPermissionName(perm)
		names = append(names, name)
	}
	return strings.Join(names, ", ")
}

func findPermissionName(permission int64) string {
	for _, perm := range Permissions {
		if perm.Bitmask == permission {
			return perm.Name
		}
	}
	return "Unknown Permission"
}

func FindMissingPermissions(userPermissions int64, requiredPermissions []int64) []int64 {
	var missing []int64
	for _, perm := range requiredPermissions {
		if !hasSinglePermission(userPermissions, perm) {
			missing = append(missing, perm)
		}
	}
	return missing
}

func (c *Vanguardcommand) HasPermission(m *discordgo.Member) bool {
	if m == nil {
		return false
	}

	if len(c.UserPermissions) == 0 {
		return true
	}

	return HasPermission(m, c.UserPermissions)
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
