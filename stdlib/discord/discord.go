package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dop251/goja"
)

// NewDiscordBind : create new discord js binding
func NewDiscordBind(session *discordgo.Session, runtime *goja.Runtime) Discord {
	return Discord{
		session: session,
		runtime: runtime,
	}
}

// Discord : discord lib object
type Discord struct {
	session *discordgo.Session
	runtime *goja.Runtime
}

func (d Discord) SendMessage(channelid, content string) {
	d.session.ChannelMessageSend(channelid, content)
}

func (d Discord) SendEmbed(channelid string, embed Embed) {
	d.session.ChannelMessageSendEmbed(channelid, &discordgo.MessageEmbed{
		Title:       embed.Title,
		Description: embed.Description,
	})
}

func (c Discord) Embed(call goja.ConstructorCall) *goja.Object {
	call.This.Set("setTitle", func(new string) *goja.Object {
		call.This.Set("title", new)
		return call.This
	})

	call.This.Set("setDescription", func(new string) *goja.Object {
		call.This.Set("description", new)
		return call.This
	})

	return nil
}
