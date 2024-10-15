package handler

import(
	"auth-library/backend/entity"
	"net/http"
	"fmt"
	"encoding/json"
	"io"
	"errors"
)

func handleHTTPRequest(w http.ResponseWriter, r *http.Request) entity.User {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return entity.User{}
    }

    // Read and parse the request body
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Failed to read request body", http.StatusBadRequest)
        return entity.User{}
    }


    var user entity.User
        
    err = json.Unmarshal(body, &user)

    if err != nil {
        http.Error(w, "Invalid request body or missing fields", http.StatusBadRequest)
        return entity.User{}
    }

    return user
}

func handleErrors(w http.ResponseWriter, err error) {
	if errors.Is(err, entity.ErrUserNotFound) {
		http.Error(w, "User is not registered", http.StatusNotFound)
	} else if errors.Is(err, entity.ErrInvalidPassword) {
		http.Error(w, "Oopsie incorrect password", http.StatusUnauthorized)
	} else if errors.Is(err, entity.ErrEmailAlreadyUsed) {
		http.Error(w, "Already registered !!", http.StatusConflict)
	} else {
		http.Error(w, fmt.Sprint(err), http.StatusUnauthorized)
	}
}