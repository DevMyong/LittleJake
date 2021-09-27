package command

type Middleware interface {
	Exec(ctx *Context, cmd Commander) (next bool, err error)
}
