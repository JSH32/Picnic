package discord

import "github.com/bwmarrin/discordgo"

// NewDiscordBind : create new discord js binding
func NewDiscordBind(session *discordgo.Session) Discord {
	return Discord{
		session: session,
	}
}

// Discord : discord lib object
type Discord struct {
	session *discordgo.Session
}

func (d Discord) SendMessage(channelid, content string) {
	d.session.ChannelMessageSend(channelid, content)
}
