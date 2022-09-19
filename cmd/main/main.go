package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sk25469/go-mongodb-server/pkg/config"
	"github.com/sk25469/go-mongodb-server/pkg/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// database_name := "im_sahil"
	// database_password := "sahilsarwar"
	database_name := os.Getenv("DATABASE_USERNAME")
	database_password := os.Getenv("DATABASE_PASSWORD")

	if err := config.Connect(database_name, database_password); err != nil {
		log.Fatal(err)
	}

	routes.RegisterRoutes()

}
