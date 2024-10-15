package handler

import (
	"auth-library/backend/controller"
	"database/sql"
	"net/http"
)

func ServeRegisterUserAPI(db *sql.DB) {
    http.HandleFunc("/register", RegisterUserHandler(db))
}

func RegisterUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Mapping request
		user := handleHTTPRequest(w, r)

		registerRequest := controller.RegisterRequest{
			User: user,
		}

		// Business logic
		err := controller.RegisterUser(db, registerRequest)

		// Error handling
		if err != nil {
			 handleErrors(w, err)
			 return
		}

		// Success response
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User added successfully"))
	}
}
