package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Anjani100/go-login-api/models"
)

type CreateUserInput struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}	// UserID will be generated automatically by the database

type LoginUserInput struct {
	ID int
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}	// LoginUserID will be generated automatically by the database

func CreateUser(c *gin.Context) {
	user := CreateUserInput{}
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `
		INSERT INTO Users (Username, Password)
		VALUES ($1, $2)
		RETURNING id`
	id := 0
	row := models.DB.QueryRow(sqlStatement, user.Username, user.Password)
	err1 := row.Scan(&id)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists."})
		return
	}
	fmt.Println("New record ID is:", id)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func LoginUser(c *gin.Context) {
	user := LoginUserInput{}
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `SELECT * FROM Users where Username = ($1) AND PASSWORD = ($2)`
	row := models.DB.QueryRow(sqlStatement, user.Username, user.Password)
	err2 := row.Scan(&user.ID, &user.Username, &user.Password)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Username or Password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func RegisterForm(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

func LoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{})
}