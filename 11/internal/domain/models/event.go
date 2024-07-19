package models

import "time"

type Event struct {
	Id    int64     `json:"id"`
	Uid   int64     `json:"uid"`
	Date  time.Time `json:"date"`
	Title string    `json:"title"`
	Body  string    `json:"body"`
}
