package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func enableCORS(w http.ResponseWriter) {
	
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func sendJSONResponse(w http.ResponseWriter, msj string, statusCode int) {
    w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": msj}
    w.WriteHeader(statusCode)

    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
    }
}


func createHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}

	sendJSONResponse(w,  "Create endpoint hit!", http.StatusOK) 
}

// Handler for the "/get" route
func getHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}
	sendJSONResponse(w,  "Get endpoint hit!", http.StatusOK) 
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
