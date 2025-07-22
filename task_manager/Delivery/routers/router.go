package routers

import (
	"github.com/gin-gonic/gin"
	"task_manager/Delivery/controller"
	"task_manager/Infrastructure"
)

func SetUpRoutes(userController *controller.UserController, taskController *controller.TaskController) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/register", userController.Register)
			users.POST("/login", userController.Login)
			users.POST("/promote/:id", Infrastructure.AuthMiddleware(), Infrastructure.Admin(), userController.Promote)
		}

		tasks := api.Group("/tasks")
		tasks.Use(Infrastructure.AuthMiddleware())
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
