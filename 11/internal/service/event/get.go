package eventGet

import (
	"context"
	"fmt"
	"time"
	"wbl2httpserver/internal/domain/models"
)

const (
	Day   = 1
	Week  = 2
	Month = 3
)

type EventGet struct {
	evProvider EventProvider
}

type EventProvider interface {
	EventsForDay(ctx context.Context, uid int64, date time.Time) *[]models.Event
	EventsForWeek(ctx context.Context, uid int64, date time.Time) *[]models.Event
	EventsForMonth(ctx context.Context, uid int64, date time.Time) *[]models.Event
}

func New(evtProvider EventProvider) *EventGet {
	return &EventGet{
		evProvider: evtProvider,
	}
}

func (e *EventGet) GetEvents(ctx context.Context, duration int8, uid int64, date time.Time) (*[]models.Event, error) {
	switch duration {
	case Day:
		return e.evProvider.EventsForDay(ctx, uid, date), nil
	case Week:
		return e.evProvider.EventsForWeek(ctx, uid, date), nil
	case Month:
		return e.evProvider.EventsForMonth(ctx, uid, date), nil
	}

	return nil, fmt.Errorf("error")
}
