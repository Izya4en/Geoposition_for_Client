FROM golang:1.24-alpine

WORKDIR /app
COPY . .
RUN go mod tidy && go build -o user-simulator ./cmd/main.go

CMD ["./user-simulator"]
