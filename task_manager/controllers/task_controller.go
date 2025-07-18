package controllers

import (
	"fmt"
	"net/http"

	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	registed, err := data.Register(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, registed)
}

func Login(ctx *gin.Context) {
	var input models.User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := data.Login(input)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func PromoteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	promote, err := data.PromoteUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, promote)
}

func GetAllTask(ctx *gin.Context) {
	tasks, err := data.GetAlltask()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve tasks"})
		return
	}
	fmt.Println(tasks)
	ctx.JSON(http.StatusOK, tasks)
}

func GetTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
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
	created, err := data.CreateTask(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create task"})
		return
	}
	ctx.JSON(http.StatusOK, created)
}

func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
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
	id := ctx.Param("id")
	deleted := data.DeleteTask(id)
	ctx.JSON(http.StatusOK, deleted)
}
