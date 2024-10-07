package main

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelwayne/go-crud/controllers"
	"github.com/michaelwayne/go-crud/handlers"
	"github.com/michaelwayne/go-crud/initializers"
	"github.com/michaelwayne/go-crud/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// authentication
    publicRoutes := r.Group("/public")
    {
        publicRoutes.POST("/login", handlers.Login)
        publicRoutes.POST("/register", handlers.Register)
		// publicRoutes.GET("/users", handlers.UsersIndex)
		// publicRoutes.DELETE("/users/:id", handlers.UsersDelete)
    }

	// crud
	protectedRoutes := r.Group("/protected")
	protectedRoutes.Use(middleware.AuthenticationMiddleware())
	{
		// Protected routes here
		protectedRoutes.POST("/posts", controllers.PostsCreate)
		protectedRoutes.GET("/posts", controllers.PostsIndex)
		protectedRoutes.GET("/posts/:id", controllers.PostsShow)
		protectedRoutes.PUT("/posts/:id", controllers.PostsUpdate)
		protectedRoutes.DELETE("/posts/:id", controllers.PostsDelete)
	}

	r.Run()
}
