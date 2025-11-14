package main

import (
	"log"
	"terminal/internal/config"
	"terminal/internal/delivery/http"
	"terminal/internal/repository"
	"terminal/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect db:", err)
	}

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	h := http.NewHandler(svc)

	r := gin.Default()
	h.RegisterRoutes(r)

	log.Printf("Terminal service running on :%s", cfg.Port)
	r.Run(":" + cfg.Port)
}
