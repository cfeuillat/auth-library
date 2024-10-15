package repository

import (
	"auth-library/backend/entity"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	// Initialize the database
	db, err := sql.Open("sqlite3", "./auth_db.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open the database: %w", err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT,
			"name" TEXT,
			"email" TEXT UNIQUE,
			"password" TEXT
		);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		db.Close() // Close connection if creating the table fails
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	fmt.Println("Database initialized!")
	return db, nil
}

func InsertUser(db *sql.DB, user entity.User) error {
	// SQL statement for inserting a new user
	insertUserSQL := `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`

	// Execute the query with the provided parameters (name, email)
	_, err := db.Exec(insertUserSQL, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil

}

func FindUserPassword(db *sql.DB, user entity.User) (string, error)  {
	// SQL statement to find the user's hashed password by email or name
	findPasswordSQL := `SELECT password FROM users WHERE email = ?`

	var storedPassword string

	// Execute the query to find the user by email
	err := db.QueryRow(findPasswordSQL, user.Email).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
		return "", fmt.Errorf("failed to query user: %w", err)
	}

	return storedPassword, nil
}
