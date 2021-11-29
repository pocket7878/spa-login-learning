package main

import (
	"log"
	"os"

	"github.com/pocket7878/spa_login_learning_backend/infrastructure/presentation"
	"github.com/pocket7878/spa_login_learning_backend/infrastructure/repository"
	"github.com/pocket7878/spa_login_learning_backend/usecase"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	userRepo, err := repository.NewUserRepository()
	if err != nil {
		log.Fatal("Failed to setup UserRepository")
	}
	userUsecase := usecase.NewUserUsecase(userRepo)

	todoRepo, err := repository.NewTodoRepository()
	if err != nil {
		log.Fatal("Failed to setup TodoRepository")
	}
	todoUsecase := usecase.NewTodoUsecase(todoRepo)

	r := presentation.NewRouter(userUsecase, todoUsecase)
	r.Run(":" + port)
}
