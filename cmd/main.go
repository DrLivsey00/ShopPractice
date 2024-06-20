package main

import (
	"context"
	"github.com/DrLivsey00/ShopPraactice/internal/server"
	"github.com/DrLivsey00/ShopPraactice/pkg/database"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	db, err := database.ConnectDB(dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}
	defer db.Conn.Close(context.Background())
	server.Run(db)
}
