package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/redis/go-redis/v9"
	"github.com/shegai01/msg/internal/config"
	"github.com/shegai01/msg/internal/handlers"
)

func main() {

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		slog.Info("info", "rdb.Ping(ctx).Result()", err)
		return
	}

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

	newHandl := handlers.NewHandlers(db, rdb)

	mux.HandleFunc("/health", handlers.Health)
	mux.HandleFunc("/users", newHandl.CreateUser())
	mux.HandleFunc("user/list", newHandl.ListUsers(ctx))

	server := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: mux,
	}

	slog.Info("info", "server started on port ", cfg.HTTPPort)
	// shut := make(chan os.Signal, 1)

	log.Fatal(server.ListenAndServe())

}
