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


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		rows, err := conn.Query(ctx, "SELECT name, url FROM links")
		if err != nil {
			fmt.Printf("query error : %v", err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var name string
			var url string
			err := rows.Scan(&name, &url)
			if err != nil {
				fmt.Printf("scan error: %v", err)
				return
			}
			fmt.Printf("name: %s, url: %s \n", name, url)
		}

		if rows.Err() != nil {
			fmt.Printf("rows error: %v", rows.Err())
			return
		}
	})
	http.ListenAndServe(":8000", nil)
}
