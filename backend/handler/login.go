package handler

import (
	"auth-library/backend/controller"
	"database/sql"
	
	"net/http"
)

func ServeLoginUserAPI(db *sql.DB) {
    http.HandleFunc("/login", LoginUserHandler(db))
}

func LoginUserHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Mapping request
        user := handleHTTPRequest(w, r)
        loginRequest := controller.LoginRequest{
            User: user,
        }

        // Business logic
        err := controller.LoginUser(db, loginRequest)

        // Error handling
        if err != nil {
           handleErrors(w, err)
           return
        }

        // Success response
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("User logged in successfully"))
    }
}



