package main

import (
	"os"
	"github.com/joho/godotenv"
	"task_manager/Infrastructure/db"
	"task_manager/Infrastructure"
	"task_manager/Repositories"
	"task_manager/Usecases"
	"task_manager/Delivery/controller"
	"task_manager/Delivery/routers"
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

	passwordService := Infrastructure.NewPasswordService()
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "secret123"
	}
	jwtService := Infrastructure.NewJWTService(jwtSecret)

	userRepo := Repositories.NewUserRepository(userCollection)
	taskRepo := Repositories.NewTaskRepository(taskCollection)

	userUsecase := Usecases.NewUserUseCase(userRepo, passwordService, jwtService)
	taskUsecase := Usecases.NewTaskUseCase(taskRepo)

	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)

	r := routers.SetUpRoutes(userController, taskController)
	r.Run(":8080")
}
