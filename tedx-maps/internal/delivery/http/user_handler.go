package http

import (
	"net/http"

	"tedx-maps/internal/entity"


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


	
}
