package command

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

// [prefix][cmd] [arg1] [arg2] [arg3] ....

type CommandHandler struct {
	prefix string

	cmdInstances []Commander
	cmdMap       map[string]Commander
	middlewares  []Middleware

	OnError func(err error, ctx *Context)
}

func NewCommandHandler(prefix string) *CommandHandler {
	return &CommandHandler{
		prefix:       prefix,
		cmdInstances: make([]Commander, 0),
		cmdMap:       make(map[string]Commander),
		middlewares:  make([]Middleware, 0),
		OnError:      func(error, *Context) {},
	}
}

func (c *CommandHandler) RegisterCommand(cmd Commander) {
	c.cmdInstances = append(c.cmdInstances, cmd)
	for _, invoke := range cmd.Invokes() {
		c.cmdMap[invoke] = cmd
	}
}
func (c *CommandHandler) RegisterMiddleware(mw Middleware) {
	c.middlewares = append(c.middlewares, mw)
}

func (c *CommandHandler) HandleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.Author.Bot || !strings.HasPrefix(m.Content, c.prefix) {
		return
	}

	split := strings.Split(m.Content[len(c.prefix):], " ")
	if len(split) < 1 {
		return
	}

	invoke := strings.ToLower(split[0])
	args := split[1:]

	cmd, ok := c.cmdMap[invoke]
	if !ok || cmd == nil {
		return
	}
	ctx := &Context{
		Session:   s,
		inputArgs: args,
		Handler:   c,
		Message:   m.Message,
	}

	for _, mw := range c.middlewares {
		next, err := mw.Exec(ctx, cmd)
		if err != nil {
			c.OnError(err, ctx)
			return
		}
		if !next {
			return
		}
	}
	if err := cmd.Exec(ctx); err != nil {
		c.OnError(err, ctx)
	}
}
