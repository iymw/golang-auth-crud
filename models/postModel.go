package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string
	Body  string
} 

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}
