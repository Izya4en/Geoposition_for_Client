FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod tidy && go build -o reservation ./cmd/main.go

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/reservation .
CMD ["./reservation"]
