package http

import (
	"net/http"
	"strconv"

	"tedx-maps/internal/entity"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePoint(c *gin.Context) {
	var p entity.Point
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := h.pointService.CreatePoint(c, &p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (h *Handler) GetPoints(c *gin.Context) {
	points, err := h.pointService.GetAllPoints(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, points)
}

func (h *Handler) GetNearestPoints(c *gin.Context) {
	lat, _ := strconv.ParseFloat(c.Query("lat"), 64)
	lon, _ := strconv.ParseFloat(c.Query("lon"), 64)
	radius, _ := strconv.ParseFloat(c.Query("radius"), 64)

	points, err := h.pointService.GetNearestPoints(c, lat, lon, radius)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, points)
}
