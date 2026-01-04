package models

import (
	"fmt"
	"time"

	"go-test/db"
)

type Event struct {
	ID          int64     `json:"id",`
	Name        string    `json:"name"  binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
	UserID      string    `json:"user_id"`
}

var events = []Event{}

func (e *Event) Save() error {
	query := `INSERT INTO events (user_id, name, description, location, date) VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.UserID, e.Name, e.Description, e.Location, e.Date)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id
	return err
}

func GetEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.UserID, &e.Name, &e.Description, &e.Location, &e.Date)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		events = append(events, e)
	}

	return events, nil
}
