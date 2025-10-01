package commands

import (
	"discord-go-music-bot/internal/state"
	"github.com/bwmarrin/discordgo"
)

func Oss(ctx *state.Context) {
	repoURL := "https://github.com/H-Edward/Discord-Go-Music-Bot"

	embed := &discordgo.MessageEmbed{
		Title:       "ℹ️ Información del Bot",
		Description: "Bot de Música para Discord",
		Color:       0x9b59b6,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "👑 Dueño",
				Value:  "<@1404572152014962708> (Galaxy A06)",
				Inline: false,
			},
			{
				Name:   "📦 Código Fuente",
				Value:  "Este bot es de código abierto\n[Ver en GitHub](" + repoURL + ")",
				Inline: false,
			},
			{
				Name:   "🛠️ Tecnología",
				Value:  "Desarrollado en Go con discordgo",
				Inline: false,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Bot de Música Discord by Galaxy A06",
		},
	}
	ctx.ReplyEmbed(embed)
}
