package main

import (
	"go-gin-workshop/config"
	// "log"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	// _ "github.com/lib/pq" // Import PostgreSQL driver
)

func main() {
	// เรียกใช้ LoadConfig เพื่อเชื่อมต่อฐานข้อมูล
	config.LoadConfig()

	// สร้าง instance ใหม่ของ Gin
	r := gin.Default()

	// สร้าง endpoint เพื่อทดสอบการเชื่อมต่อฐานข้อมูล
	r.GET("/test-db", func(c *gin.Context) {
		var result string

		// ทดสอบ query ข้อมูลจากฐานข้อมูล
		err := config.DB.QueryRow(context.Background(), "SELECT 'Database connected successfully!'").Scan(&result)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": result})
	})

	// รันเซิร์ฟเวอร์
	r.Run(":8080")
}
