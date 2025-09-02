package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := map[string]string{"status": "alive"}
	json.NewEncoder(w).Encode(resp)

	fmt.Printf("%s %s\n", r.Method, r.URL.Path)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	resp := map[string]string{"error": "not found"}
	json.NewEncoder(w).Encode(resp)

	fmt.Printf("404: %s %s\n", r.Method, r.URL.Path)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/health-check", healthCheckHandler)
	mux.HandleFunc("/", notFoundHandler)

	fmt.Println("Server running on port 3000...")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		panic(err)
	}
}
