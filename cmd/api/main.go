package main

import (
	"log"
	"log/slog"
	"net/http"

	// "github.com/hashicorp/consul/agent/config
	"github.com/shegai01/msg/internal/config"
	"github.com/shegai01/msg/internal/handlers"
)

func main() {
	cfg := config.Load()
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.Health)

	server := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: mux,
	}

	slog.Info("info", "server started on port ", cfg.HTTPPort)
	// shut := make(chan os.Signal, 1)

	log.Fatal(server.ListenAndServe())

}
