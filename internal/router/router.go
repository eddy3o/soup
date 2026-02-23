package router

import (
	"soup/internal/auth"
	"soup/internal/store"

	"github.com/gin-gonic/gin"
)

func RegisterRouteGroups(r *gin.Engine, rds *store.Redis) {
	authApi := r.Group("/auth")
	authHandler := auth.NewHandler(rds)

	auth.RegisterRoutes(authApi, authHandler)
}
