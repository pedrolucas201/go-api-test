package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	http.HandleFunc("/health", loggingMiddleware(healthHandler))
	http.HandleFunc("/message", loggingMiddleware(messageHandler))

	log.Println("API rodando na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// GET /health
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]string{
		"status": "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// POST /message
func messageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(msg.Text) == "" {
		http.Error(w, "Campo 'text' não pode ser vazio", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}

// PLUS: middleware de log
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	}
}
