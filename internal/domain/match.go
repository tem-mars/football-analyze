package domain

import (
	"time"
)

type Match struct {
	ID           string    `json:"id"`
	HomeTeamID   string    `json:"home_team_id"`
	AwayTeamID   string    `json:"away_team_id"`
	Date         time.Time `json:"date"`
	Venue        string    `json:"venue"`
	Competition  string    `json:"competition"`
	HomeScore    int       `json:"home_score"`
	AwayScore    int       `json:"away_score"`
	Status       string    `json:"status"` // scheduled, ongoing, completed, cancelled
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type MatchRepository interface {
	Create(match *Match) error
	GetByID(id string) (*Match, error)
	Update(match *Match) error
	Delete(id string) error
	List() ([]*Match, error)
	ListByTeamID(teamID string) ([]*Match, error)
	ListByDateRange(start, end time.Time) ([]*Match, error)
} 