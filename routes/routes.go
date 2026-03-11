// routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"todo_API/handlers"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/tasks", handlers.CreateTask)
	r.GET("/tasks", handlers.GetTasks)
	r.GET("/tasks/:id", handlers.GetTaskByID)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.DELETE("/tasks/:id", handlers.DeleteTask)

} 