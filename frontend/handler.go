package frontend

import (
	"net/http"
)

func ServeFrontend() {

	 fs := http.FileServer(http.Dir("./frontend/ui/static"))
	 http.Handle("/static/", http.StripPrefix("/static/", fs))
 
	 // Serve the index.html at the root (home page)
	 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		 http.ServeFile(w, r, "./frontend/ui/index.html")
	 })
	 
	 // Serve the login page
	 http.HandleFunc("/login-form", func(w http.ResponseWriter, r *http.Request) {
		 http.ServeFile(w, r, "./frontend/ui/login-form.html")
	 })
 
	 // Serve the register page
	 http.HandleFunc("/register-form", func(w http.ResponseWriter, r *http.Request) {
		 http.ServeFile(w, r, "./frontend/ui/register-form.html")
	 })
}