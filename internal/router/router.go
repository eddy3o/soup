package router

import (
	"soup/internal/auth"
	"soup/internal/store"

	"github.com/gin-gonic/gin"
)

func RegisterRouteGroups(r *gin.Engine, rds *store.Redis, db *store.Database) {
	authApi := r.Group("/auth")

	authRepo := auth.NewRepository(rds, db)
	authService := auth.NewService(authRepo, rds)
	authHandler := auth.NewHandler(authService)

	auth.RegisterRoutes(authApi, authHandler)
}
