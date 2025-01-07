package controllers

import (
	"go-services/internal/dto"
	"go-services/internal/services"
	"go-services/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService *services.TaskService
}

func NewTaskController() *TaskController {
	return &TaskController{
		taskService: services.NewTaskService(),
	}
}

func (tc *TaskController) Create(c *gin.Context) {
	var req dto.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (tc *TaskController) GetOne(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	task, err := tc.taskService.GetByID(uint(id))
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "Task not found")
		return
	}
	response.SuccessResponse(c, http.StatusOK, task)
}

func (tc *TaskController) GetAll(c *gin.Context) {
	tasks, err := tc.taskService.GetAll()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, tasks)
}

func (tc *TaskController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var req dto.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	task, err := tc.taskService.GetByID(uint(id))
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "Task not found")
		return
	}

	if req.Title != "" {
		task.Title = req.Title
	}
	if req.Description != "" {
		task.Description = req.Description
	}

	if req.Completed != task.Completed {
		task.Completed = req.Completed
	}

	err = tc.taskService.Update(task)

	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, task)
}

func (tc *TaskController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	err = tc.taskService.Delete(uint(id))
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Task deleted successfully")
}
