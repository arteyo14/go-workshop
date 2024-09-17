package users

import (
	"context"
	"database/sql"
	"go-gin-workshop/config"
	response "go-gin-workshop/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type GetUsersInterface struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Firstname string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Birthdate string    `json:"birthdate"`
	CreateAt  time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

func GetUsers(c *gin.Context) {
	var users []GetUsersInterface
	query := "SELECT id, username, first_name, last_name, email, birthdate, created_at, updated_at FROM users"

	rows, err := config.DB.Query(context.Background(), query)
	if err != nil {
		response.JSONErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var user GetUsersInterface
		var birthdate sql.NullString

		err := rows.Scan(&user.ID, &user.Username, &user.Firstname, &user.LastName, &user.Email, &birthdate, &user.CreateAt, &user.UpdateAt)
		if err != nil {
			response.JSONErrorResponse(c, http.StatusInternalServerError, err)
			return
		}

		if birthdate.Valid {
			user.Birthdate = birthdate.String
		} else {
			user.Birthdate = "" // กำหนดค่าเริ่มต้นหากเป็น NULL
		}

		users = append(users, user)
	}

	response.JSONResponse(c, http.StatusOK, true, users)
}
