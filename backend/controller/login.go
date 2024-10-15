package controller

import (
    "fmt"
    "errors"
    "database/sql"
    "auth-library/backend/entity"
    "auth-library/backend/repository"
)

type LoginRequest struct{
    User entity.User
}

func LoginUser(db *sql.DB, loginRequest LoginRequest) (error) {
    user := loginRequest.User
    storedPassword, err := repository.FindUserPassword(db, user)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
			return entity.ErrUserNotFound
		}
        return fmt.Errorf("Error logging in user : %w", err)
    }

    if storedPassword != user.Password {
        return entity.ErrInvalidPassword
    }

    fmt.Println("User logged in successfully !")
    return nil
}
