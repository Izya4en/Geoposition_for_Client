# Stage 1: Builder
FROM golang:1.25-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы зависимостей и скачиваем их
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарник
# CGO_ENABLED=0 отключает привязку к си-библиотекам (важно для alpine)
RUN CGO_ENABLED=0 GOOS=linux go build -o atm-service ./cmd/api/main.go

# Stage 2: Runner (Финальный образ)
FROM alpine:latest

WORKDIR /root/

# Устанавливаем сертификаты (если будут внешние HTTPS запросы)
RUN apk --no-cache add ca-certificates

# Копируем бинарник из первого этапа
COPY --from=builder /app/atm-service .

# Копируем конфигурацию и миграции (если они нужны при старте)
COPY --from=builder /app/config ./config
COPY --from=builder /app/migrations ./migrations

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./atm-service"]