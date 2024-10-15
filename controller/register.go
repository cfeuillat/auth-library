package controller

import (
    "fmt"
	"database/sql"
	"auth-library/entity"
	"auth-library/repository"
)


func RegisterUser(db *sql.DB, user entity.User) (error) {
  err := repository.InserUser(db, user)
  if err != nil {
    return fmt.Errorf("Error registering user : %w", err)
  }

  fmt.Println("User registered successfully !")
  return nil
}
