package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Riku32/Picnic/handler/command"
	"github.com/Riku32/Picnic/listeners"
	"github.com/Riku32/Picnic/stdlib/logger"
	"github.com/Riku32/Picnic/util/config"
	"github.com/bwmarrin/discordgo"
)

func main() {
	conf := config.Load()

	session, err := discordgo.New("Bot " + conf.Token)
	if err != nil {
		logger.Error("could not log in to discord")
		return
	}

	handler := command.Loader()

	commandlistener := listeners.Command{
		Registry: handler,
		Prefix:   conf.Prefix,
	}

	session.AddHandler(commandlistener.Handler)

	session.Open()

	// Wait here until CTRL-C or other term signal is received.
	logger.Info("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	session.Close()
}
