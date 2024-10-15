package handler

import(
	"net/http"
	"fmt"
	"database/sql"
	"auth-library/entity"
	"encoding/json"
	"auth-library/controller"
	"io"
)

func LoginUserHandler(db *sql.DB) http.HandlerFunc {
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

        if err != nil || user.Email == "" || user.Password == "" {
            http.Error(w, "Invalid request body or missing fields", http.StatusBadRequest)
            return
        }

        // Insert the user into the database
        err = controller.LoginUser(db, user)

        if err != nil {
            http.Error(w, fmt.Sprint(err), http.StatusUnauthorized)
            return
        }

        // Success response
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("User logged in successfully"))
    }
}
