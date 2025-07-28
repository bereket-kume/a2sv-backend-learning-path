package domain

import "time"

type User struct {
	ID       string
	Name     string
	Password string
	Role     string
}

type Task struct {
	ID          string
	Title       string
	Description string
	DueDate     time.Time
	Status      string
}

type UserRepository interface {
	Register(user *User) (*User, error)
	GetUserByUsername(username string) (*User, error)
	PromoteUser(id string) (*User, error)
}

type TaskRepository interface {
	GetAlltask() ([]*Task, error)
	GetTaskById(id string) (*Task, error)
	CreateTask(task *Task) (*Task, error)
	UpdateTask(id string, updated *Task) (*Task, error)
	DeleteTask(id string) error
}

type UserUseCaseInterface interface {
	Register(user *User) (*User, error)
	Login(user *User) (string, error)
	GetUserByUsername(username string) (*User, error)
	PromoteUser(id string) (*User, error)
}

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)
