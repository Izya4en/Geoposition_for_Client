package server

import (
	"fmt"

	"tedx-maps/config"
	"tedx-maps/internal/delivery/http"
	"tedx-maps/internal/repository"
	"tedx-maps/internal/service"
	"tedx-maps/pkg/mapsapi"
	"tedx-maps/utils"

	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) Run() error {
	// 1. Инициализация логгера
	logger := utils.NewLogger()

	// 2. Подключение к БД
	db, err := repository.NewPostgresDB(s.cfg)
	if err != nil {
		return fmt.Errorf("db init error: %v", err)
	}

	// 3. Репозитории
	pointRepo := repository.NewPointRepository(db)
	routeRepo := repository.NewRouteRepository(db)
	userRepo := repository.NewUserRepository(db)

	// 4. Внешний клиент (API карт)
	mapClient := mapsapi.NewClient(s.cfg.MapsAPIKey)

	// 5. Сервисы
	pointService := service.NewPointService(pointRepo)
	routeService := service.NewRouteService(routeRepo)
	userService := service.NewUserService(userRepo)
	mapService := service.NewMapService(mapClient)

	// 6. Инициализация HTTP-хендлеров
	handler := http.NewHandler(pointService, routeService, userService, mapService)

	router := gin.Default()
	handler.RegisterRoutes(router)

	logger.Info("Starting server on port ", s.cfg.ServerPort)
	return router.Run(":" + s.cfg.ServerPort)
}
