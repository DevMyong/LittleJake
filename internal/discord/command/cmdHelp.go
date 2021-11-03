package command

import (
	"github.com/DevMyong/LittleJake/internal/config"
	"strings"
)

type Help Command

func NewHelp() *Help {
	cfg, _ := config.ParseConfigFromJSONFile(config.FileName)
	return &Help{
		invokes:         []string{"help", "h"},
		usage:           strings.Join(cfg.Usages["help"], "\n"),
		description:     "It shows how to use the Jake and command.",
		isAdminRequired: false,
		validArgs: []string{"help", "user"},
	}
}
