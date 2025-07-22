package routers

import (
	"task_manager/Delivery/controller"
	infrastructure "task_manager/Infrastructure"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(userController *controller.UserController, taskController *controller.TaskController) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/register", userController.Register)
			users.POST("/login", userController.Login)
			users.POST("/promote/:id", infrastructure.AuthMiddleware(), infrastructure.Admin(), userController.Promote)
			users.GET(":username", infrastructure.AuthMiddleware(), userController.GetUserByUsername)
		}

		tasks := api.Group("/tasks")
		tasks.Use(infrastructure.AuthMiddleware())
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
