package commands

import (
	"discord-go-music-bot/internal/state"
	"discord-go-music-bot/internal/validation"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func NukeMessages(ctx *state.Context) {
	if !validation.HasPermission(ctx, discordgo.PermissionManageMessages) {
		embed := &discordgo.MessageEmbed{
			Title:       "❌ Sin Permiso",
			Description: "No tienes permiso para usar este comando.",
			Color:       0xe74c3c,
		}
		ctx.ReplyEmbed(embed)
		return
	}

	if ctx.Arguments["count"] == "" {
		embed := &discordgo.MessageEmbed{
			Title:       "❌ Error",
			Description: "Uso: /limpiar <número de mensajes>",
			Color:       0xe74c3c,
		}
		ctx.ReplyEmbed(embed)
		return
	}
	num, err := strconv.Atoi(ctx.Arguments["count"])
	if err != nil {
		embed := &discordgo.MessageEmbed{
			Title:       "❌ Error",
			Description: "Número de mensajes inválido",
			Color:       0xe74c3c,
		}
		ctx.ReplyEmbed(embed)
		return
	}
	if num < 1 || num > 100 {
		embed := &discordgo.MessageEmbed{
			Title:       "❌ Error",
			Description: "Por favor especifica un número entre 1 y 100",
			Color:       0xe74c3c,
		}
		ctx.ReplyEmbed(embed)
		return
	}
	num++

	messages, err := ctx.GetSession().ChannelMessages(ctx.GetChannelID(), num, "", "", "")
	if err != nil {
		embed := &discordgo.MessageEmbed{
			Title:       "❌ Error",
			Description: "Error al obtener mensajes",
			Color:       0xe74c3c,
		}
		ctx.ReplyEmbed(embed)
		return
	}
	for _, message := range messages {
		ctx.GetSession().ChannelMessageDelete(ctx.GetChannelID(), message.ID)
		time.Sleep(20 * time.Millisecond)
	}
	embed := &discordgo.MessageEmbed{
		Title:       "🧹 Limpieza Completa",
		Description: "Se han eliminado " + strconv.Itoa(num-1) + " mensajes.",
		Color:       0x2ecc71,
	}
	ctx.ReplyEmbed(embed)
}
