package repositories

import (
	"context"
	domain "task_manager/Domain"
	repositories "task_manager/Repositories"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepositoryTestSuite struct {
	suite.Suite
	Client         *mongo.Client
	Database       *mongo.Database
	TaskCollection *mongo.Collection
	Repository     *repositories.TaskRepository
}

func (suite *TaskRepositoryTestSuite) SetupTest() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	suite.NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	suite.NoError(err)

	suite.Client = client
	suite.Database = client.Database("task_manager_test")
	suite.TaskCollection = suite.Database.Collection("tasks")
	suite.Repository = repositories.NewTaskRepository(suite.TaskCollection)

}

func (suite *TaskRepositoryTestSuite) TearDownTest() {
	err := suite.Database.Drop(context.Background())
	suite.NoError(err)
	err = suite.Client.Disconnect(context.Background())
	suite.NoError(err)
}

func (suite *TaskRepositoryTestSuite) TestGetAllTasks() {
	expectedTasks := []*domain.Task{
		{ID: "1", Title: "Task 1", Description: "Description 1", DueDate: time.Now(), Status: "Pending"},
		{ID: "2", Title: "Task 2", Description: "Description 2", DueDate: time.Now(), Status: "Completed"},
	}
	tasksInterface := make([]interface{}, len(expectedTasks))
	for i, task := range expectedTasks {
		tasksInterface[i] = task
	}
	suite.TaskCollection.InsertMany(context.Background(), tasksInterface)

	tasks, err := suite.Repository.GetAlltask()
	suite.NoError(err)
	suite.Equal(len(expectedTasks), len(tasks))
	suite.Equal(expectedTasks[0].Title, tasks[0].Title)
}

func (suite *TaskRepositoryTestSuite) TestGetTaskById() {
	task := bson.M{
		"Title":       "Task 1",
		"Description": "Description 1",
		"DueDate":     time.Now(),
		"Status":      "Pending",
	}
	res, err := suite.TaskCollection.InsertOne(context.Background(), task)
	suite.NoError(err)
	objId := res.InsertedID.(primitive.ObjectID)
	taskId := objId.Hex()
	retrievedTask, err := suite.Repository.GetTaskById(taskId)
	suite.NoError(err)
	suite.Equal(task["Title"], retrievedTask.Title)
	suite.Equal(task["Description"], retrievedTask.Description)
	suite.Equal(task["Status"], retrievedTask.Status)
}

func (suite *TaskRepositoryTestSuite) TestCreateTask() {
	task := &domain.Task{
		Title:       "New Task",
		Description: "New Task Description",
		DueDate:     time.Now(),
		Status:      "Pending",
	}
	createdTask, err := suite.Repository.CreateTask(task)
	suite.NoError(err)
	suite.NotEmpty(createdTask.ID)
	suite.Equal(task.Title, createdTask.Title)
	suite.Equal(task.Description, createdTask.Description)
}

func (suite *TaskRepositoryTestSuite) TestUpdateTask() {
	task := bson.M{
		"Title":       "Task to Update",
		"Description": "Description to Update",
		"DueDate":     time.Now(),
		"Status":      "Pending",
	}
	res, err := suite.TaskCollection.InsertOne(context.TODO(), task)
	suite.NoError(err)

	id := res.InsertedID.(primitive.ObjectID).Hex()
	updateTask := &domain.Task{
		Title:       "Updated Task",
		Description: "Updated Description",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "Completed",
	}
	updated, err := suite.Repository.UpdateTask(id, updateTask)
	suite.NoError(err)
	suite.Equal(updateTask.Title, updated.Title)
	suite.Equal(updateTask.Description, updated.Description)
	suite.Equal(updateTask.Status, updated.Status)
	suite.NotEqual(task["DueDate"], updated.DueDate)
}

func (suite *TaskRepositoryTestSuite) TestDeleteTask() {
	task := bson.M{
		"Title":       "Task to Delete",
		"Description": "Description to Delete",
		"DueDate":     time.Now(),
		"Status":      "Pending",
	}
	res, err := suite.TaskCollection.InsertOne(context.TODO(), task)
	suite.NoError(err)

	id := res.InsertedID.(primitive.ObjectID).Hex()
	err = suite.Repository.DeleteTask(id)
	suite.NoError(err)

	_, err = suite.Repository.GetTaskById(id)
	suite.Error(err)
	count, err := suite.TaskCollection.CountDocuments(context.Background(), bson.M{"_id": res.InsertedID})
	suite.NoError(err)
	suite.Equal(int64(0), count)
}

func TestTaskRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryTestSuite))
}
