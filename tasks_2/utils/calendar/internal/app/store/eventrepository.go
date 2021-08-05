package store

import (
	"errors"

	"github.com/vlasove/materials/tasks_2/utils/calendar/internal/app/models"
)

// EventRepository ...
type EventRepository struct {
	store *Store
}

var (
	// BaseTimeSample ...
	BaseTimeSample        = "2006-01-02"
	errEventAlreadyExists = errors.New("event already exists")
	errEventDoesNotExists = errors.New("event does not exists")
	errBadDateParsing     = errors.New("can not parse date properly")
)

// CreateEvent ...
func (e *EventRepository) CreateEvent(event *models.Event) error {
	if !e.checkIfExists(event) {
		id := len(e.store.db) + 1
		e.store.db[id] = event
		return nil
	}
	return errEventAlreadyExists
}

// UpdateEvent ...
func (e *EventRepository) UpdateEvent(event *models.Event) error {
	if e.checkIfExists(event) {
		for id := range e.store.db {
			if id == event.ID {
				e.store.db[event.ID] = event
				return nil
			}
		}
	}
	return errEventDoesNotExists
}

// DeleteEvent ...
func (e *EventRepository) DeleteEvent(event *models.Event) error {
	if e.checkIfExists(event) {
		delete(e.store.db, event.ID)
		return nil
	}
	return errEventDoesNotExists
}

// GetEventsForDates ...
func (e *EventRepository) GetEventsForDates(dateLHS, dateRHS string) ([]*models.Event, error) {
	events := []*models.Event{}
	for _, val := range e.store.db {
		if val.Date <= dateRHS && val.Date >= dateLHS {
			events = append(events, val)
		}
	}
	return events, nil
}

func (e *EventRepository) checkIfExists(event *models.Event) bool {
	for _, val := range e.store.db {
		if *val == *event {
			return true
		}
	}
	return false
}
