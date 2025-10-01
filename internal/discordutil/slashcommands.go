package discordutil

import (
        "discord-go-music-bot/internal/constants"
        "discord-go-music-bot/internal/logging"
        "discord-go-music-bot/internal/state"

        "github.com/bwmarrin/discordgo"
)

func SetupSlashCommands(s *discordgo.Session) {
        logging.InfoLog(constants.ANSIBlue + "Setting up slash commands")
        var theNumberOneAsFloat float64 = 1.0

        commands := []*discordgo.ApplicationCommand{
                {Name: "ping", Description: "Responde con Pong"},
                {Name: "pong", Description: "Responde con Ping"},
                {Name: "reproducir", Description: "Reproduce música (URL o nombre de canción)",
                        Options: []*discordgo.ApplicationCommandOption{
                                {
                                        Type:        discordgo.ApplicationCommandOptionString,
                                        Name:        "url",
                                        Description: "URL de YouTube o nombre de la canción",
                                        Required:    true,
                                },
                        },
                },
                {Name: "buscar", Description: "Busca una canción para reproducir",
                        Options: []*discordgo.ApplicationCommandOption{
                                {
                                        Type:        discordgo.ApplicationCommandOptionString,
                                        Name:        "query",
                                        Description: "La búsqueda a realizar",
                                        Required:    true,
                                },
                        },
                },
                {Name: "saltar", Description: "Salta la canción actual"},
                {Name: "cola", Description: "Muestra la cola actual"},
                {Name: "detener", Description: "Detiene la reproducción y limpia la cola"},
                {Name: "pausar", Description: "Pausa la canción actual"},
                {Name: "reanudar", Description: "Reanuda la canción actual"},
                {Name: "volumen", Description: "Establece el volumen (0-200)",
                        Options: []*discordgo.ApplicationCommandOption{
                                {
                                        Type:        discordgo.ApplicationCommandOptionInteger,
                                        Name:        "level",
                                        Description: "El nivel de volumen (0-200)",
                                        Required:    false,
                                },
                        },
                },
                {Name: "volumenactual", Description: "Muestra el volumen actual"},
                {Name: "limpiar", Description: "Elimina una cantidad de mensajes",
                        Options: []*discordgo.ApplicationCommandOption{
                                {
                                        Type:        discordgo.ApplicationCommandOptionInteger,
                                        Name:        "count",
                                        Description: "La cantidad de mensajes a eliminar (1-200)",
                                        Required:    true,
                                        MaxValue:    200.0,
                                        MinValue:    &theNumberOneAsFloat,
                                },
                        },
                },
                {Name: "tiempoactivo", Description: "Muestra el tiempo activo del bot"},
                {Name: "version", Description: "Muestra la versión del bot"},
                {Name: "ayuda", Description: "Muestra información de ayuda"},
                {Name: "info", Description: "Muestra información del bot"},
        }

        // Registrar comandos globalmente usando Application ID
        existingCommands, err := s.ApplicationCommands(state.ApplicationID, "")
        if err != nil {
                logging.FatalLog("No se pudieron obtener los comandos existentes: " + err.Error())
        }

        // Crear un mapa de comandos deseados
        desiredCommands := make(map[string]*discordgo.ApplicationCommand)
        for _, cmd := range commands {
                if !state.DisabledCommands[cmd.Name] {
                        desiredCommands[cmd.Name] = cmd
                }
        }

        // Eliminar comandos que ya no existen
        for _, existingCmd := range existingCommands {
                if _, exists := desiredCommands[existingCmd.Name]; !exists {
                        err := s.ApplicationCommandDelete(state.ApplicationID, "", existingCmd.ID)
                        if err != nil {
                                logging.WarningLog("No se pudo eliminar el comando: " + existingCmd.Name)
                        } else {
                                logging.InfoLog("Comando eliminado: " + existingCmd.Name)
                        }
                }
        }

        // Registrar o actualizar comandos
        for _, cmd := range commands {
                if state.DisabledCommands[cmd.Name] {
                        logging.WarningLog("Omitiendo comando deshabilitado: " + cmd.Name)
                        continue
                }
                
                found := false
                var existingCmdID string
                for _, existingCmd := range existingCommands {
                        if cmd.Name == existingCmd.Name {
                                found = true
                                existingCmdID = existingCmd.ID
                                break
                        }
                }
                
                if found {
                        // Actualizar el comando existente
                        _, err := s.ApplicationCommandEdit(state.ApplicationID, "", existingCmdID, cmd)
                        if err != nil {
                                logging.WarningLog("No se pudo actualizar el comando: " + cmd.Name + " - " + err.Error())
                        } else {
                                logging.InfoLog("Comando actualizado: " + cmd.Name)
                        }
                } else {
                        // Crear nuevo comando
                        _, err := s.ApplicationCommandCreate(state.ApplicationID, "", cmd)
                        if err != nil {
                                logging.FatalLog("No se pudo crear el comando: " + cmd.Name + " " + err.Error())
                        } else {
                                logging.InfoLog("Comando registrado: " + cmd.Name)
                        }
                }
        }
        logging.InfoLog("Configuración de comandos slash completa.")
}
