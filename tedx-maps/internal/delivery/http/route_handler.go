package http

import (
	"net/http"

	"tedx-maps/internal/entity"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateRoute(c *gin.Context) {
	var r entity.Route
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := h.routeService.CreateRoute(c, &r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (h *Handler) GetRoutes(c *gin.Context) {
	routes, err := h.routeService.GetRoutes(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, routes)
}
