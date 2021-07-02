package main

import (
	"facom-bot/internal/interaction"
	"facom-bot/internal/logger"
	"facom-bot/internal/util"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := util.GetEnv("DISCORD_BOT_TOKEN")
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		logger.Error.Fatalf("Failed to initialize bot: %s", err.Error())
	}

	session.AddHandler(func(session *discordgo.Session, _ *discordgo.Ready) {
		logger.Info.Println("Bot is ready to Go!")
	})

	err = session.Open()
	if err != nil {
		logger.Error.Fatalf("Failed to create a websocket connection to Discord: %s", err.Error())
	}
	defer session.Close()

	interaction.RegisterCommands(session)
	session.AddHandler(interaction.InteractionHandler)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interrupt
	logger.Info.Println("Shutting down bot...")
}
