package commands

import (
	"discord-go-music-bot/internal/state"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func ShowQueue(ctx *state.Context) {
	state.QueueMutex.Lock()
	defer state.QueueMutex.Unlock()

	if len(state.Queue[ctx.GetGuildID()]) == 0 {
		embed := &discordgo.MessageEmbed{
			Title:       "📋 Cola Vacía",
			Description: "No hay canciones en la cola.",
			Color:       0x95a5a6,
		}
		ctx.ReplyEmbed(embed)
		return
	}

	var formattedQueue []string
	for i, song := range state.Queue[ctx.GetGuildID()] {
		formattedQueue = append(formattedQueue, fmt.Sprintf("`[%d]` %s", i+1, song))
	}

	embed := &discordgo.MessageEmbed{
		Title:       "📋 Cola de Reproducción",
		Description: strings.Join(formattedQueue, "\n"),
		Color:       0x9b59b6,
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("Total: %d canciones", len(state.Queue[ctx.GetGuildID()])),
		},
	}
	ctx.ReplyEmbed(embed)
}
