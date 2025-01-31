package main

import (
    "encoding/json"
    "fmt"         // Adicione esta linha!
    "net/http"
    "time"
)

type Response struct {
    Mensagem   string `json:"mensagem"`
    Horario    string `json:"horario"`
    Endpoint   string `json:"endpoint"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    resposta := Response{
        Mensagem: "Olá do servidor Go!",
        Horario:  time.Now().Format("2006-01-02 15:04:05"),
        Endpoint: r.URL.Path,
    }
    json.NewEncoder(w).Encode(resposta)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Servidor rodando em http://localhost:8080") // Agora funciona!
    http.ListenAndServe(":8080", nil)
}