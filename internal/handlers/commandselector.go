package handlers

import (
        "discord-go-music-bot/internal/commands"
        "discord-go-music-bot/internal/state"
)

// Both handlers can use this to forward to the correct command
func commandSelector(ctx *state.Context) {
        if state.DisabledCommands[ctx.CommandName] {
                ctx.Reply("This command has been disabled.")
                return
        }

        switch ctx.CommandName {
        case "ping":
                commands.Pong(ctx)
        case "pong":
                commands.Ping(ctx)
        case "play", "reproducir":
                commands.AddSong(ctx, false)
        case "search", "buscar":
                commands.AddSong(ctx, true)
        case "skip", "saltar":
                commands.SkipSong(ctx)
        case "queue", "cola":
                commands.ShowQueue(ctx)
        case "stop", "detener":
                commands.StopSong(ctx)
        case "pause", "pausar", "resume", "reanudar":
                commands.PauseSong(ctx)
        case "volume", "volumen":
                commands.SetVolume(ctx)
        case "currentvolume", "volumenactual":
                commands.CurrentVolume(ctx)
        case "nuke", "limpiar":
                commands.NukeMessages(ctx)
        case "uptime", "tiempoactivo":
                commands.Uptime(ctx)
        case "version":
                commands.Version(ctx)
        case "help", "ayuda":
                commands.Help(ctx)
        case "oss", "info":
                commands.Oss(ctx)
        default:
                commands.Unknown(ctx)
        }
}
