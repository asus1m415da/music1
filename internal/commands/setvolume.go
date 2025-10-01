package commands

import (
	"discord-go-music-bot/internal/state"
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func SetVolume(ctx *state.Context) {
	volume := ctx.Arguments["level"]

	if len(volume) < 1 {
		CurrentVolume(ctx)
		return
	}

	newVolume, err := strconv.ParseFloat(volume, 64)
	if err != nil || newVolume < 0.0 || newVolume > 200.0 {
		embed := &discordgo.MessageEmbed{
			Title:       "❌ Error",
			Description: "Valor de volumen inválido. Por favor especifica un número entre 0 y 200.",
			Color:       0xe74c3c,
		}
		ctx.ReplyEmbed(embed)
		return
	}
	var preservedVolume float64 = newVolume
	newVolume = newVolume / 100.0

	state.VolumeMutex.Lock()
	if _, ok := state.Volume[ctx.GetGuildID()]; !ok {
		state.Volume[ctx.GetGuildID()] = 1.0
	}
	state.Volume[ctx.GetGuildID()] = newVolume
	state.VolumeMutex.Unlock()

	embed := &discordgo.MessageEmbed{
		Title:       "🔊 Volumen Ajustado",
		Description: fmt.Sprintf("El volumen se ha establecido a **%.1f%%**", preservedVolume),
		Color:       0x3498db,
	}
	ctx.ReplyEmbed(embed)
}
