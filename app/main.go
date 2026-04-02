package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ajayr/devops-cicd-eks-pipeline/app/internal/server"
)

func main() {
	cfg := server.Config{
		AppName:     valueOrDefault("APP_NAME", "demo-service"),
		Environment: valueOrDefault("APP_ENV", "dev"),
		Version:     valueOrDefault("APP_VERSION", "local"),
		Port:        valueOrDefault("PORT", "8080"),
		StartedAt:   time.Now().UTC(),
	}

	srv := server.New(cfg)

	log.Printf("starting %s on :%s", cfg.AppName, cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, srv.Routes()); err != nil {
		log.Fatal(err)
	}
}

func valueOrDefault(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

