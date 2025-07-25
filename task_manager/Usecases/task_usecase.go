package usecases

import (
	domain "task_manager/Domain"
)

type TaskUseCase struct {
	TaskRepo domain.TaskRepository
}

func NewTaskUseCase(taskRepo domain.TaskRepository) *TaskUseCase {
	return &TaskUseCase{TaskRepo: taskRepo}
}

func (u *TaskUseCase) GetAllTasks() ([]*domain.Task, error) {
	return u.TaskRepo.GetAlltask()
}

func (u *TaskUseCase) GetTaskById(id string) (*domain.Task, error) {
	return u.TaskRepo.GetTaskById(id)
}

func (u *TaskUseCase) CreateTask(task *domain.Task) (*domain.Task, error) {
	return u.TaskRepo.CreateTask(task)
}

func (u *TaskUseCase) UpdateTask(id string, updated *domain.Task) (*domain.Task, error) {
	return u.TaskRepo.UpdateTask(id, updated)
}

func (u *TaskUseCase) DeleteTask(id string) error {
	return u.TaskRepo.DeleteTask(id)
}
