package controllers

import (
	"net/http"

	"task_manager/service"

	"github.com/gin-gonic/gin"
)

func GetAllTask(ctx *gin.Context) {
	tasks := service.GetAllTask()
	ctx.JSON(http.StatusOK, tasks)
}
