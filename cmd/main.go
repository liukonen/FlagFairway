package main

import (
	"net/http"
)

func main() {
	// Serve static files (HTML, CSS, JS) from the "ui" directory
	fs := http.FileServer(http.Dir("../internal/ui/build"))
	http.Handle("/", fs)

	// Define API endpoints
	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		// Your API logic here
	})

	// Start server
	http.ListenAndServe(":8080", nil)
}
