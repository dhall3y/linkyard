package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secret := os.Getenv("POSTGRES_PASSWORD")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(secret)
	})
	http.ListenAndServe(":8000", nil)
}
