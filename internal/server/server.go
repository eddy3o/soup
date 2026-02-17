package server

import (
	"net/http"
	"os"
	"pipeline/internal/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port int
}

func NewServer() *Server {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}

	return &Server{
		port: port,
	}
}

func (s *Server) Run() error {
	router := gin.Default()

	api := router.Group("/api")

	authHandler := auth.NewHandler()

	auth.RegisterRoutes(api, authHandler)

	httpServer := &http.Server{
		Addr:    ":" + strconv.Itoa(s.port),
		Handler: router,
	}

	return httpServer.ListenAndServe()
}
