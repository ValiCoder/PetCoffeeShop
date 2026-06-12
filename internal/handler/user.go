package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler будет содержать зависимости, например, usecase или db
type UserHandler struct{}

// NewUserHandler — конструктор для инициализации в main.go
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// PingHandler проверяет работоспособность сервера
// @Summary      Ping-Pong check
// @Description  Simple endpoint to check if server is alive
// @Tags         test
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /ping [get]
func (h *UserHandler) PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
