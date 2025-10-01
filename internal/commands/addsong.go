package commands

import (
	"discord-go-music-bot/internal/audio"
	"discord-go-music-bot/internal/discordutil"
	"discord-go-music-bot/internal/logging"
	"discord-go-music-bot/internal/state"
	"discord-go-music-bot/internal/validation"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func AddSong(ctx *state.Context, search_mode bool) {
	var url string

	if !discordutil.IsUserInVoiceChannel(ctx) {
		embed := &discordgo.MessageEmbed{
			Title:       "❌ Error",
			Description: "Debes estar en un canal de voz para usar este comando.",
			Color:       0xe74c3c,
		}
		ctx.ReplyEmbed(embed)
		return
	}

	if search_mode {
		if ctx.SourceType == state.SourceTypeInteraction {
			embed := &discordgo.MessageEmbed{
				Title:       "🔍 Buscando...",
				Description: "Buscando la canción en YouTube...",
				Color:       0x3498db,
			}
			ctx.ReplyEmbed(embed)
		}

		var hadToSanitise bool

		searchQuery := strings.TrimSpace(ctx.Arguments["query"])

		if !validation.IsValidSearchQuery(searchQuery) {
			var searchQuerySafeToUse bool
			searchQuery, searchQuerySafeToUse = validation.SanitiseSearchQuery(searchQuery)
			hadToSanitise = true
			if !searchQuerySafeToUse {
				embed := &discordgo.MessageEmbed{
					Title:       "❌ Error",
					Description: "Búsqueda inválida",
					Color:       0xe74c3c,
				}
				ctx.ReplyEmbed(embed)
				return
			}
		}

		var found_result bool
		url, found_result = audio.SearchYoutube(searchQuery)

		if !found_result {
			logging.ErrorLog("No se encontraron resultados para: " + searchQuery)
			embed := &discordgo.MessageEmbed{
				Title:       "❌ Sin Resultados",
				Description: "No se encontraron resultados para: " + searchQuery,
				Color:       0xe74c3c,
			}
			ctx.ReplyEmbed(embed)
			return
		}

		if hadToSanitise {
			embed := &discordgo.MessageEmbed{
				Title:       "✅ Encontrado",
				Description: "**URL:** " + url + "\n**Búsqueda:** " + searchQuery,
				Color:       0x2ecc71,
			}
			ctx.ReplyEmbed(embed)
		} else {
			embed := &discordgo.MessageEmbed{
				Title:       "✅ Encontrado",
				Description: "**URL:** " + url,
				Color:       0x2ecc71,
			}
			ctx.ReplyEmbed(embed)
		}
	} else {
		// Modo /reproducir - puede ser URL o búsqueda
		if len(ctx.Arguments["url"]) < 3 {
			embed := &discordgo.MessageEmbed{
				Title:       "❌ Error",
				Description: "Proporciona una URL o el nombre de una canción",
				Color:       0xe74c3c,
			}
			ctx.ReplyEmbed(embed)
			return
		}

		input := strings.TrimSpace(ctx.Arguments["url"])

		// Detectar si es una URL o un término de búsqueda
		isURL := validation.IsValidURL(input)

		if isURL {
			// Es una URL válida
			url = input
		} else {
			// No es una URL, buscar en YouTube automáticamente
			if ctx.SourceType == state.SourceTypeInteraction {
				embed := &discordgo.MessageEmbed{
					Title:       "🔍 Buscando...",
					Description: "Buscando \"" + input + "\" en YouTube...",
					Color:       0x3498db,
				}
				ctx.ReplyEmbed(embed)
			}

			// Validar el término de búsqueda
			searchQuery := input
			if !validation.IsValidSearchQuery(searchQuery) {
				var searchQuerySafeToUse bool
				searchQuery, searchQuerySafeToUse = validation.SanitiseSearchQuery(searchQuery)
				if !searchQuerySafeToUse {
					embed := &discordgo.MessageEmbed{
						Title:       "❌ Error",
						Description: "Búsqueda inválida",
						Color:       0xe74c3c,
					}
					ctx.ReplyEmbed(embed)
					return
				}
			}

			var found_result bool
			url, found_result = audio.SearchYoutube(searchQuery)

			if !found_result {
				logging.ErrorLog("No se encontraron resultados para: " + searchQuery)
				embed := &discordgo.MessageEmbed{
					Title:       "❌ Sin Resultados",
					Description: "No se encontraron resultados para: " + searchQuery,
					Color:       0xe74c3c,
				}
				ctx.ReplyEmbed(embed)
				return
			}

			// Mostrar qué se encontró
			embed := &discordgo.MessageEmbed{
				Title:       "✅ Encontrado",
				Description: "**Búsqueda:** " + searchQuery + "\n**URL:** " + url,
				Color:       0x2ecc71,
			}
			ctx.ReplyEmbed(embed)
		}
	}

	state.QueueMutex.Lock()
	state.Queue[ctx.GetGuildID()] = append(state.Queue[ctx.GetGuildID()], url)
	state.QueueMutex.Unlock()

	state.PlayingMutex.Lock()
	isAlreadyPlaying := state.Playing[ctx.GetGuildID()]
	state.PlayingMutex.Unlock()

	embed := &discordgo.MessageEmbed{
		Title:       "🎵 Añadido a la Cola",
		Description: "La canción se ha añadido a la cola de reproducción.",
		Color:       0x9b59b6,
	}
	ctx.ReplyEmbed(embed)

	if !isAlreadyPlaying {
		state.PlayingMutex.Lock()
		state.Playing[ctx.GetGuildID()] = true
		state.PlayingMutex.Unlock()
		audio.ProcessQueue(ctx)
	}
}
