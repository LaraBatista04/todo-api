// task_service.go
package services

import (
	"errors"
	"strings"
	"time"
	"todo_API/models"
	"todo_API/repositories"

	"github.com/google/uuid"
)

func ValidateTask(task *models.Task) error {
	task.Title = strings.TrimSpace(task.Title)

	if len(task.Title) < 3 || len(task.Title) > 100 {
		return errors.New("title must be between 3 and 100 characters")
	}

	if task.Status == "" {
		task.Status = "pending"
	} else if !isValidStatus(task.Status) {
		return errors.New("invalid status")
	}

	if !isValidPriority(task.Priority) {
		if task.Priority == "" {
			task.Priority = "low"
		} else {
			return errors.New("invalid priority, must be low, medium, or high")
		}
	}

	if task.DueDate != "" {
		dueDate, err := time.Parse("2006-01-02", task.DueDate)
		if err != nil {
			return errors.New("due_date must be in YYYY-MM-DD format")
		}
        
        // compare to today in UTC at 00:00:00
        today := time.Now().UTC().Truncate(24 * time.Hour)
		if dueDate.UTC().Before(today) {
			return errors.New("due_date cannot be in the past")
		}
	}

	return nil
}

func CreateTask(task *models.Task) error {
	if err := ValidateTask(task); err != nil {
		return err
	}

	task.ID = uuid.New().String()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	return repositories.CreateTask(*task)
}

func UpdateTask(id string, input models.Task) (*models.Task, error) {
	existingTask, err := repositories.GetTaskByID(id)
	if err != nil {
		return nil, errors.New("task not found")
	}

	if existingTask.Status == "completed" {
		return nil, errors.New("cannot edit a completed task")
	}

	input.Title = strings.TrimSpace(input.Title)
	if input.Title != "" {
		if len(input.Title) < 3 || len(input.Title) > 100 {
			return nil, errors.New("title must be between 3 and 100 characters")
		}
		existingTask.Title = input.Title
	}

	if input.Status != "" {
		if !isValidStatus(input.Status) {
			return nil, errors.New("invalid status")
		}
		existingTask.Status = input.Status
	}

	if input.Priority != "" {
		if !isValidPriority(input.Priority) {
			return nil, errors.New("invalid priority, must be low, medium, or high")
		}
		existingTask.Priority = input.Priority
	}

	if input.DueDate != "" {
		dueDate, err := time.Parse("2006-01-02", input.DueDate)
		if err != nil {
			return nil, errors.New("due_date must be in YYYY-MM-DD format")
		}
        today := time.Now().UTC().Truncate(24 * time.Hour)
		if dueDate.UTC().Before(today) {
			return nil, errors.New("due_date cannot be in the past")
		}
		existingTask.DueDate = input.DueDate
	}

	if input.Description != "" {
		existingTask.Description = input.Description
	}

	existingTask.UpdatedAt = time.Now()

	err = repositories.UpdateTask(id, existingTask)
	if err != nil {
		return nil, err
	}

	return &existingTask, nil
}

func GetTasks(status, priority string) ([]models.Task, error) {
	return repositories.GetTasks(status, priority)
}

func GetTaskByID(id string) (models.Task, error) {
	return repositories.GetTaskByID(id)
}

func DeleteTask(id string) error {
	_, err := repositories.GetTaskByID(id)
	if err != nil {
		return errors.New("task not found")
	}
	return repositories.DeleteTask(id)
}

func isValidStatus(status string) bool {
	switch status {
	case "pending", "in_progress", "completed", "cancelled":
		return true
	}
	return false
}

func isValidPriority(priority string) bool {
	switch priority {
	case "low", "medium", "high":
		return true
	}
	return false
}  