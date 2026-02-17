package main

import "pipeline/internal/server"

func main() {
	server := server.NewServer()

	server.Run()
}
