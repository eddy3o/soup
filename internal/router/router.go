package router

import (
	"soup/internal/auth"
	"soup/internal/categories"
	"soup/internal/middleware"
	"soup/internal/orders"
	"soup/internal/products"
	"soup/internal/store"
	"soup/internal/users"

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

	// User routes
	userApi := r.Group("/users")

	userRepo := users.NewRepository(db)
	userService := users.NewService(userRepo)
	userHandler := users.NewHandler(userService)

	users.RegisterRoutes(userApi, userHandler, middleware)

	// Product routes
	productsApi := r.Group("/products")

	productsRepo := products.NewRepository(db)
	productsService := products.NewService(productsRepo)
	productsHandler := products.NewHandler(productsService)

	products.RegisterRoutes(productsApi, productsHandler, middleware)

	// Category routes
	categoryApi := r.Group("/categories")

	categoryRepo := categories.NewRepository(db)
	categoryService := categories.NewService(categoryRepo)
	categoryHandler := categories.NewHandler(categoryService)

	categories.RegisterRoutes(categoryApi, categoryHandler, middleware)

	// Orders routes
	ordersApi := r.Group("/orders")

	ordersRepo := orders.NewRepository(db)
	ordersService := orders.NewService(ordersRepo)
	ordersHandler := orders.NewHandler(ordersService)

	orders.RegisterRoutes(ordersApi, ordersHandler, middleware)
}
