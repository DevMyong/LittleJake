package command

import "github.com/DevMyong/LittleJake/internal/discord/config"

type Commander interface {
	Invokes() []string
	Usage() string
	Description() string

	IsAdminRequired() bool

	NArgs([]string) error
	ValidArgs() config.ArgPattern
	ValidFlags() config.FlagPattern

	Exec(ctx *Context) error
}

type Command struct {
	invokes     []string
	usage       string
	description string

	isAdminRequired bool
	nArgs           PositionalArgs

	args       []string
	flagMap    map[string]string
	validArgs  config.ArgPattern
	validFlags config.FlagPattern
}
type PositionalArgs func(args []string) error
