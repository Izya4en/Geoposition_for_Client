package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"navigation-service/internal/service"
)

type Handler struct {
	service *service.NavigationService
}

func NewHandler(s *service.NavigationService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "navigation service is running"})
	})

	r.POST("/route", h.findRoute)
}

type RouteRequest struct {
	UserLat float64 `json:"user_lat"`
	UserLon float64 `json:"user_lon"`
	Amount  float64 `json:"amount"`
}

func (h *Handler) findRoute(c *gin.Context) {
	var req RouteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.service.FindNearest(req.UserLat, req.UserLon, req.Amount)
	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "no terminal found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"nearest_terminal": result})
}
