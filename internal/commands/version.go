package commands

import (
	"discord-go-music-bot/internal/state"
	"discord-go-music-bot/internal/validation"

	"github.com/bwmarrin/discordgo"
)

func Version(ctx *state.Context) {
	if !validation.HasPermission(ctx, discordgo.PermissionAdministrator) {
		embed := &discordgo.MessageEmbed{
			Title:       "❌ Sin Permiso",
			Description: "No tienes permiso para usar este comando.",
			Color:       0xe74c3c,
		}
		ctx.ReplyEmbed(embed)
		return
	}
	embed := &discordgo.MessageEmbed{
		Title:       "📌 Versión del Bot",
		Description: "**Versión:** `" + state.GoSourceHash + "`",
		Color:       0x2ecc71,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Bot de Música Discord",
		},
	}
	ctx.ReplyEmbed(embed)
}
