package http

import (
	"net/http"
	"reservation/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	srv *service.ReservationService
}

func NewHandler(srv *service.ReservationService) *Handler {
	return &Handler{srv: srv}
}

func (h *Handler) Run(addr string) {
	r := gin.Default()

	r.POST("/reserve", func(c *gin.Context) {
		userID := c.PostForm("user_id")
		amount, _ := strconv.ParseFloat(c.PostForm("amount"), 64)

		h.srv.ReserveMoney(userID, amount)
		c.JSON(http.StatusOK, gin.H{"status": "reserved"})
	})

	r.GET("/reservations", func(c *gin.Context) {
		c.JSON(http.StatusOK, h.srv.GetAll())
	})

	r.Run(addr)
}
