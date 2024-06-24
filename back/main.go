package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func enableCORS(w http.ResponseWriter) {
	
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins (replace "*" with specific origin in production)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}


// Handler for the "/create" route
func createHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}
	fmt.Fprintf(w, "Create endpoint hit!")
}

// Handler for the "/get" route
func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
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
