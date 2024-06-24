package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
    testCases := []struct {
        name           string
        endpoint       string
        expectedStatus int
        expectedMessage string
    }{
        {"CreateHandler", "/create", http.StatusOK, "Create endpoint hit!"},
        {"GetHandler", "/get", http.StatusOK, "Get endpoint hit!"},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            req, err := http.NewRequest("GET", tc.endpoint, nil)
            if err != nil {
                t.Fatal(err)
            }

            rr := httptest.NewRecorder()

            // Determine handler based on endpoint
            var handler http.HandlerFunc
            switch tc.endpoint {
            case "/create":
                handler = createHandler
            case "/get":
                handler = getHandler
            default:
                t.Fatal("Invalid endpoint")
            }

            handler.ServeHTTP(rr, req)

            // Check status code
            if status := rr.Code; status != tc.expectedStatus {
                t.Errorf("handler returned wrong status code: got %v want %v",
                    status, tc.expectedStatus)
            }

            // Check Content-Type
            if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
                t.Errorf("handler returned wrong content type: got %v want %v",
                    contentType, "application/json")
            }

            // Parse and check JSON
            var response map[string]string
            if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
                t.Fatal(err)
            }

            if response["message"] != tc.expectedMessage {
                t.Errorf("handler returned unexpected body: got %v want %v",
                    response, tc.expectedMessage)
            }
        })
    }
}
