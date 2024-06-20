package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Handler for the "/create" route
func createHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create endpoint hit!")
}

// Handler for the "/get" route
func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get endpoint hit!")
}

func main() {
	// Map routes to handlers
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/get", getHandler)

	// Get the port from environment variables, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Printf("Listening on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
