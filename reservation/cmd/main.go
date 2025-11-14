package main

import (
	"log"
	"reservation/internal/config"
	"reservation/internal/delivery/http"
	"reservation/internal/kafka"
	"reservation/internal/repository"
	"reservation/internal/service"
)

func main() {
	cfg := config.Load()

	producer := kafka.NewProducer(cfg.KafkaBroker)
	repo := repository.NewReservationRepo()
	srv := service.NewReservationService(repo, producer)
	handler := http.NewHandler(srv)

	log.Println("Reservation service running on :8081")
	handler.Run(":8081")
}
