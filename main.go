package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := map[string]string{"status": "ok"}
	json.NewEncoder(w).Encode(resp)

	fmt.Printf("%s %s\n", r.Method, r.URL.Path)
}

func main() {
	http.HandleFunc("/api/health-check", healthCheckHandler)

	fmt.Println("Server running on port 3000...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
