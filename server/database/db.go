package database

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	_ = godotenv.Load(".env", "cmd/server/.env", "../.env")

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	err := mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: host,
	})
	if err != nil {
		log.Fatal("Failed to register TLS config:", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=tidb&charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	DB, err = gorm.Open(gormMysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to TiDB:", err)
	}

	log.Println("Successfully connected to TiDB Cloud!")
}
