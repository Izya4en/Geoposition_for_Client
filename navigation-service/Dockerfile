FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o navigation ./cmd

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/navigation .
EXPOSE 8083
CMD ["./navigation"]
