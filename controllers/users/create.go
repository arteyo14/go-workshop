package users

import (
	"context"
	"go-gin-workshop/config"
	response "go-gin-workshop/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Firstname string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Birthdate string    `json:"birthdate"`
	CreateAt  time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

func CreateUser(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		response.JSONErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	hashedPassword, err := response.HashPassword(user.Password)
	if err != nil {
		response.JSONErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	user.Password = hashedPassword
	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()

	query := "INSERT INTO users (username, password, first_name, last_name, email, birthdate, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"

	err = config.DB.QueryRow(context.Background(), query, user.Username, user.Password, user.Firstname, user.LastName, user.Email, user.Birthdate, user.CreateAt, user.UpdateAt).Scan(&user.ID)
	if err != nil {
		response.JSONErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.JSONResponse(c, http.StatusCreated, true, "Created user successfully!")
}
