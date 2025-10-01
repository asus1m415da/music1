package commands

import (
	"discord-go-music-bot/internal/state"
	"github.com/bwmarrin/discordgo"
)

func Help(ctx *state.Context) {
	prefix := "/"
	if ctx.GetSourceType() == int(state.SourceTypeMessage) {
		prefix = "!"
	}

	embed := &discordgo.MessageEmbed{
		Title:       "📋 Comandos del Bot",
		Description: "Lista de comandos disponibles",
		Color:       0x3498db,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "🎵 Música",
				Value:  prefix + "reproducir <url> - Reproduce una canción desde la URL\n" + prefix + "buscar <búsqueda> - Busca y reproduce una canción\n" + prefix + "saltar - Salta la canción actual\n" + prefix + "cola - Muestra la cola de reproducción\n" + prefix + "detener - Detiene la reproducción y limpia la cola\n" + prefix + "pausar - Pausa la reproducción\n" + prefix + "reanudar - Reanuda la reproducción",
				Inline: false,
			},
			{
				Name:   "🔊 Volumen",
				Value:  prefix + "volumen <valor> - Establece el volumen (0 a 200)\n" + prefix + "volumenactual - Muestra el volumen actual",
				Inline: false,
			},
			{
				Name:   "🛠️ Utilidades",
				Value:  prefix + "ping - Responde con Pong\n" + prefix + "pong - Responde con Ping\n" + prefix + "limpiar <número> - Elimina la cantidad especificada de mensajes\n" + prefix + "tiempoactivo - Muestra el tiempo activo del bot\n" + prefix + "version - Muestra la versión del bot\n" + prefix + "info - Información del bot\n" + prefix + "ayuda - Muestra este mensaje",
				Inline: false,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Bot de Música Discord",
		},
	}
	ctx.ReplyEmbed(embed)
}
