# 🏗️ Etapa de construcción en Go
FROM golang:1.23-alpine AS builder

# Instala herramientas necesarias para compilación con CGO
RUN apk add --no-cache build-base

WORKDIR /app

# Copia y descarga dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto del código
COPY . .

# Compila el bot con CGO habilitado
RUN CGO_ENABLED=1 GOOS=linux go build -o music-bot ./cmd/bot

# 🧱 Etapa final: entorno de ejecución
FROM alpine:latest

# Agrega repositorio community y actualiza
RUN echo "http://dl-cdn.alpinelinux.org/alpine/latest-stable/community" >> /etc/apk/repositories && \
    apk update && \
    apk add --no-cache \
    ffmpeg \
    opus \
    opus-dev \
    python3 \
    py3-pip \
    py3-setuptools \
    ca-certificates && \
    python3 -m ensurepip && \
    pip3 install --upgrade pip && \
    pip3 install --break-system-packages yt-dlp && \
    addgroup -S botuser && \
    adduser -S botuser -G botuser

WORKDIR /app

# Copia el binario compilado
COPY --from=builder /app/music-bot .

# Asigna permisos al usuario no root
RUN chown -R botuser:botuser /app

USER botuser

# Variables configurables en tiempo de ejecución
ENV UNKNOWN_COMMANDS="ignore"

CMD ["./music-bot"]
