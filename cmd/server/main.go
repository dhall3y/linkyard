package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type Server struct {
	DB *pgx.Conn
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to db: %v", err)
	}
	defer conn.Close(context.Background())

	// this is a composite literal basically the same as:
	// var srv Server
	// srv.DB = conn
	// and then passing srv as a pointer with &srv
	srv := &Server{
		DB: conn,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		srv.handleGetLinks(w, r)
	})

	http.ListenAndServe(":8000", nil)
}
