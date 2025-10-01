# Bot de Música para Discord

## Descripción
Bot de música para Discord desarrollado en Go con comandos slash en español. Reproduce música desde YouTube en canales de voz.

## Dueño
- **Galaxy A06** (<@1404572152014962708>)

## Tecnología
- Lenguaje: Go 1.23
- Framework: discordgo
- Audio: yt-dlp, ffmpeg, opus

## Comandos Disponibles

### 🎵 Música
- `/reproducir <url>` - Reproduce una canción desde una URL de YouTube
- `/buscar <consulta>` - Busca y reproduce una canción de YouTube
- `/saltar` - Salta la canción actual
- `/cola` - Muestra la cola de reproducción
- `/detener` - Detiene la reproducción y limpia la cola
- `/pausar` - Pausa la reproducción
- `/reanudar` - Reanuda la reproducción

### 🔊 Volumen
- `/volumen <nivel>` - Establece el volumen (0-200)
- `/volumenactual` - Muestra el volumen actual

### 🛠️ Utilidades
- `/ping` - Responde con Pong
- `/pong` - Responde con Ping
- `/limpiar <cantidad>` - Elimina mensajes (requiere permisos)
- `/tiempoactivo` - Muestra el tiempo activo del bot
- `/version` - Muestra la versión del bot
- `/info` - Información del bot y dueño
- `/ayuda` - Muestra todos los comandos

## Configuración

### Variables de Entorno Requeridas
- `DISCORD_BOT_TOKEN` - Token del bot de Discord
- `DISCORD_APPLICATION_ID` - Application ID del bot

### Variables Opcionales
- `DISABLED_COMMANDS` - Comandos deshabilitados (separados por comas)
- `UNKNOWN_COMMANDS` - Comportamiento para comandos desconocidos (ignore/help/error)

## Características
- Comandos slash registrados globalmente
- Respuestas con embeds bonitos y coloreados
- Soporte para búsqueda de YouTube
- Cola de reproducción
- Control de volumen
- Sistema de pausar/reanudar
- Mensajes en español

## Despliegue
- Compatible con Docker (Dockerfile multi-stage optimizado)
- Compatible con Railway, Render, Fly.io y otros servicios
- Compatible con Replit

## Cambios Recientes
- 2025: Traducción completa al español
- 2025: Implementación de embeds para todas las respuestas
- 2025: Actualización de información del dueño a Galaxy A06
- 2025: Registro de comandos globales con Application ID
- 2025: Dockerfile optimizado para multi-plataforma
