package data

import (
	"errors"
	"task_manager/models"
	"time"
)

var tasks = []models.Task{
	{
		ID:          1,
		Title:       "Complete Backend Assignment",
		Description: "Finish the backend learning path assignment by the end of the week.",
		DueDate:     time.Now().Add(72 * time.Hour),
		Status:      "Pending",
	},
	{
		ID:          2,
		Title:       "Prepare for Meeting",
		Description: "Prepare slides and notes for the upcoming team meeting.",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "In Progress",
	},
	{
		ID:          3,
		Title:       "Code Review",
		Description: "Review the pull requests submitted by the team.",
		DueDate:     time.Now().Add(48 * time.Hour),
		Status:      "Pending",
	},
}

func GetAlltask() []models.Task {
	return tasks
}

func GetTaskById(id int) (models.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, errors.New("task not found")

}

func CreateTask(task models.Task) models.Task {
	task.ID = getNextID()
	tasks = append(tasks, task)
	return task
}

func UpdateTask(id int, updated models.Task) (models.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			if updated.Title != "" {
				task.Title = updated.Title
			}
			if updated.Description != "" {
				task.Description = updated.Description
			}
			if updated.Status != "" {
				task.Status = updated.Status
			}
			return task, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func DeleteTask(id int) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}

	return errors.New("task not found")
}

func getNextID() int {
	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks)-1].ID + 1
}
