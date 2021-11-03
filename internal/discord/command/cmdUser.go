package command

import (
	"archeage/pkg/bond"
	"github.com/DevMyong/LittleJake/internal/discord/config"
	"strings"
)

type UserInfo Command

func NewUserInfo() *UserInfo {
	cfg, _ := config.ParseConfigFromJSONFile(config.FileName)
	return &UserInfo{
		invokes:         []string{"user", "u", "정보"},
		usage:           strings.Join(cfg.Usages["user"], "\n"),
		isAdminRequired: false,
		nArgs:       MinimumNArgs(1),
		validFlags: config.FlagPattern{
			"stat":    nil,
			"save":    nil,
			"load":    {"@int 1 100"}, // 백분률, 1~100 사이의 숫자
			"diff":    {"@int 1 100"}, // 백분률, 1~100 사이의 숫자
			"history": {"name", "server", "union", "expedition", "all"},
		},
	}
}
func (u *UserInfo) Invokes() []string {
	return u.invokes
}
func (u *UserInfo) Usage() string {
	return u.usage
}
func (u *UserInfo) Description() string {
	return u.description
}

func (u *UserInfo) IsAdminRequired() bool {
	return u.isAdminRequired
}
func (u *UserInfo) NArgs(args []string) error{
	return u.nArgs(args)
}

func (u *UserInfo) ValidFlags() config.FlagPattern {
	return u.validFlags
}
func (u *UserInfo) ValidArgs() config.ArgPattern {
	return u.validArgs
}
func (u *UserInfo) Exec(ctx *Context) (err error) {
	_, err = archeage.GetUserInfo(ctx.Args, ctx.FlagMap)
	if err != nil{
		return
	}


	return
}