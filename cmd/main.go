package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/okaniyoshiii/openknowledge/internal/routes"
)

var address = flag.String("address", "127.0.0.1:8080", "tcp address the HTTP server will listen to")

func main() {
	flag.Parse()

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

	tmplDir := "./website/templates"
	baseTmpl := template.Must(template.New("base.html").ParseFiles(tmplDir + "/base.html"))

	config := routes.Config{}
	config.TemplateDir = tmplDir
	config.BaseTemplate = baseTmpl

	mux := http.NewServeMux()
	mux.Handle("GET /", &routes.PostsHandler{
		Config: config,
	})

	server := http.Server{
		Addr:              *address,
		Handler:           mux,
		ReadTimeout:       time.Second * 60,
		ReadHeaderTimeout: time.Second * 10,
		WriteTimeout:      time.Second * 60,
		IdleTimeout:       time.Second * 10,
		MaxHeaderBytes:    1 << 20, // 1MB
	}

	fmt.Printf("Listening on TCP address : %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
