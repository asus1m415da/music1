package bot

import (
        "discord-go-music-bot/internal/discordutil"
        "discord-go-music-bot/internal/handlers"
        "discord-go-music-bot/internal/logging"
        "discord-go-music-bot/internal/state"
        "os"
        "os/exec"
        "strings"

        "github.com/bwmarrin/discordgo"
        "github.com/joho/godotenv"
)

func setup() { // find env, get bot token, check dependencies

        // Intentar cargar .env si existe (para desarrollo local)
        // En producción (Replit, Railway, etc.) las variables ya están en el entorno
        _ = godotenv.Load()
        state.Token = os.Getenv("DISCORD_BOT_TOKEN")
        if state.Token == "" {
                logging.FatalLog("Token no encontrado - verifica el archivo .env")
        }

        state.ApplicationID = os.Getenv("DISCORD_APPLICATION_ID")
        if state.ApplicationID == "" {
                logging.FatalLog("Application ID no encontrado - verifica el archivo .env")
        }

        if _, err := exec.LookPath("yt-dlp"); err != nil {
                logging.FatalLog("yt-dlp no encontrado. Por favor instálalo con: pip install yt-dlp")
        }

        if _, err := exec.LookPath("ffmpeg"); err != nil {
                logging.FatalLog("ffmpeg no encontrado. Por favor instálalo con tu gestor de paquetes")
        }

        // Parse disabled commands from .env
        disabled := os.Getenv("DISABLED_COMMANDS")
        for _, cmd := range strings.Split(disabled, ",") {
                cmd = strings.TrimSpace(cmd)
                if cmd != "" {
                        state.DisabledCommands[cmd] = true
                }
        }
}

func Run() {
        setup()
        dg, err := discordgo.New("Bot " + state.Token)
        if err != nil {
                logging.FatalLog("Error creating Discord session: " + err.Error())
        }

        dg.AddHandler(handlers.HandleMessageCreate)
        dg.AddHandler(handlers.HandleInteractionCreate)

        err = dg.Open()

        discordutil.SetupSlashCommands(dg)

        if err != nil {
                logging.FatalLog("Error opening connection: " + err.Error())
        }
        defer dg.Close()
        logging.InfoLog("Versión: " + state.GoSourceHash)
        logging.InfoLog("Bot en ejecución. Presiona CTRL-C para salir.")
        select {} // block forever
}
