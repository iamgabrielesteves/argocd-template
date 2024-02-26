package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Define um logger para usar
	logger := log.New(os.Stdout, "myapp ", log.LstdFlags|log.Lmicroseconds)

	// Configura o handler para responder a todas as requisições HTTP com uma mensagem
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("Received request from %s\n", r.RemoteAddr)
		w.Write([]byte("App Funcional - UPDATE V2"))
	})

	// Inicia o servidor na porta 8080
	logger.Println("Iniciando servidor na porta 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatalf("Erro ao iniciar servidor: %s", err)
	}
}
