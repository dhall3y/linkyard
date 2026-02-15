package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Link struct {
	Name string
	Url  string
}

func (s *Server) getLinks(ctx context.Context) ([]Link, error) {
	var links []Link

	rows, err := s.DB.Query(ctx, "SELECT name, url FROM links")
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

// by adding (s *Server) we're saying this function belongs to the Server struct
func (s *Server) handleGetLinks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	links, err := s.getLinks(ctx)
	if err != nil {
		fmt.Println("error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(links)

}
