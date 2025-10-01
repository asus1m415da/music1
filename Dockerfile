FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o music-bot ./cmd/bot

FROM alpine:latest

RUN apk add --no-cache ffmpeg opus opus-dev python3 py3-pip ca-certificates && \
    pip3 install --no-cache-dir yt-dlp && \
    addgroup -S botuser && adduser -S botuser -G botuser

WORKDIR /app

COPY --from=builder /app/music-bot .

RUN chown -R botuser:botuser /app

USER botuser

ENV DISCORD_BOT_TOKEN="" \
    DISCORD_APPLICATION_ID="" \
    DISABLED_COMMANDS="" \
    UNKNOWN_COMMANDS="ignore"

CMD ["./music-bot"]