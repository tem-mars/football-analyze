package domain

import (
	"time"
)

type Team struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Country   string    `json:"country"`
	League    string    `json:"league"`
	Logo      string    `json:"logo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TeamRepository interface {
	Create(team *Team) error
	GetByID(id string) (*Team, error)
	Update(team *Team) error
	Delete(id string) error
	List() ([]*Team, error)
} 