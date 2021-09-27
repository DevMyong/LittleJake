package command

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type MwPermissions struct{}

func (mw *MwPermissions) Exec(ctx *Context, cmd Commander) (next bool, err error) {
	if !cmd.IsAdminRequired() {
		next = true
		return
	}

	defer func() {
		if !next && err != nil {
			_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
				fmt.Sprintf("%s 명령어를 사용할 권한이 없습니다.", cmd.Invokes()))
		}
	}()

	guild, err := ctx.Session.Guild(ctx.Message.GuildID)
	if err != nil {
		return
	}

	if guild.OwnerID == ctx.Message.Author.ID {
		next = true
	}

	roleMap := make(map[string]*discordgo.Role)
	for _, role := range guild.Roles {
		roleMap[role.ID] = role
	}

	for _, rID := range ctx.Message.Member.Roles {
		if role, ok := roleMap[rID]; ok && role.Permissions&discordgo.PermissionAdministrator > 0 {
			next = true
			break
		}
	}
	return
}
