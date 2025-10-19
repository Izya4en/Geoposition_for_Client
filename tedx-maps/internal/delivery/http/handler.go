package http

import (
	"net/http"

	"tedx-maps/internal/service"


	"github.com/gin-gonic/gin"
)

type Handler struct {
	pointService *service.PointService
	routeService *service.RouteService
	userService  *service.UserService
	mapService   *service.MapService
}

func NewHandler(
	pointService *service.PointService,
	routeService *service.RouteService,
	userService *service.UserService,
	mapService *service.MapService,
) *Handler {
	return &Handler{
		pointService: pointService,
		routeService: routeService,
		userService:  userService,
		mapService:   mapService,
	}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")

	// публичные маршруты
	api.POST("/register", h.RegisterUser)
	api.GET("/map", h.GetMapData)

	// защищённые маршруты
	protected := api.Group("/")
	

	{
		protected.POST("/points", h.CreatePoint)
		protected.GET("/points", h.GetPoints)
		protected.GET("/points/nearest", h.GetNearestPoints)

		protected.POST("/routes", h.CreateRoute)
		protected.GET("/routes", h.GetRoutes)
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}
