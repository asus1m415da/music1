package commands

import (
	"discord-go-music-bot/internal/discordutil"
	"discord-go-music-bot/internal/state"

	"github.com/bwmarrin/discordgo"
)

func PauseSong(ctx *state.Context) {
	if !discordutil.BotInChannel(ctx) {
		embed := &discordgo.MessageEmbed{
			Title:       "❌ Error",
			Description: "No estoy en un canal de voz.",
			Color:       0xe74c3c,
		}
		ctx.ReplyEmbed(embed)
		return
	}

	state.PauseMutex.Lock()
	currentState := state.Paused[ctx.GetGuildID()]
	state.Paused[ctx.GetGuildID()] = !currentState
	state.PauseMutex.Unlock()

	state.PauseChMutex.Lock()
	if ch, exists := state.PauseChs[ctx.GetGuildID()]; exists {
		select {
		case ch <- !currentState:
		default:
		}
	}
	state.PauseChMutex.Unlock()

	if currentState {
		embed := &discordgo.MessageEmbed{
			Title:       "▶️ Reanudado",
			Description: "Se ha reanudado la reproducción.",
			Color:       0x2ecc71,
		}
		ctx.ReplyEmbed(embed)
	} else {
		embed := &discordgo.MessageEmbed{
			Title:       "⏸️ Pausado",
			Description: "Se ha pausado la reproducción.",
			Color:       0xf39c12,
		}
		ctx.ReplyEmbed(embed)
	}
}
