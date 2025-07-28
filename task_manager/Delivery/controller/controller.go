package controller

import (
	"net/http"
	"time"

	"task_manager/Delivery/controller/dto"
	domain "task_manager/Domain"
	usecases "task_manager/Usecases"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUseCase domain.UserUseCaseInterface
}

type TaskController struct {
	TaskUseCase *usecases.TaskUseCase
}

func NewUserController(userUseCase domain.UserUseCaseInterface) *UserController {
	return &UserController{UserUseCase: userUseCase}
}

func NewTaskController(taskUseCase *usecases.TaskUseCase) *TaskController {
	return &TaskController{TaskUseCase: taskUseCase}
}

// User Handlers
func (uc *UserController) Register(c *gin.Context) {
	var userDto dto.UserDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := domain.User{
		Name:     userDto.Name,
		Password: userDto.Password,
		Role:     userDto.Role,
	}
	createdUser, err := uc.UserUseCase.Register(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdUser)
}

func (uc *UserController) Login(c *gin.Context) {
	var userDto dto.UserDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := domain.User{
		Name:     userDto.Name,
		Password: userDto.Password,
	}

	token, err := uc.UserUseCase.Login(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc *UserController) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := uc.UserUseCase.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) Promote(c *gin.Context) {
	id := c.Param("id")
	user, err := uc.UserUseCase.PromoteUser(id)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Task Handlers
func (tc *TaskController) GetAllTasks(c *gin.Context) {
	tasks, err := tc.TaskUseCase.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.TaskUseCase.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var taskDto dto.TaskDto
	if err := c.ShouldBindJSON(&taskDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task := domain.Task{}
	if task.DueDate.IsZero() {
		if dueDateStr, ok := c.GetPostForm("due_date"); ok {
			dueDate, err := time.Parse(time.RFC3339, dueDateStr)
			if err == nil {
				task.DueDate = dueDate
			}
		}
	}
	createdTask, err := tc.TaskUseCase.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdTask)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTask, err := tc.TaskUseCase.UpdateTask(id, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTask)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := tc.TaskUseCase.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
