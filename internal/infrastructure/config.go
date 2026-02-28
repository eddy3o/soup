package infrastructure

import (
	"net/http"
	"os"
	"soup/internal/router"
	"soup/internal/store"
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
	r := gin.Default()

	rds := store.NewRedis()

	router.RegisterRouteGroups(r, rds)

	httpServer := &http.Server{
		Addr:    ":" + strconv.Itoa(s.port),
		Handler: r,
	}

	return httpServer.ListenAndServe()
}
