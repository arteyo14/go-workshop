package routes

import (
	"go-gin-workshop/controllers/users"

	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine) {
	r.POST("/users", users.CreateUser)
}
