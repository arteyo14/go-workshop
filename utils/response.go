package utils

import (
	"github.com/gin-gonic/gin"
)

func JSONResponse(c *gin.Context, code int, status bool, data interface{}, err interface{}) {

	if err != nil {
		if errorMessage, ok := err.(string); ok && errorMessage != "" {

			c.JSON(code, gin.H{
				"status":  false,
				"code":    code,
				"message": errorMessage,
			})
			return
		}
	}

	c.JSON(code, gin.H{
		"status": status,
		"code":   code,
		"data":   data,
	})
}
