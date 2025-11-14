package main

import (
	"log"
	"tedx-api/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
