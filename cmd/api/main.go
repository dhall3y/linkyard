package main

import (
	"context"
	"linkyard/internal/imports"
	"linkyard/internal/links"
	"linkyard/internal/server"
	"log"
	"net/http"
	"os"

	"github.com/gofrs/uuid"
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

	linksStore := links.NewStore(conn)
	linksHandler := links.NewHandler(linksStore)

	importsHandler := imports.NewHandler(linksStore, uuid.NewGen())

	server := server.NewServer(linksHandler, importsHandler)
	http.ListenAndServe(":8000", server)
}
