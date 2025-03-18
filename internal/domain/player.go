package domain

import (
	"errors"
	"time"
)

// ValidPositions contains all valid football positions
var ValidPositions = map[string]bool{
	"Goalkeeper": true,
	"Defender":   true,
	"Midfielder": true,
	"Forward":    true,
}

// Player represents a football player entity
type Player struct {
	ID          int64     `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Nationality string    `json:"nationality"`
	BirthDate   time.Time `json:"birth_date"`
	Position    string    `json:"position"`
	Height      int       `json:"height"` // in centimeters
	Weight      int       `json:"weight"` // in kilograms
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Validate checks if the player data is valid
func (p *Player) Validate() error {
	if p.FirstName == "" {
		return errors.New("first name is required")
	}
	if p.LastName == "" {
		return errors.New("last name is required")
	}
	if p.Nationality == "" {
		return errors.New("nationality is required")
	}
	if !ValidPositions[p.Position] {
		return errors.New("invalid position")
	}
	if p.Height <= 0 {
		return errors.New("height must be positive")
	}
	if p.Weight <= 0 {
		return errors.New("weight must be positive")
	}
	return nil
}

// CalculateAge calculates the current age of the player
func (p *Player) CalculateAge() int {
	if p.BirthDate.IsZero() {
		return 0
	}
	now := time.Now()
	age := now.Year() - p.BirthDate.Year()
	
	// Adjust age if birthday hasn't occurred this year
	if now.Month() < p.BirthDate.Month() || 
		(now.Month() == p.BirthDate.Month() && now.Day() < p.BirthDate.Day()) {
		age--
	}
	return age
}