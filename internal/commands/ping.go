package commands

import (
        "discord-go-music-bot/internal/state"
        "github.com/bwmarrin/discordgo"
)

func Ping(ctx *state.Context) {
        embed := &discordgo.MessageEmbed{
                Title:       "🏓 Ping",
                Description: "¡Pong!",
                Color:       0x00ff00,
        }
        ctx.ReplyEmbed(embed)
}
