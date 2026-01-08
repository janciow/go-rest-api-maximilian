package services

import (
	"errors"
	"go-test/internal/models"
	"go-test/internal/repository"
)

type EventService struct {
	repo *repository.EventRepository
}

func NewEventService(repo *repository.EventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) GetAllEvents() ([]models.Event, error) {
	if s.repo == nil {
		return nil, errors.New("repository not initialized")
	}
	return s.repo.GetAll()
}

func (s *EventService) GetEventByID(id int64) (*models.Event, error) {
	if id <= 0 {
		return nil, errors.New("invalid event ID")
	}
	if s.repo == nil {
		return nil, errors.New("repository not initialized")
	}
	return s.repo.GetByID(id)
}

func (s *EventService) CreateEvent(event *models.Event) error {
	if event == nil {
		return errors.New("event cannot be nil")
	}
	if event.Name == "" {
		return errors.New("event name is required")
	}
	if event.Description == "" {
		return errors.New("event description is required")
	}
	if event.Location == "" {
		return errors.New("event location is required")
	}
	if s.repo == nil {
		return errors.New("repository not initialized")
	}
	return s.repo.Save(event)
}

func (s *EventService) UpdateEvent(event *models.Event) error {
	if event == nil || event.ID <= 0 {
		return errors.New("invalid event")
	}
	if event.Name == "" {
		return errors.New("event name is required")
	}
	if s.repo == nil {
		return errors.New("repository not initialized")
	}
	return s.repo.Update(event)
}

func (s *EventService) DeleteEvent(id int64) error {
	if id <= 0 {
		return errors.New("invalid event ID")
	}
	if s.repo == nil {
		return errors.New("repository not initialized")
	}
	return s.repo.Delete(id)
}
