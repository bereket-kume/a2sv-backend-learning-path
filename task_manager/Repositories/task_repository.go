package repositories

import (
	"context"
	"errors"
	"time"
	domain "task_manager/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepositoryImpl struct {
	TaskCollection *mongo.Collection
}

func NewTaskRepository(db *mongo.Collection) *TaskRepositoryImpl {
	return &TaskRepositoryImpl{TaskCollection: db}
}

func (r *TaskRepositoryImpl) GetAlltask() ([]*domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := r.TaskCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var tasks []*domain.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepositoryImpl) GetTaskById(id string) (*domain.Task, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ObjectId format")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var task domain.Task
	err = r.TaskCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepositoryImpl) CreateTask(task *domain.Task) (*domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := r.TaskCollection.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}
	if objId, ok := res.InsertedID.(primitive.ObjectID); ok {
		task.ID = objId
	} else {
		return nil, errors.New("failed to retrieve inserted ID")
	}
	return task, nil
}

func (r *TaskRepositoryImpl) UpdateTask(id string, updated *domain.Task) (*domain.Task, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ObjectId format")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	update := bson.M{"$set": bson.M{
		"title":       updated.Title,
		"description": updated.Description,
		"duedate":     updated.DueDate,
		"status":      updated.Status,
	}}
	_, err = r.TaskCollection.UpdateByID(ctx, objId, update)
	if err != nil {
		return nil, err
	}
	var task domain.Task
	err = r.TaskCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepositoryImpl) DeleteTask(id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ObjectId format")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = r.TaskCollection.DeleteOne(ctx, bson.M{"_id": objId})
	return err
}
