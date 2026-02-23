package main

import "soup/internal/infraestructure"

func main() {
	server := infraestructure.NewServer()
	server.Run()
}
