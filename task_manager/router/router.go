package router

import (
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes() *gin.Engine {
	r := gin.Default()
	taskRoutes := r.Group("api/tasks")
	{
		taskRoutes.GET("/", controllers.GetAllTask)
		taskRoutes.GET("/:id", controllers.GetTaskById)
		taskRoutes.POST("/", controllers.CreateTask)
		taskRoutes.PUT("/:id", controllers.UpdateTask)
		taskRoutes.DELETE("/:id", controllers.DeleteTask)
	}
	return r
}
