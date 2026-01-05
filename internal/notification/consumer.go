package notification

import (
	"log/slog"

	"github.com/shegai01/msg/internal/shared"
)

func Start(events <-chan shared.UserCreatedEvent) {
	go func() {
		for e := range events {
			slog.String("[Notification] new user:", e.UserName)
			slog.String("[Notification] new email:", e.Email)
		}
	}()
}
