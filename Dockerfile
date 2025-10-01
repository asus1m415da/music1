# 🏗️ Etapa de construcción en Go
FROM golang:1.23-alpine AS builder

RUN apk add --no-cache build-base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

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
    yt-dlp \
    python3 \
    py3-pip \
    py3-setuptools \
    ca-certificates \
    && addgroup -S botuser \
    && adduser -S botuser -G botuser

WORKDIR /app

COPY --from=builder /app/music-bot .

RUN chown -R botuser:botuser /app

USER botuser

ENV UNKNOWN_COMMANDS="ignore"

CMD ["./music-bot"]
