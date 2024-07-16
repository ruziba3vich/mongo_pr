package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/to-do-with-mongo/internal/models"
	"github.com/ruziba3vich/to-do-with-mongo/internal/service"
)

type Handler struct {
	service *service.Service
	context context.Context
}

func New(service *service.Service, context context.Context) *Handler {
	return &Handler{
		service: service,
		context: context,
	}
}

func (h *Handler) CreateTaskHandler(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response, err := h.service.CreateTask(h.context, &task)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}

func (h *Handler) GetTaskByIdHandler(c *gin.Context) {
	var taskId string
	if err := c.ShouldBindJSON(&taskId); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response, err := h.service.GetTaskById(h.context, taskId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}

func (h *Handler) UpdateTaskStatusHandler(c *gin.Context) {
	var req models.UpdateTaskStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response, err := h.service.UpdateTaskStatus(h.context, &req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}

func (h *Handler) GetIncompleteSubTasksHandler(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response, err := h.service.GetIncompleteSubTasks(h.context, &req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}

func (h *Handler) GetTasksUntilDateHandler(c *gin.Context) {
	var req models.GetTasksUntilDateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response, err := h.service.GetTasksUntilDate(h.context, &req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}

func (h *Handler) UpdateSubTaskStatusHandler(c *gin.Context) {
	var req models.UpdateSubTaskStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response, err := h.service.UpdateSubTaskStatus(h.context, &req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}

func (h *Handler) AddNewSubTaskIntoTaskHandler(c *gin.Context) {
	var req models.AddNewSubTaskIntoTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response, err := h.service.AddNewSubTaskIntoTask(h.context, &req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}

func (h *Handler) ChangeTaskUserHandler(c *gin.Context) {
	var req models.ChangeTaskUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response, err := h.service.ChangeTaskUser(h.context, &req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}
