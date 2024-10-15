package controller

import (
    "fmt"
	"database/sql"
	"auth-library/backend/entity"
	"auth-library/backend/repository"
  "github.com/mattn/go-sqlite3"
)

type RegisterRequest struct{
  User entity.User
}


func RegisterUser(db *sql.DB, registerRequest RegisterRequest) (error) {
  user := registerRequest.User

  err := repository.InsertUser(db, user)

  if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
				return entity.ErrEmailAlreadyUsed
			}
    return fmt.Errorf("Error registering user : %w", err)
  }
}

  fmt.Println("User registered successfully !")
  return nil
}