package command

import (
	"github.com/bwmarrin/discordgo"
)

// [prefix][cmd] [arg1] [arg2] [arg3] ....

type Context struct {
	Session   *discordgo.Session
	Message   *discordgo.Message
	inputArgs []string
	Handler   *CommandHandler

	// Args is actual Arguments parsed from inputArgs
	Args []string
	// FlagMap is actual Flags with parameter parsed from inputArgs
	FlagMap map[string]string
}
