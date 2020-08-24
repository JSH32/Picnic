package discord

import "github.com/bwmarrin/discordgo"

// NewDiscordBind : create new discord js binding
func NewDiscordBind(session *discordgo.Session) Discord {
	return Discord{
		session: session,
		SendMessage: func(channelid, content string) {
			session.ChannelMessageSend(channelid, content)
		},
	}
}

// Discord : discord lib object
type Discord struct {
	session     *discordgo.Session
	SendMessage func(string, string) `json:"sendMessage"`
}
