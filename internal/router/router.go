package router

import (
	"soup/internal/auth"
	"soup/internal/middleware"
	"soup/internal/products"
	"soup/internal/store"

	"github.com/gin-gonic/gin"
)

func RegisterRouteGroups(r *gin.Engine, rds *store.Redis, db *store.Database) {
	middleware := middleware.AuthMiddleware(rds)

	// Auth routes
	authApi := r.Group("/auth")

	authRepo := auth.NewRepository(rds, db)
	authService := auth.NewService(authRepo, rds)
	authHandler := auth.NewHandler(authService, rds)

	auth.RegisterRoutes(authApi, authHandler)

	// Product routes
	productsApi := r.Group("/products")

	productsRepo := products.NewRepository(db)
	productsService := products.NewService(productsRepo)
	productsHandler := products.NewHandler(productsService)

	products.RegisterRoutes(productsApi, productsHandler, middleware)
}
