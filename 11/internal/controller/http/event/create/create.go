package createEventHandler

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"
	"wbl2httpserver/internal/domain/models"
)

type EventCreater interface {
	CreateEvent(ctx context.Context, event *models.Event) (int64, error)
}

type createEventRequestBody struct {
	Uid   int64  `json:"uid"`
	Date  string `json:"date"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func getUpdateEventHandler(log *slog.Logger, eventCreater EventCreater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event createEventRequestBody

		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		date, err := time.Parse("2006-01-02", event.Date)
		if err != nil {
			fmt.Println("error parse uid")
		}

		id, err := eventCreater.CreateEvent(r.Context(), &models.Event{
			Uid:   event.Uid,
			Date:  date,
			Title: event.Title,
			Body:  event.Body,
		})
		if err != nil {
			w.WriteHeader(200)
			return
		}

		resp := map[string]string{"event_id": strconv.Itoa(int(id))}
		jsonData, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.Write(jsonData)

		w.WriteHeader(200)
	}
}
