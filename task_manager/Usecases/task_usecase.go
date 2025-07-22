package usecases

import (
	domain "task_manager/Domain"
)

type TaskUseCaseImpl struct {
	TaskRepo domain.TaskRepository
}

func NewTaskUseCase(taskRepo domain.TaskRepository) *TaskUseCaseImpl {
	return &TaskUseCaseImpl{TaskRepo: taskRepo}
}

func (u *TaskUseCaseImpl) GetAllTasks() ([]*domain.Task, error) {
	return u.TaskRepo.GetAlltask()
}

func (u *TaskUseCaseImpl) GetTaskById(id string) (*domain.Task, error) {
	return u.TaskRepo.GetTaskById(id)
}

func (u *TaskUseCaseImpl) CreateTask(task *domain.Task) (*domain.Task, error) {
	return u.TaskRepo.CreateTask(task)
}

func (u *TaskUseCaseImpl) UpdateTask(id string, updated *domain.Task) (*domain.Task, error) {
	return u.TaskRepo.UpdateTask(id, updated)
}

func (u *TaskUseCaseImpl) DeleteTask(id string) error {
	return u.TaskRepo.DeleteTask(id)
}
