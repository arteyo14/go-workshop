package utils

import (
	"github.com/gin-gonic/gin"
)

func JSONResponse(c *gin.Context, code int, status bool, data interface{}) {
	response := gin.H{
		"status": status,
		"code":   code,
	}

	// เช็คประเภทข้อมูลที่เป็น `string` เพื่อใช้เป็น `message`
	if message, ok := data.(string); ok {
		response["message"] = message
	} else {
		response["data"] = data
	}

	c.JSON(code, response)
}

func JSONErrorResponse(c *gin.Context, code int, err error) {
	if err != nil {
		JSONResponse(c, code, false, err.Error())
		return
	}
}
