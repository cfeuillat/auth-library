package entity

import "errors"

// Custom errors
var (
    ErrUserNotFound     = errors.New("user not found")
    ErrInvalidPassword  = errors.New("invalid password")
    ErrEmailAlreadyUsed = errors.New("email already used")
    ErrDatabase         = errors.New("database error")
)

