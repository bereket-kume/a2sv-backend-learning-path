package main

import (
	"os"
	"task_manager/Delivery/controller"
	"task_manager/Delivery/routers"
	infrastructure "task_manager/Infrastructure"
	"task_manager/Infrastructure/db"
	repositories "task_manager/Repositories"
	usecases "task_manager/Usecases"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	mongo, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	dbName := os.Getenv("DB_NAME")
	userCol := os.Getenv("USER_COLLECTION")
	taskCol := os.Getenv("TASK_COLLECTION")

	userCollection := mongo.GetCollection(dbName, userCol)
	taskCollection := mongo.GetCollection(dbName, taskCol)

	passwordService := infrastructure.NewPasswordService()
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "secret123"
	}
	jwtService := infrastructure.NewJWTService(jwtSecret)

	userRepo := repositories.NewUserRepository(userCollection)
	taskRepo := repositories.NewTaskRepository(taskCollection)

	userUsecase := usecases.NewUserUseCase(userRepo, passwordService, jwtService)
	taskUsecase := usecases.NewTaskUseCase(taskRepo)

	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)

	r := routers.SetUpRoutes(userController, taskController)
	r.Run(":8080")
}
