package routers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserController struct {
	mock.Mock
}

func (m *MockUserController) Register(c *gin.Context) {
	m.Called(c)
	c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

func (m *MockUserController) Login(c *gin.Context) {
	m.Called(c)
	c.JSON(http.StatusOK, gin.H{"message": "User logged in"})
}

func (m *MockUserController) GetUserByUsername(c *gin.Context) {
	m.Called(c)
	c.JSON(http.StatusOK, gin.H{"message": "User found"})
}

func (m *MockUserController) Promote(c *gin.Context) {
	m.Called(c)
	c.JSON(http.StatusOK, gin.H{"message": "User promoted"})
}

type MockTaskController struct {
	mock.Mock
}

func (m *MockTaskController) GetAllTasks(c *gin.Context) {
	m.Called(c)
	c.JSON(http.StatusOK, gin.H{"message": "All tasks retrieved"})
}

func (m *MockTaskController) GetTaskById(c *gin.Context) {
	m.Called(c)
	c.JSON(http.StatusOK, gin.H{"message": "Task retrieved"})
}

func (m *MockTaskController) CreateTask(c *gin.Context) {
	m.Called(c)
	c.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

func (m *MockTaskController) UpdateTask(c *gin.Context) {
	m.Called(c)
	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func (m *MockTaskController) DeleteTask(c *gin.Context) {
	m.Called(c)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

// Mock Middleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Bypass authentication for testing
	}
}

func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Bypass admin check for testing
	}
}

// SetUpRoutes for testing
func SetUpTestRoutes(userController *MockUserController, taskController *MockTaskController) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/register", userController.Register)
			users.POST("/login", userController.Login)
			users.POST("/promote/:id", AuthMiddleware(), Admin(), userController.Promote)
			users.GET(":username", AuthMiddleware(), userController.GetUserByUsername)
		}

		tasks := api.Group("/tasks")
		tasks.Use(AuthMiddleware())
		{
			tasks.GET("/", taskController.GetAllTasks)
			tasks.GET(":id", taskController.GetTaskById)
			tasks.POST("/", taskController.CreateTask)
			tasks.PUT(":id", taskController.UpdateTask)
			tasks.DELETE(":id", taskController.DeleteTask)
		}
	}

	return r
}

func TestSetUpRoutes(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)

	mockUserController := new(MockUserController)
	mockTaskController := new(MockTaskController)

	router := SetUpTestRoutes(mockUserController, mockTaskController)

	t.Run("Test POST /api/users/register", func(t *testing.T) {
		mockUserController.On("Register", mock.Anything).Return()

		reqBody := `{"name": "testuser", "password": "password123"}`
		req, _ := http.NewRequest(http.MethodPost, "/api/users/register", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.JSONEq(t, `{"message": "User registered"}`, resp.Body.String())
		mockUserController.AssertCalled(t, "Register", mock.Anything)
	})

	t.Run("Test POST /api/users/login", func(t *testing.T) {
		mockUserController.On("Login", mock.Anything).Return()

		reqBody := `{"name": "testuser", "password": "password123"}`
		req, _ := http.NewRequest(http.MethodPost, "/api/users/login", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.JSONEq(t, `{"message": "User logged in"}`, resp.Body.String())
		mockUserController.AssertCalled(t, "Login", mock.Anything)
	})

	t.Run("Test GET /api/tasks/", func(t *testing.T) {
		mockTaskController.On("GetAllTasks", mock.Anything).Return()

		req, _ := http.NewRequest(http.MethodGet, "/api/tasks/", nil)

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.JSONEq(t, `{"message": "All tasks retrieved"}`, resp.Body.String())
		mockTaskController.AssertCalled(t, "GetAllTasks", mock.Anything)
	})

	t.Run("Test DELETE /api/tasks/:id", func(t *testing.T) {
		mockTaskController.On("DeleteTask", mock.Anything).Return()

		req, _ := http.NewRequest(http.MethodDelete, "/api/tasks/1", nil)

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.JSONEq(t, `{"message": "Task deleted"}`, resp.Body.String())
		mockTaskController.AssertCalled(t, "DeleteTask", mock.Anything)
	})
}
