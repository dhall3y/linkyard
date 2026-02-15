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

type Link struct {
	name string
	url  string
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


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		rows, err := conn.Query(ctx, "SELECT name, url FROM links")
		if err != nil {
			fmt.Printf("query error : %v", err)
			return
		}
		defer rows.Close()

		var links []Link

		for rows.Next() {
			var newLink Link
			err := rows.Scan(&newLink.name, &newLink.url)
			if err != nil {
				fmt.Printf("scan error: %v", err)
				return
			}
			links = append(links, newLink)
		}

		if rows.Err() != nil {
			fmt.Printf("rows error: %v", rows.Err())
			return
		}

		fmt.Printf("links: %+v", links)
	})
	http.ListenAndServe(":8000", nil)
}
