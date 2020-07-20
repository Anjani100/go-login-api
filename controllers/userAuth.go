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

type LoginUserInput struct {
	ID int
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}	// LoginUserID will be generated automatically by the database

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `
		INSERT INTO Users (Username, Password)
		VALUES ($1, $2)
		RETURNING id`
	id := 0
	row := models.DB.QueryRow(sqlStatement, input.Username, input.Password)
	err1 := row.Scan(&id)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	fmt.Println("New record ID is:", id)
	c.JSON(http.StatusOK, gin.H{"data": input})
}

func LoginUser(c *gin.Context) {
	var input LoginUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sqlStatement := `SELECT * FROM Users where Username = ($1) AND PASSWORD = ($2)`
	row := models.DB.QueryRow(sqlStatement, input.Username, input.Password)
	err := row.Scan(&input.ID, &input.Username, &input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": input})
}
