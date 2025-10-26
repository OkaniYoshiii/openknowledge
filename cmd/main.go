package main

import (
	"database/sql"
	"fmt"
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

	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	address := os.Getenv("DATABASE_ADDRESS")
	protocol := os.Getenv("DATABASE_PROTOCOL")
	name := os.Getenv("DATABASE_NAME")

	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8", user, password, protocol, address, name)
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
