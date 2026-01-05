package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type User struct {
	ID        int    `json:"id"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

//todo custom error, func contentType for delete to repeate code

func (h *Handlers) CreateUser() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var u User

		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err := h.db.Exec("insert into users (username, email) values ($1,$2)", u.UserName, u.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}

		w.WriteHeader(http.StatusCreated)
	}
}

func (h *Handlers) ListUsers(ctx context.Context) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		cached, err := h.rdb.Get(ctx, "users:list").Result()
		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(cached))
			return
		}

		rows, err := h.db.Query("select id, username,email, created_at from users")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		users := []User{}
		for rows.Next() {
			var u User
			if err := rows.Scan(&u.ID, &u.UserName, &u.CreatedAt, &u.Email); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			users = append(users, u)
		}

		data := json.NewEncoder(w).Encode(users)
		h.rdb.Set(ctx, "users:list", data, 3*time.Minute)
		w.Header().Set("Content-Type", "application/json")

	}
}
