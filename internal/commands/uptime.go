package commands

import (
	"discord-go-music-bot/internal/state"
	"discord-go-music-bot/internal/validation"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Uptime(ctx *state.Context) {
	if !validation.HasPermission(ctx, discordgo.PermissionAdministrator) {
		embed := &discordgo.MessageEmbed{
			Title:       "❌ Sin Permiso",
			Description: "No tienes permiso para usar este comando.",
			Color:       0xe74c3c,
		}
		ctx.ReplyEmbed(embed)
		return
	}
	timeNow := time.Now()
	uptime := timeNow.Sub(state.StartTime)
	days := int(uptime.Hours() / 24)
	hours := int(uptime.Hours()) % 24
	minutes := int(uptime.Minutes()) % 60
	seconds := int(uptime.Seconds()) % 60

	var uptimeMessage strings.Builder
	if days > 0 {
		if days == 1 {
			uptimeMessage.WriteString("1 día, ")
		} else {
			uptimeMessage.WriteString(strconv.Itoa(days) + " días, ")
		}
	}
	if hours > 0 {
		if hours == 1 {
			uptimeMessage.WriteString("1 hora, ")
		} else {
			uptimeMessage.WriteString(strconv.Itoa(hours) + " horas, ")
		}
	}
	if minutes > 0 {
		if minutes == 1 {
			uptimeMessage.WriteString("1 minuto y ")
		} else {
			uptimeMessage.WriteString(strconv.Itoa(minutes) + " minutos y ")
		}
	}
	if seconds > 0 {
		if seconds == 1 {
			uptimeMessage.WriteString("1 segundo")
		} else {
			uptimeMessage.WriteString(strconv.Itoa(seconds) + " segundos")
		}
	}

	embed := &discordgo.MessageEmbed{
		Title:       "⏱️ Tiempo Activo",
		Description: "**Tiempo en línea:** " + uptimeMessage.String(),
		Color:       0xf39c12,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Bot de Música Discord",
		},
	}
	ctx.ReplyEmbed(embed)
}
