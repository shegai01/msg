package main

import (
	"database/sql"
	"fmt"
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

	dns := fmt.Sprintf("host%s port=%s user=%s password=%s dbname=%s sslmode=dissable", cfg.DB_host, cfg.DB_Port, cfg.DB_user, cfg.DB_pass, cfg.DB_Name)

	db, err := sql.Open("postgres", dns)
	if err != nil {
		slog.Info("sql.open")
		return
	}
	defer db.Close()

	_, err = db.Exec(`	CREATE TABLE IF not EXISTS users(
		id serial PRIMARY key,
    username VARCHAR(50) not NULL UNIQUE,
    email VARCHAR(100) not NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT TIMESTAMP
	);`)

	if err != nil {
		slog.Info("err")
		return
	}

	mux.HandleFunc("/health", handlers.Health)
	mux.HandleFunc("/users", handlers.CreateUser(db))
	mux.HandleFunc("user/list", handlers.ListUsers(db))

	server := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: mux,
	}

	slog.Info("info", "server started on port ", cfg.HTTPPort)
	// shut := make(chan os.Signal, 1)

	log.Fatal(server.ListenAndServe())

}
