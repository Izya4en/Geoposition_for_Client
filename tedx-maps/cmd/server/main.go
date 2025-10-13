package main

import (
	"log"
	"os"

	"tedx-maps/config"
	"tedx-maps/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	// 1️⃣ Загружаем переменные окружения из .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  Warning: .env file not found or couldn't be loaded — continuing with system env")
	}

	// 2️⃣ Загружаем YAML-конфиг
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	// 3️⃣ Опционально: переопределяем значения из .env (если есть)
	if port := os.Getenv("APP_PORT"); port != "" {
		cfg.Server.Port = port
	}
	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		cfg.Database.URL = dbURL
	}

	// 4️⃣ Создаём сервер
	srv := server.NewServer(cfg)

	// 5️⃣ Запускаем сервер
	if err := srv.Run(); err != nil {
		log.Fatalf("❌ Server stopped: %v", err)
	}

	log.Println("🚀 TEDx Maps server is running on port:", cfg.Server.Port)
}
