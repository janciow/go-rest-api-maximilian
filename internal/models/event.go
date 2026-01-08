package models

import (
	"database/sql"
	"fmt"
	"time"

	"go-test/internal/db"
)

type Event struct {
	ID          int64     `json:"id"`
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

func GetEventByID(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`

	row := db.DB.QueryRow(query, id)
	var e Event
	err := row.Scan(&e.ID, &e.UserID, &e.Name, &e.Description, &e.Location, &e.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &e, nil
}

func (event Event) Update() error {
	query := `UPDATE events SET name = ?, description = ?, location = ?, date = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.Date, event.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEventByID(id int64) error {
	query := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}
