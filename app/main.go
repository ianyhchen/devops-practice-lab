package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := getEnv("PORT", "8080")

	http.HandleFunc("/healthz", healthHandler)
	http.HandleFunc("/readyz", readyHandler)
	http.HandleFunc("/api/v1/items", itemsHandler)

	log.Printf("server started on port %s", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok")
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ready")
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]interface{}{
		"items": []string{"apple", "banana", "orange"},
		"time":  time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(resp)
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
