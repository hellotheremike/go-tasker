package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hellotheremike/go-tasker/internal/middleware"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAll(c *gin.Context) {
	users, err := h.service.GetAll(c.Request.Context())
	userId, _ := c.Get(string(middleware.ContextUserKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return;
	}
	c.JSON(http.StatusOK, gin.H{"user_id": userId, "users": users})
}