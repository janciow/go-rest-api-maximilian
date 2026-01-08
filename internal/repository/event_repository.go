package repository

import (
	"database/sql"
	"errors"
	"go-test/internal/db"
	"go-test/internal/models"
)

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository() *EventRepository {
	return &EventRepository{db: db.DB}
}

func (r *EventRepository) Save(event *models.Event) error {
	if r.db == nil {
		return errors.New("database not initialized")
	}
	return event.Save()
}

func (r *EventRepository) GetAll() ([]models.Event, error) {
	if r.db == nil {
		return nil, errors.New("database not initialized")
	}
	return models.GetEvents()
}

func (r *EventRepository) GetByID(id int64) (*models.Event, error) {
	if r.db == nil {
		return nil, errors.New("database not initialized")
	}
	return models.GetEventByID(id)
}

func (r *EventRepository) Update(event *models.Event) error {
	if r.db == nil {
		return errors.New("database not initialized")
	}
	return event.Update()
}

func (r *EventRepository) Delete(id int64) error {
	if r.db == nil {
		return errors.New("database not initialized")
	}
	return models.DeleteEventByID(id)
}
