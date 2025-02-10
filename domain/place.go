package domain

import (
	"errors"
	"time"
)

type Place struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

func (p *Place) Validate() error {
	if p.Name == "" {
		return errors.New("name cannot be empty")
	}
	if p.Latitude < -90 || p.Latitude > 90 {
		return errors.New("invalid latitude")
	}
	if p.Longitude < -180 || p.Longitude > 180 {
		return errors.New("invalid longitude")
	}
	return nil
}

func (p *Place) Update(updated Place) {
	p.Name = updated.Name
	p.Description = updated.Description
	p.Latitude = updated.Latitude
	p.Longitude = updated.Longitude
	p.UpdatedAt = time.Now()
}
