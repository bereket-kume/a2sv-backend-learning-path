package data

import (
	"context"
	"errors"
	"task_manager/database"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAlltask() ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := database.TaskCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New("find failed")
	}
	defer cursor.Close(ctx)

	var tasks []models.Task

	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, errors.New("failed to decode task")
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.New("cursor iteration error")
	}

	return tasks, nil
}

func GetTaskById(id string) (models.Task, error) {
	objId, er := primitive.ObjectIDFromHex(id)
	if er != nil {
		return models.Task{}, errors.New("invalid objectid")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var task models.Task
	err := database.TaskCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&task)
	if err != nil {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}

func CreateTask(task models.Task) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := database.TaskCollection.InsertOne(ctx, task)
	if err != nil {
		return models.Task{}, errors.New("insert failed")
	}

	if objId, ok := res.InsertedID.(primitive.ObjectID); ok {
		task.ID = objId
	}

	return task, nil
}

func UpdateTask(id string, updated models.Task) (models.Task, error) {
	objId, er := primitive.ObjectIDFromHex(id)
	if er != nil {
		return models.Task{}, errors.New("invalid objectid")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{"$set": bson.M{
		"title":       updated.Title,
		"description": updated.Description,
		"status":      updated.Status,
	}}
	_, err := database.TaskCollection.UpdateByID(ctx, objId, update)
	if err != nil {
		return models.Task{}, errors.New("update failed")
	}
	return updated, nil
}

func DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ObjectID")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := database.TaskCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return errors.New("delete task failed")
	}
	if res.DeletedCount == 0 {
		return errors.New("no task found to delete")
	}
	return nil
}
