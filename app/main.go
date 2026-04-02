package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ajayr/devops-cicd-eks-pipeline/app/internal/server"
)

func main() {
	port := valueOrDefault("PORT", "8080")
	srv := server.New(server.Config{Port: port})

	log.Printf("starting service on :%s", srv.Port())
	if err := http.ListenAndServe(":"+srv.Port(), srv.Routes()); err != nil {
		log.Fatal(err)
	}
}

func valueOrDefault(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

