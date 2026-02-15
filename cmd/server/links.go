package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type Link struct {
	Name string
	Url  string
}

func getLinks(conn *pgx.Conn, ctx context.Context) ([]Link, error) {
	var links []Link

	rows, err := conn.Query(ctx, "SELECT name, url FROM links")
	if err != nil {
		fmt.Printf("query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var newLink Link
		err := rows.Scan(&newLink.Name, &newLink.Url)
		if err != nil {
			fmt.Printf("scan error: %v", err)
			return nil, err
		}

		links = append(links, newLink)
	}

	if rows.Err() != nil {
		fmt.Printf("rows error: %v", err)
		return nil, rows.Err()
	}

	return links, nil

}

func handleGetLinks(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	ctx := r.Context()

	links, err := getLinks(conn, ctx)
	if err != nil {
		fmt.Println("error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(links)

}
