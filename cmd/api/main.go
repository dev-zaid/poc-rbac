package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)


func main() {
	r := chi.NewRouter()

	// Define routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Chi Auth Server!"))
	})

	r.Post("/signup", signupHandler)
    r.Post("/login", loginHandler)

	// Start the server
	port := 8080
	fmt.Printf("Starting server on port %d...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
