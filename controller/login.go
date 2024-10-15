package controller

import (
    "fmt"
    "database/sql"
    "auth-library/entity"
    "auth-library/repository"
)

func LoginUser(db *sql.DB, user entity.User) (error) {
    err := repository.FindUser(db, user)
    if err != nil {
        return fmt.Errorf("Error logging in user : %w", err)
    }
    fmt.Println("User logged in successfully !")
    return nil
}
