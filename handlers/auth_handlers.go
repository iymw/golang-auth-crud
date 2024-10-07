package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelwayne/go-crud/initializers"
	"github.com/michaelwayne/go-crud/models"
	"github.com/michaelwayne/go-crud/utils"
	"golang.org/x/crypto/bcrypt"
)

// Function for logging in
func Login(c *gin.Context) {
	var user models.User

	// Check user credentials and generate a JWT token
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Find the user in the database by username
	var dbUser models.User
	
	if err := initializers.DB.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare the hashed password with the input password
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate a JWT token
	token, err := utils.GenerateToken(dbUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Function for registering a new user (for demonstration purposes)
func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := initializers.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// func UsersIndex(c *gin.Context) {
// 	var users []models.User
// 	initializers.DB.Find(&users)

// 	c.JSON(200, gin.H{
// 		"user": users,
// 	})
// }

// func UsersDelete(c *gin.Context) {
// 	id := c.Param("id")

// 	initializers.DB.Delete(&models.User{}, id)

// 	c.Status(200)
// }
