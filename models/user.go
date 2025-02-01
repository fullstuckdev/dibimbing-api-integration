package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Telephone string `json:"telephone"`
	Age int `json:"age"`
	Posts []Post `json:"posts" gorm:"foreignKey:UserID"`
}

type Post struct {
	gorm.Model
	Title string `json:"title"`
	Content string `json:"content"`
	UserID uint `json:"user_id"`
	User   User `json:"-" gorm:"foreignKey:UserID"`
}