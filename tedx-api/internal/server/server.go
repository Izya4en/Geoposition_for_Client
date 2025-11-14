package server

import (
	"log"
	"tedx-api/internal/config"
	"tedx-api/internal/delivery/http"
	"tedx-api/internal/repository"
	"tedx-api/internal/service"

	"github.com/gin-gonic/gin"
)

func Run() error {
	cfg := config.Load()
	repo := repository.NewInMemoryUserRepo()
	auth := service.NewAuthService(repo, cfg)

	router := gin.Default()
	handler := http.NewHandler(auth)
	handler.RegisterRoutes(router)

	log.Printf("TEDx API running on :%s", cfg.Port)
	return router.Run(":" + cfg.Port)
}
