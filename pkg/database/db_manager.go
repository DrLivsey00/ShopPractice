package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

type DB struct {
	Conn *pgx.Conn
}

func ConnectDB(url string) (*DB, error) {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		return nil, err
	}
	log.Println("Database connection established")
	return &DB{Conn: conn}, nil
}
