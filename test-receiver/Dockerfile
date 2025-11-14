# Этап сборки
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o test-receiver ./main.go   # бинарник создаётся в /app/test-receiver

# Финальный образ
FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/test-receiver ./  

EXPOSE 8080

CMD ["./test-receiver"]   # запускаем бинарник относительно WORKDIR
