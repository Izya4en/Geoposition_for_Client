package http

import (
	"terminal/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Svc *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{Svc: svc}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
}
