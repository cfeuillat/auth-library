package main

import (
	"auth-library/backend/handler"
	"auth-library/backend/repository"
	"auth-library/frontend"
	"log"
	"net/http"
)


func main() {
	 // Initialize the database
	 db, err := repository.InitDB()
	 if err != nil {
		 log.Fatalf("Error initializing the database: %v", err)
	 }

	 defer db.Close()  // Defer closing until the end of the program

	 
	frontend.ServeFrontend()

	// API routes
	handler.ServeLoginUserAPI(db)
	handler.ServeRegisterUserAPI(db)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
