package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	dsn := os.Getenv("DATABASE_DSN")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(4 * time.Minute)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}
