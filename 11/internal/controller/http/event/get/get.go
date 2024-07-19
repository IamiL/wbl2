package getEventHandler

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"
	"wbl2httpserver/internal/domain/models"
	eventGet "wbl2httpserver/internal/service/event"
)

type EventGet interface {
	GetEvents(ctx context.Context, duration int8, uid int64, date time.Time) (*[]models.Event, error)
}

func getEventForDayHandler(log *slog.Logger, eventGetter EventGet) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid, date, err := getQueryParameters(r)
		if err != nil {
			w.WriteHeader(200)
			return
		}

		events, err := eventGetter.GetEvents(r.Context(), eventGet.Day, uid, date)
		if err != nil {
			w.WriteHeader(200)
			return
		}

		data, err := json.Marshal(*events)
		if err != nil {
			w.WriteHeader(200)
			return
		}

		w.WriteHeader(200)
		w.Write(data)
	}
}

func getEventForWeekHandler(log *slog.Logger, eventGetter EventGet) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid, date, err := getQueryParameters(r)
		if err != nil {
			w.WriteHeader(200)
			return
		}

		events, err := eventGetter.GetEvents(r.Context(), eventGet.Week, uid, date)
		if err != nil {
			w.WriteHeader(200)
			return
		}

		data, err := json.Marshal(*events)
		if err != nil {
			w.WriteHeader(200)
			return
		}

		w.WriteHeader(200)
		w.Write(data)
	}
}

func getEventForMonthHandler(log *slog.Logger, eventGetter EventGet) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid, date, err := getQueryParameters(r)
		if err != nil {
			w.WriteHeader(200)
			return
		}

		events, err := eventGetter.GetEvents(r.Context(), eventGet.Month, uid, date)
		if err != nil {
			w.WriteHeader(200)
			return
		}

		data, err := json.Marshal(*events)
		if err != nil {
			w.WriteHeader(200)
			return
		}

		w.WriteHeader(200)
		w.Write(data)
	}
}

func getQueryParameters(r *http.Request) (int64, time.Time, error) {
	query := r.URL.Query()
	uidStr, present := query["user_id"]
	if !present || len(uidStr) != 1 {
		fmt.Println("filters not present")
	}

	dateStr, present := query["date"]
	if !present || len(dateStr) != 1 {
		fmt.Println("filters not present")
	}

	uid, err := strconv.ParseInt(uidStr[0], 10, 64)
	if err != nil {
		fmt.Println("error parse date")
	}

	date, err := time.Parse("2006-01-02", dateStr[0])
	if err != nil {
		fmt.Println("error parse uid")
	}

	return uid, date, nil
}
