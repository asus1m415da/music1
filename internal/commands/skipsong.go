package commands

import (
	"discord-go-music-bot/internal/discordutil"
	"discord-go-music-bot/internal/state"

	"github.com/bwmarrin/discordgo"
)

func SkipSong(ctx *state.Context) {
	vc, err := discordutil.GetVoiceConnection(ctx)
	if err != nil {
		embed := &discordgo.MessageEmbed{
			Title:       "❌ Error",
			Description: "No estoy en un canal de voz",
			Color:       0xe74c3c,
		}
		ctx.ReplyEmbed(embed)
		return
	}

	state.StopMutex.Lock()
	if stopChan, exists := state.StopChannels[ctx.GetGuildID()]; exists {
		close(stopChan)
		delete(state.StopChannels, ctx.GetGuildID())
	}
	state.StopMutex.Unlock()

	vc.Speaking(false)

	embed := &discordgo.MessageEmbed{
		Title:       "⏭️ Saltando",
		Description: "Saltando a la siguiente canción...",
		Color:       0x3498db,
	}
	ctx.ReplyEmbed(embed)
}
