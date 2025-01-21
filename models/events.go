package models

import (
	"event-planning/db"
	"time"
)

type Event struct {
	ID int64 `json:"id"`
	Title string `binding:"required" json:"title"`
	Description string `binding:"required" json:"description"`
	Venue string `binding:"required" json:"venue"`
	CreatedBy int64 `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
}

func (e Event) SaveEvents () error {
	query := db.InsertIntoEvents()
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.Title, e.Description, e.Venue, e.CreatedBy)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := db.GetAllEvents()
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Venue, &event.CreatedBy)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	if err = rows.Err(); err != nil {  
        return nil, err
    }  
	
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := db.GetEventByID()
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Title, &event.Description, &event.Venue, &event.CreatedBy)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) UpdateEventByID () error {
	query := db.UpdateEvent()
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(e.Title, e.Description, e.Venue, e.CreatedBy, e.ID)

	return err
}

func (e Event) DeleteEvent() error {
	query := db.DeleteEvent()
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	return err
}