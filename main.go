package main

import (
	"auth-library/handler"
	"auth-library/repository"
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

	// API routes
	http.HandleFunc("/register", handler.RegisterUserHandler(db))
	http.HandleFunc("/login", handler.LoginUserHandler(db))

	// Serve static files (CSS, JS, etc.)
	fs := http.FileServer(http.Dir("./ui/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve the index.html at the root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./ui/index.html")
	})

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
