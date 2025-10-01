package commands

import (
	"discord-go-music-bot/internal/state"
	"os"

	"github.com/bwmarrin/discordgo"
)

func Unknown(ctx *state.Context) {
	unknown_commands := os.Getenv("UNKNOWN_COMMANDS")
	switch unknown_commands {
	case "help":
		Help(ctx)
	case "error":
		embed := &discordgo.MessageEmbed{
			Title:       "❌ Comando Desconocido",
			Description: "Comando desconocido. Escribe /ayuda para ver la lista de comandos.",
			Color:       0xe74c3c,
		}
		ctx.ReplyEmbed(embed)
	default:
		return
	}
}
