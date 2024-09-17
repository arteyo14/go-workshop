package main

import (
	"go-gin-workshop/config"
	middleware "go-gin-workshop/middlewares"
	"go-gin-workshop/routes"

	// "log"

	"github.com/gin-gonic/gin"
	// _ "github.com/lib/pq" // Import PostgreSQL driver
)

func main() {
	// เรียกใช้ LoadConfig เพื่อเชื่อมต่อฐานข้อมูล
	config.LoadConfig()

	// สร้าง instance ใหม่ของ Gin
	r := gin.Default()

	r.Use(middleware.Header())

	routes.SetupRoute(r)

	// รันเซิร์ฟเวอร์
	r.Run(":8080")
}
