package data

import (
	"task_manager/models"
	"time"

	"github.com/gin-gonic/gin"
)

var tasks = []models.Task{
	{
		ID:          1,
		Title:       "Complete Backend Assignment",
		Description: "Finish the backend learning path assignment by the end of the week.",
		DueDate:     time.Now().Add(72 * time.Hour), // 3 days from now
		Status:      "Pending",
	},
	{
		ID:          2,
		Title:       "Prepare for Meeting",
		Description: "Prepare slides and notes for the upcoming team meeting.",
		DueDate:     time.Now().Add(24 * time.Hour), // 1 day from now
		Status:      "In Progress",
	},
	{
		ID:          3,
		Title:       "Code Review",
		Description: "Review the pull requests submitted by the team.",
		DueDate:     time.Now().Add(48 * time.Hour), // 2 days from now
		Status:      "Pending",
	},
}

func GetAlltask(ctx *gin.Context) {
	return Task
}
