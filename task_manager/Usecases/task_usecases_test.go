package usecases

import (
	domain "task_manager/Domain"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TaskRepositoryMock struct {
	mock.Mock
}

func (m *TaskRepositoryMock) GetAlltask() ([]*domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Task), args.Error(1)
}

func (m *TaskRepositoryMock) GetTaskById(id string) (*domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *TaskRepositoryMock) CreateTask(task *domain.Task) (*domain.Task, error) {
	args := m.Called(task)
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *TaskRepositoryMock) UpdateTask(id string, update *domain.Task) (*domain.Task, error) {
	args := m.Called(id, update)
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *TaskRepositoryMock) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

type TaskUseCaseTestSuite struct {
	suite.Suite
	TaskUseCase  *TaskUseCase
	TaskRepoMock *TaskRepositoryMock
}

func (suite *TaskUseCaseTestSuite) SetupTest() {
	suite.TaskRepoMock = new(TaskRepositoryMock)
	suite.TaskUseCase = NewTaskUseCase(suite.TaskRepoMock)
}

func (suite *TaskUseCaseTestSuite) TestGetAllTasks() {
	expectedTasks := []*domain.Task{
		{ID: "1", Title: "Task 1", Description: "Description 1", DueDate: time.Now(), Status: "Pending"},
		{ID: "2", Title: "Task 2", Description: "Description 2", DueDate: time.Now(), Status: "Completed"},
	}
	suite.TaskRepoMock.On("GetAlltask").Return(expectedTasks, nil)
	tasks, err := suite.TaskUseCase.GetAllTasks()
	suite.NoError(err)
	suite.Equal(expectedTasks, tasks)
	suite.TaskRepoMock.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestGetTaskById() {
	expectedTask := &domain.Task{ID: "1", Title: "Task 1", Description: "Description 1", DueDate: time.Now(), Status: "Pending"}
	suite.TaskRepoMock.On("GetTaskById", "1").Return(expectedTask, nil)
	task, err := suite.TaskUseCase.GetTaskById("1")

	suite.NoError(err)
	suite.Equal(expectedTask, task)
	suite.TaskRepoMock.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestCreateTask() {
	newTask := &domain.Task{Title: "New Task", Description: "New Description", DueDate: time.Now(), Status: "Pending"}
	suite.TaskRepoMock.On("CreateTask", newTask).Return(newTask, nil)
	createdTask, err := suite.TaskUseCase.CreateTask(newTask)

	suite.NoError(err)
	suite.Equal(newTask, createdTask)
	suite.TaskRepoMock.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestUpdateTask() {
	updateTask := &domain.Task{Title: "Updated Task", Description: "Updated Description", DueDate: time.Now(), Status: "Completed"}
	suite.TaskRepoMock.On("UpdateTask", "1", updateTask).Return(updateTask, nil)
	updatedTask, err := suite.TaskUseCase.UpdateTask("1", updateTask)
	suite.NoError(err)
	suite.Equal(updatedTask, updateTask)
	suite.TaskRepoMock.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestDeleteTask() {
	suite.TaskRepoMock.On("DeleteTask", "1").Return(nil)
	err := suite.TaskUseCase.DeleteTask("1")
	suite.NoError(err)
	suite.TaskRepoMock.AssertExpectations(suite.T())
}

func TestTaskUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(TaskUseCaseTestSuite))
}
