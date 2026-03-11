// task_handler.go
package handlers

import (
	"net/http"
	"log/slog"
	"github.com/gin-gonic/gin"

	"todo_API/models"
	"todo_API/services"
)

func CreateTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		slog.Warn("Tentativa de criação com JSON inválido", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateTask(&task)
	if err != nil {

		slog.Error("Erro ao criar tarefa",
			"title", task.Title,
			"error", err.Error(),
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	slog.Info("Tarefa criada com sucesso", "id", task.ID)
	c.JSON(http.StatusCreated, task)
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	task, err := services.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func GetTasks(c *gin.Context) {
	status := c.Query("status")
	priority := c.Query("priority")

	tasks, err := services.GetTasks(status, priority)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var taskInput models.Task

	if err := c.ShouldBindJSON(&taskInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask, err := services.UpdateTask(id, taskInput)
	if err != nil {
		if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	err := services.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}
