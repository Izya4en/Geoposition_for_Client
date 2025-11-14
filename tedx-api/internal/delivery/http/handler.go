package http

import (
	"net/http"
	"tedx-api/internal/dto"
	"tedx-api/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService *service.AuthService
}

func NewHandler(auth *service.AuthService) *Handler {
	return &Handler{authService: auth}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.POST("/login", h.Login)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}

func (h *Handler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{Token: token})
}
