package commands

import (
	"discord-go-music-bot/internal/state"
	"github.com/bwmarrin/discordgo"
)

func Pong(ctx *state.Context) {
	embed := &discordgo.MessageEmbed{
		Title:       "🏓 Pong",
		Description: "¡Ping!",
		Color:       0x00ff00,
	}
	ctx.ReplyEmbed(embed)
}
