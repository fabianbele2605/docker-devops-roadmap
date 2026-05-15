package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `{"status": "ok"}`)
}

func main() {
	fmt.Println("Servidor iniciado con éxito en http://localhost:8080")
	http.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":8080", nil)
}