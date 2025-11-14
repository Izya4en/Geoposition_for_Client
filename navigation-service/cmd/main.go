package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"navigation-service/internal/config"
	"navigation-service/internal/delivery/http"
	"navigation-service/internal/repository"
	"navigation-service/internal/service"
)

func main() {
	cfg := config.Load()

	repo := repository.NewTerminalRepository()
	svc := service.NewNavigationService(repo)
	handler := http.NewHandler(svc)

	r := gin.Default()
	handler.RegisterRoutes(r)

	addr := fmt.Sprintf(":%s", cfg.Port)
	r.Run(addr)
}
