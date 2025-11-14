package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"user-simulator/internal/model"
	"user-simulator/internal/service"
	"user-simulator/internal/transport"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://localhost:8080"
	}

	user := model.NewUser(1, 10000.0, 43.238949, 76.889709)
	client := transport.NewHTTPClient(apiURL)
	sim := service.NewSimulator(user, client)

	log.Printf("🚀 Симулятор запущен, отправка данных в %s...", apiURL)
	sim.Start(5 * time.Second)
}
