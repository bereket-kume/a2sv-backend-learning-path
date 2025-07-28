package controller

import (
	domain "task_manager/Domain"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) Register(user *domain.User) (*domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserUsecase) Login(user *domain.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func (m *MockUserUsecase) GetUserByUsername(username string) (*domain.User, error) {
	args := m.Called(username)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserUsecase) PromoteUser(id string) (*domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.User), args.Error(1)
}

type MockTaskUseCase struct {
	mock.Mock
}

func (m *MockTaskUseCase) GetAllTasks() ([]*domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Task), args.Error(1)
}

func (m *MockTaskUseCase) GetTaskById(id string) (*domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *MockTaskUseCase) CreateTask(task *domain.Task) (*domain.Task, error) {
	args := m.Called(task)
	return args.Get(0).(*domain.Task), args.Error(1)
}
func (m *MockTaskUseCase) UpdateTask(id string, updated *domain.Task) (*domain.Task, error) {
	args := m.Called(id, updated)
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *MockTaskUseCase) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

type ControllerTestSuite struct {
	suite.Suite
	UserUseCaseMock *MockUserUsecase
	TaskUseCaseMock *MockTaskUseCase
}

func (suite *ControllerTestSuite) SetupTest() {
	suite.UserUseCaseMock = new(MockUserUsecase)
	suite.TaskUseCaseMock = new(MockTaskUseCase)
}

func (suite *ControllerTestSuite) TestRegisterUser() {
	user := &domain.User{
		Name:     "testuser",
		Password: "password123",
		Role:     domain.RoleAdmin,
	}
	expectedUser := &domain.User{
		ID:       "1",
		Name:     "testuser",
		Password: "password123",
		Role:     domain.RoleAdmin,
	}
	suite.UserUseCaseMock.On("Register", user).Return(expectedUser, nil)
	registeredUser, err := suite.UserUseCaseMock.Register(user)
	suite.NoError(err)
	suite.NotNil(registeredUser)
	suite.Equal(expectedUser, registeredUser)
	suite.UserUseCaseMock.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestLoginUser() {
	user := &domain.User{
		Name:     "testuser",
		Password: "password123",
	}
	expected_token := "mocked_token"
	suite.UserUseCaseMock.On("Login", user).Return(expected_token, nil)

	token, err := suite.UserUseCaseMock.Login(user)

	suite.NoError(err)
	suite.Equal(expected_token, token)
	suite.UserUseCaseMock.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestGetAlltask() {
	expectedTasks := []*domain.Task{
		{ID: "1", Title: "Task 1", Description: "Description 1", DueDate: time.Now(), Status: "Pending"},
		{ID: "2", Title: "Task 2", Description: "Description 2", DueDate: time.Now(), Status: "Completed"},
	}
	suite.TaskUseCaseMock.On("GetAllTasks").Return(expectedTasks, nil)

	tasks, err := suite.TaskUseCaseMock.GetAllTasks()
	suite.NoError(err)
	suite.Equal(len(expectedTasks), len(tasks))
	suite.Equal(expectedTasks[0].Title, tasks[0].Title)
	suite.TaskUseCaseMock.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestGetTaskById() {
	taskID := "1"
	expectedTask := &domain.Task{ID: taskID, Title: "Task 1", Description: "Description 1", DueDate: time.Now(), Status: "Pending"}
	suite.TaskUseCaseMock.On("GetTaskById", taskID).Return(expectedTask, nil)
	task, err := suite.TaskUseCaseMock.GetTaskById(taskID)
	suite.NoError(err)
	suite.Equal(expectedTask, task)
	suite.TaskUseCaseMock.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestCreateTask() {
	task := &domain.Task{
		Title:       "New Task",
		Description: "New Task Description",
		DueDate:     time.Now(),
		Status:      "Pending",
	}
	expectedTask := &domain.Task{
		ID:          "1",
		Title:       "New Task",
		Description: "New Task Description",
		DueDate:     time.Now(),
		Status:      "Pending",
	}
	suite.TaskUseCaseMock.On("CreateTask", task).Return(expectedTask, nil)

	createdTask, err := suite.TaskUseCaseMock.CreateTask(task)
	suite.NoError(err)
	suite.NotEmpty(createdTask.ID)
	suite.Equal(expectedTask.Title, createdTask.Title)
	suite.Equal(expectedTask.Description, createdTask.Description)
	suite.TaskUseCaseMock.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestUpdateTask() {
	taskID := "1"
	updatedTask := &domain.Task{
		Title:       "Updated Task",
		Description: "Updated Description",
		DueDate:     time.Now(),
		Status:      "Completed",
	}
	suite.TaskUseCaseMock.On("UpdateTask", taskID, updatedTask).Return(updatedTask, nil)

	result, err := suite.TaskUseCaseMock.UpdateTask(taskID, updatedTask)
	suite.NoError(err)
	suite.Equal(updatedTask.Title, result.Title)
	suite.Equal(updatedTask.Description, result.Description)
	suite.TaskUseCaseMock.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestDeleteTask() {
	taskID := "1"
	suite.TaskUseCaseMock.On("DeleteTask", taskID).Return(nil)

	err := suite.TaskUseCaseMock.DeleteTask(taskID)
	suite.NoError(err)
	suite.TaskUseCaseMock.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestGetUserByUsername() {
	username := "testuser"
	expectedUser := &domain.User{
		ID:       "1",
		Name:     username,
		Password: "password123",
		Role:     domain.RoleAdmin,
	}
	suite.UserUseCaseMock.On("GetUserByUsername", username).Return(expectedUser, nil)

	user, err := suite.UserUseCaseMock.GetUserByUsername(username)
	suite.NoError(err)
	suite.NotNil(user)
	suite.Equal(expectedUser.Name, user.Name)
	suite.UserUseCaseMock.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestPromoteUser() {
	userID := "1"
	expectedUser := &domain.User{
		ID:       userID,
		Name:     "testuser",
		Password: "password123",
		Role:     domain.RoleAdmin,
	}
	suite.UserUseCaseMock.On("PromoteUser", userID).Return(expectedUser, nil)

	user, err := suite.UserUseCaseMock.PromoteUser(userID)
	suite.NoError(err)
	suite.NotNil(user)
	suite.Equal(expectedUser.Role, user.Role)
	suite.UserUseCaseMock.AssertExpectations(suite.T())
}

func TestControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ControllerTestSuite))
}
