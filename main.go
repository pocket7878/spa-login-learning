package main

import (
	"log"
	"os"

	"github.com/pocket7878/spa_login_learning_backend/infrastructure/presentation"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	r := presentation.NewRouter()
	r.Run(":" + port)
}
