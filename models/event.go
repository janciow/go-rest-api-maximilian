package models

import "time"

type Event struct {
	ID          string    `json:"id",`
	Name        string    `json:"name"  binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
	UserID      string    `json:"user_id"`
}

var events = []Event{}

func (e *Event) Save() error {
	events = append(events, *e)
	return nil
}

func GetEvents() []Event {
	return events
}
