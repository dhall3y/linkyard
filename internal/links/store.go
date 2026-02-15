package links

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Store struct {
	db *pgx.Conn
}

func NewStore(db *pgx.Conn) *Store {
	// this is a composite literal basically the same as:
	// var srv Server
	// srv.DB = conn
	// and then passing srv as a pointer with &srv
	return &Store{
		db: db,
	}
}

func (s *Store) getLinks(ctx context.Context) ([]Link, error) {
	var links []Link

	rows, err := s.db.Query(ctx, "SELECT name, url FROM links")
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
