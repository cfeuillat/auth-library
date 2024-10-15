package handler

import (
	"auth-library/controller"
	"auth-library/entity"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func RegisterUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Read and parse the request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}

		var user entity.User
		err = json.Unmarshal(body, &user)

		if err != nil || user.Username == "" || user.Email == "" {
			http.Error(w, "Invalid request body or missing fields", http.StatusBadRequest)
			return
		}

		// Insert the user into the database
		err = controller.RegisterUser(db, user)

		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		// Success response
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User added successfully"))
	}
}
