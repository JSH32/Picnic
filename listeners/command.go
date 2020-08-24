package listeners

import (
	"regexp"
	"strings"

	"github.com/Riku32/Picnic/handler/command"
	"github.com/Riku32/Picnic/javascript"
	"github.com/Riku32/Picnic/stdlib/discord"
	"github.com/Riku32/Picnic/stdlib/logger"
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Registry command.Handler
	Prefix   string
}

func (c Command) Handler(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Message.Author.ID == s.State.User.ID || e.Author.Bot {
		return
	}

	_, err := s.Channel(e.ChannelID)
	if err != nil {
		logger.Error("Failed getting discord channel from ID " + e.ChannelID + " : " + err.Error())
		return
	}

	if !strings.HasPrefix(e.Message.Content, c.Prefix) {
		return
	}

	re := regexp.MustCompile(`(?:[^\s"]+|"[^"]*")+`)
	contSplit := re.FindAllString(e.Message.Content, -1)
	for i, k := range contSplit {
		if strings.Contains(k, "\"") {
			contSplit[i] = strings.Replace(k, "\"", "", -1)
		}
	}

	invoke := contSplit[0][len(c.Prefix):]
	invoke = strings.ToLower(invoke)

	if command, ok := c.Registry.GetCommand(invoke); ok {
		runtime := javascript.NewVM()

		// set discord library object
		runtime.SetGlobal("discord", discord.NewDiscordBind(s))

		// channel object
		channel := discord.Channel{
			ID: e.ChannelID,
		}

		// author user object
		author := discord.User{
			ID: e.Author.ID,
		}

		// set command argument object
		runtime.SetGlobal("args", discord.Args{
			Author: author,
			Message: discord.Message{
				ID:      e.ID,
				Channel: channel,
				Author:  author,
				Content: e.Content,
			},
			Channel: channel,
			Args:    contSplit[1:],
		})

		// execute the command
		runtime.Execute(command.Command)
	}
}
