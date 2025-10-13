package http

import (
	"net/http"

	"tedx-maps/internal/entity"
	"tedx-maps/pkg/auth"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterUser(c *gin.Context) {
	var u entity.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.Register(c, &u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, _ := auth.GenerateJWT(u.Email)
	c.JSON(http.StatusCreated, gin.H{"token": token})
}
