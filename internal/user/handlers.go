package user

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type Handlers struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewHandlers(postgres *sql.DB, redisclient *redis.Client) *Handlers {
	return &Handlers{
		db:  postgres,
		rdb: redisclient,
	}
}

type HealthResponse struct {
	Status string `json:"status"`
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthResponse{
		Status: "ok",
	})
}
