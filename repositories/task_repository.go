// task_repository.go
package repositories

import (
	"context"

	"todo_API/database"
	"todo_API/models"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateTask(task models.Task) error {

	collection := database.DB.Collection("tasks")

	_, err := collection.InsertOne(context.Background(), task)

	return err
}

func GetTasks(status, priority string) ([]models.Task, error) {

	collection := database.DB.Collection("tasks")

	filter := bson.M{}
	if status != "" {
		filter["status"] = status
	}
	if priority != "" {
		filter["priority"] = priority
	}

	cursor, err := collection.Find(context.Background(), filter)

	var tasks []models.Task

	if err != nil {
		return tasks, err
	}

	err = cursor.All(context.Background(), &tasks)
	if tasks == nil {
		tasks = []models.Task{} // Return empty array instead of null
	}

	return tasks, err
}

func GetTaskByID(id string) (models.Task, error) {

	var task models.Task

	err := database.DB.Collection("tasks").
		FindOne(context.TODO(), bson.M{"_id": id}).
		Decode(&task)

	return task, err
}

func UpdateTask(id string, task models.Task) error {
	collection := database.DB.Collection("tasks")
	
	update := bson.M{"$set": task}
	
	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	
	return err
}

func DeleteTask(id string) error {
	collection := database.DB.Collection("tasks")
	
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	
	return err
}