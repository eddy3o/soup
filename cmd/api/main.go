package main

import (
	"log"
	"soup/internal/infraestructure"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	server := infraestructure.NewServer()
	server.Run()
}
