package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/to-do-with-mongo/internal/models"
	"github.com/ruziba3vich/to-do-with-mongo/internal/service"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateTaskHandler(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response, err := h.service.CreateTask(&task)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}
