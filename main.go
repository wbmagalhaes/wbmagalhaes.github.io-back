package main

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "pong")
}

func main() {
	http.HandleFunc("/api/ping", pingHandler)

	fmt.Println("Server running on port 3000...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
