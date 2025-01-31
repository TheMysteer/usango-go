package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// Estruturas para as respostas JSON
type RootResponse struct {
	Message string   `json:"message"`
	Routes  []string `json:"routes"`
}

type HealthResponse struct {
	Status string `json:"status"`
}

type TimeResponse struct {
	UTC  string `json:"utc_time"`
	Unix int64  `json:"unix_time"`
}

// Middleware para log de requisições
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, time.Since(start))
	}
}

// Handlers
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := RootResponse{
		Message: "Bem-vindo ao servidor Go!",
		Routes:  []string{"/health", "/time", "/echo?text=..."},
	}
	json.NewEncoder(w).Encode(response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HealthResponse{Status: "ok"})
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := TimeResponse{
		UTC:  time.Now().UTC().Format(time.RFC3339),
		Unix: time.Now().Unix(),
	}
	json.NewEncoder(w).Encode(response)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, `{"error": "Parâmetro 'text' não fornecido"}`, http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"echo": strings.ToUpper(text)})
}

func main() {
	// Registra os handlers com middleware de logging
	http.HandleFunc("/", loggingMiddleware(rootHandler))
	http.HandleFunc("/health", loggingMiddleware(healthHandler))
	http.HandleFunc("/time", loggingMiddleware(timeHandler))
	http.HandleFunc("/echo", loggingMiddleware(echoHandler))

	// Inicia o servidor
	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}