package main

import "soup/internal/server"

func main() {
	server := server.NewServer()

	server.Run()
}
