package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sirockin/cucumber-screenplay-go/back-end/internal/domain/application"
	httpserver "github.com/sirockin/cucumber-screenplay-go/back-end/internal/http"
)

func main() {
	var port = flag.Int("port", 8080, "port to run server on")
	flag.Parse()

	// Create domain application service
	appService := application.New()

	// Create HTTP server wrapping the service
	httpServer := httpserver.NewServer(appService)

	// Start server
	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Starting server on http://localhost%s", addr)
	log.Printf("API endpoints:")
	log.Printf("  POST   /accounts")
	log.Printf("  GET    /accounts/{name}")
	log.Printf("  POST   /accounts/{name}/activate")
	log.Printf("  POST   /accounts/{name}/authenticate")
	log.Printf("  GET    /accounts/{name}/authentication-status")
	log.Printf("  GET    /accounts/{name}/projects")
	log.Printf("  POST   /accounts/{name}/projects")
	log.Printf("  DELETE /clear")

	server := &http.Server{
		Addr:         addr,
		Handler:      httpServer,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
