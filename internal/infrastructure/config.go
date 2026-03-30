package infrastructure

import (
	"fmt"
	"net/http"
	"os"
	"soup/internal/router"
	"soup/internal/store"
	"strconv"

	"github.com/gin-contrib/cors"
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
	r := gin.Default()
	r.Use(cors.Default())

	rds := store.NewRedis()
	db, err := store.NewDatabase()

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	router.RegisterRouteGroups(r, rds, db)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: r,
	}

	return httpServer.ListenAndServe()
}
