package commands

import (
	"discord-go-music-bot/internal/discordutil"
	"discord-go-music-bot/internal/logging"
	"discord-go-music-bot/internal/state"
	"time"

	"github.com/bwmarrin/discordgo"
)

func StopSong(ctx *state.Context) {
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

	state.QueueMutex.Lock()
	state.Queue[ctx.GetGuildID()] = []string{}
	state.QueueMutex.Unlock()

	state.PlayingMutex.Lock()
	state.Playing[ctx.GetGuildID()] = false
	state.PlayingMutex.Unlock()

	go func() {
		time.Sleep(500 * time.Millisecond)
		vc.Speaking(false)
		err = vc.Disconnect()
		if err != nil {
			logging.ErrorLog("Error al desconectar del canal de voz: " + err.Error())
		}
	}()

	embed := &discordgo.MessageEmbed{
		Title:       "⏹️ Detenido",
		Description: "Se ha detenido la reproducción y se ha limpiado la cola.",
		Color:       0xe74c3c,
	}
	ctx.ReplyEmbed(embed)
}
