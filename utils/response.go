package utils

import (
	"github.com/gin-gonic/gin"
)

func JSONResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{
		"status": true,
		"code":   code,
		"data":   data,
	})
}

func JSONSuccessResponse(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"status":  true,
		"code":    code,
		"message": message,
	})
}

func JSONErrorResponse(c *gin.Context, code int, err error) {
	if err != nil {
		errMessage := err.Error()
		c.JSON(code, gin.H{
			"status":  false,
			"code":    code,
			"message": errMessage,
		})
		return
	}
}
