package domain

import (
	"errors"
	"time"
)

// PlayerStats represents statistics for a player in a match
type PlayerStats struct {
	ID              int64     `json:"id"`
	PlayerID        int64     `json:"player_id"`
	MatchID         int64     `json:"match_id"`
	Minutes         int       `json:"minutes"`
	Goals           int       `json:"goals"`
	Assists         int       `json:"assists"`
	Shots           int       `json:"shots"`
	ShotsOnTarget   int       `json:"shots_on_target"`
	Passes          int       `json:"passes"`
	PassesCompleted int       `json:"passes_completed"`
	Tackles         int       `json:"tackles"`
	Interceptions   int       `json:"interceptions"`
	Fouls          int       `json:"fouls"`
	YellowCards    int       `json:"yellow_cards"`
	RedCards       int       `json:"red_cards"`
	PassAccuracy    float64   `json:"pass_accuracy"`
	ShotAccuracy    float64   `json:"shot_accuracy"`
	Date            time.Time `json:"date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// Validate checks if the player statistics are valid
func (s *PlayerStats) Validate() error {
	if s.PlayerID <= 0 {
		return errors.New("player ID is required")
	}
	if s.MatchID <= 0 {
		return errors.New("match ID is required")
	}
	if s.Minutes < 0 {
		return errors.New("minutes cannot be negative")
	}
	if s.Goals < 0 {
		return errors.New("goals cannot be negative")
	}
	if s.Assists < 0 {
		return errors.New("assists cannot be negative")
	}
	if s.ShotsOnTarget > s.Shots {
		return errors.New("shots on target cannot be greater than total shots")
	}
	if s.PassesCompleted > s.Passes {
		return errors.New("completed passes cannot be greater than total passes")
	}
	if s.YellowCards < 0 || s.RedCards < 0 {
		return errors.New("cards cannot be negative")
	}
	return nil
}

// Calculate computes derived statistics
func (s *PlayerStats) Calculate() {
	// Calculate pass accuracy
	if s.Passes > 0 {
		s.PassAccuracy = float64(s.PassesCompleted) / float64(s.Passes) * 100
	}

	// Calculate shot accuracy
	if s.Shots > 0 {
		s.ShotAccuracy = float64(s.ShotsOnTarget) / float64(s.Shots) * 100
	}
}