package database

import (
	"context"
	"log"
	"time"
)

func (db *DB) RegisterUser(username, password, bio, email string, createdAt time.Time) error {

	_, err := db.Conn.Exec(
		context.Background(),
		"INSERT INTO users (user_name, user_password, user_email, user_bio, user_access_level, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		username, password, email, bio, 0, createdAt, createdAt,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (db *DB) SetSessionToken(username, sessionToken string) error {
	_, err := db.Conn.Exec(
		context.Background(), "UPDATE users SET session_token = $1 WHERE username = $2", sessionToken, username,
	)
	if err != nil {
		log.Println(err)
	}
	return nil
}
