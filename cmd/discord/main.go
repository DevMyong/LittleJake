package main

import (
	"fmt"
	"github.com/DevMyong/LittleJake/internal/discord/command"
	"github.com/DevMyong/LittleJake/internal/discord/config"
	"github.com/DevMyong/LittleJake/internal/discord/events"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	const fileName = "C:/Users/KD/Projects/LittleJake/config/config.json"

	cfg, err := config.ParseConfigFromJSONFile(fileName)
	if err != nil {
		panic(err)
	}

	s, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		panic(err)
	}

	s.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	registerEvents(s)
	registerCommands(s, cfg)

	if err = s.Open(); err != nil {
		panic(err)
	}

	fmt.Println("Bot is running, Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	s.Close()
}

func registerEvents(s *discordgo.Session) {
	s.AddHandler(events.NewMessageHandler().Handler)
}

func registerCommands(s *discordgo.Session, cfg *config.Config) {
	cmdHandler := command.NewCommandHandler(cfg.Prefix)
	cmdHandler.OnError = func(err error, ctx *command.Context) {
		_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
			fmt.Sprintf("Command Execution failed: %s", err.Error()))
	}

	cmdHandler.RegisterCommand(command.NewUserInfo())
	cmdHandler.RegisterMiddleware(&command.MwArguments{})
	cmdHandler.RegisterMiddleware(&command.MwPermissions{})
	s.AddHandler(cmdHandler.HandleMessage)
}
