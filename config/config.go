package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func LoadConfig() {
	// โหลดค่าจากไฟล์ .env.local ถ้ามี
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Error loading .env.local file: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if dbPort == "" {
		dbPort = "5432" // ค่า default ของ PostgreSQL port
	}

	// ตรวจสอบว่ามีการตั้งค่าค่าทั้งหมดหรือไม่
	if dbUser == "" || dbPass == "" || dbName == "" || dbHost == "" {
		log.Fatalf("Database configuration is missing.")
	}

	// สร้าง connection string ที่ถูกต้อง
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	DB, err = pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	log.Println("Connected to the database successfully!")
}
