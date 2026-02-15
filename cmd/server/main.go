package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

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

	var name string
	var url string

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		row := conn.QueryRow(ctx, "SELECT name, url FROM links")

		err := row.Scan(&name, &url)
		if err != nil {
			if err == pgx.ErrNoRows {
				fmt.Println("no rows in table links")
			}
			return
		}
		fmt.Println(name)
		fmt.Println(url)
	})
	http.ListenAndServe(":8000", nil)
}
