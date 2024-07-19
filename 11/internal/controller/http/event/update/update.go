package updateEventHandler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"wbl2httpserver/internal/domain/models"
)

type EventUpdate interface {
	UpdateEvent(ctx context.Context, event *models.Event) error
}

func getUpdateEventHandler(log *slog.Logger, eventGetter EventUpdate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event models.Event

		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err = eventGetter.UpdateEvent(r.Context(), &event); err != nil {
			w.WriteHeader(200)
			return
		}

		w.WriteHeader(200)
	}
}
