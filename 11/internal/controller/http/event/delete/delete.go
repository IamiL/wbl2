package deleteEventHandler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type EventDeleter interface {
	DeleteEvent(ctx context.Context, eventID int64) error
}

type deleteEventRequestBody struct {
	EventID int64 `json:"event_id"`
}

func getUpdateEventHandler(log *slog.Logger, eventDeleter EventDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event deleteEventRequestBody

		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err = eventDeleter.DeleteEvent(r.Context(), event.EventID); err != nil {
			w.WriteHeader(200)
			return
		}

		w.WriteHeader(200)
	}
}
