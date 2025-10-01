package commands

import (
	"discord-go-music-bot/internal/state"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func CurrentVolume(ctx *state.Context) {
	state.VolumeMutex.Lock()
	volume, exists := state.Volume[ctx.GetGuildID()]
	if !exists {
		volume = 1.0
	}
	state.VolumeMutex.Unlock()

	volumePercentage := volume * 100.0

	embed := &discordgo.MessageEmbed{
		Title:       "🔊 Volumen Actual",
		Description: fmt.Sprintf("El volumen actual es **%.1f%%**", volumePercentage),
		Color:       0x3498db,
	}
	ctx.ReplyEmbed(embed)
}
