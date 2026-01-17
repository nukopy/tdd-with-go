package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// rootHandler handles requests to the root path
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"status": "ok"}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error encoding response: %v\n", err)
	}
}

// NewServer creates and returns a new HTTP server
func NewServer(port int) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)

	addr := fmt.Sprintf(":%d", port)
	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}
