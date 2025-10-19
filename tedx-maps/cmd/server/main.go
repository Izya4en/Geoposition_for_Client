package main

import (
	"log"

	"tedx-maps/config"
	"tedx-maps/internal/server"
)

func main() {
	// Загружаем конфиг из .env
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	// Инициализируем сервер
	srv := server.NewServer(cfg)

	// Запускаем сервер
	if err := srv.Run(); err != nil {
		log.Fatalf("❌ Server stopped: %v", err)
	}
}
