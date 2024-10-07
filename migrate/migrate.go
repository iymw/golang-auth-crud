package main

import (
	"github.com/michaelwayne/go-crud/initializers"
	"github.com/michaelwayne/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	initializers.DB.AutoMigrate(&models.User{})
}
