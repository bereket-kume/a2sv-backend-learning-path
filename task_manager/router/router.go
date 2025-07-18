package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes() *gin.Engine {
	r := gin.Default()

	userRoutes := r.Group("api/users")
	{
		userRoutes.POST("/register", controllers.Register)
		userRoutes.POST("/login", controllers.Login)
		userRoutes.POST("/promote/:id", middleware.AuthMiddleware(), middleware.Admin(), controllers.PromoteUser)
	}

	taskRoutes := r.Group("api/tasks")
	taskRoutes.Use(middleware.AuthMiddleware())
	{
		taskRoutes.GET("/", controllers.GetAllTask)
		taskRoutes.GET("/:id", controllers.GetTaskById)
		taskRoutes.POST("/", controllers.CreateTask)
		taskRoutes.PUT("/:id", controllers.UpdateTask)
		taskRoutes.DELETE("/:id", controllers.DeleteTask)
	}

	return r
}
