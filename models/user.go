package models

type User struct {
	UserID uint `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}