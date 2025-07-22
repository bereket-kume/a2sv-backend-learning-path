package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	DueDate     time.Time          `bson:"duedate" json:"duedate"`
	Status      string             `bson:"status" json:"status"`
}

type User struct {
	ID       primitive.ObjectID `bson:"id" json:"id"`
	Name     string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Role     string             `bson:"role" json:"role"`
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

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)
