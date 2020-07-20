package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Anjani100/go-login-api/models"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}	// UserID will be generated automatically by the database

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Username: input.Username, Password: input.Password}
	sqlStatement := `
		INSERT INTO User (Username, Password)
		VALUES ($1, $2)
		RETURNING id`
	id := 0
	err1 := models.DB.QueryRow(sqlStatement, user.Username, user.Password).Scan(&id)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("New record ID is:", id)
	c.JSON(http.StatusOK, gin.H{"data": user})
}