package controllers

import (
	"net/http"
	"strconv"

	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func GetAllTask(ctx *gin.Context) {
	tasks := data.GetAlltask()
	ctx.JSON(http.StatusOK, tasks)
}

func GetTaskById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
	}
	task, err := data.GetTaskById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "task not found"})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func CreateTask(ctx *gin.Context) {
	var task models.Task

	if err := ctx.ShouldBind(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created := data.CreateTask(task)
	ctx.JSON(http.StatusOK, created)
}

func UpdateTask(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var task models.Task
	if err := ctx.ShouldBind(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	updated, err := data.UpdateTask(id, task)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

func DeleteTask(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}
	deleted := data.DeleteTask(id)
	ctx.JSON(http.StatusOK, deleted)
}
