FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o terminal ./cmd/terminal

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/terminal .
COPY migrations/init.sql ./migrations/init.sql
COPY .env.example .env
EXPOSE 8080
CMD ["./terminal"]
