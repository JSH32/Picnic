package listeners

import (
	"regexp"
	"strings"

	"github.com/Riku32/Picnic/discord"
	"github.com/Riku32/Picnic/handler/command"
	"github.com/Riku32/Picnic/javascript"
	"github.com/Riku32/Picnic/logger"
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Registry command.Handler
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

	prefix := "p!"

	if !strings.HasPrefix(e.Message.Content, prefix) {
		return
	}

	re := regexp.MustCompile(`(?:[^\s"]+|"[^"]*")+`)
	contSplit := re.FindAllString(e.Message.Content, -1)
	for i, k := range contSplit {
		if strings.Contains(k, "\"") {
			contSplit[i] = strings.Replace(k, "\"", "", -1)
		}
	}

	invoke := contSplit[0][len(prefix):]
	invoke = strings.ToLower(invoke)

	if command, ok := c.Registry.GetCommand(invoke); ok {
		runtime := javascript.NewVM()

		runtime.SetGlobal("discord", discord.NewDiscordBind(s))

		runtime.SetGlobal("message", discord.Message{
			UserID:    e.Author.ID,
			MessageID: e.Message.ID,
			ChannelID: e.ChannelID,
			Content:   e.Content,
			Args:      contSplit[1:],
		})
		runtime.Execute(command.Command)
	}
}
