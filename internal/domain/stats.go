package domain

import (
	"time"
)

type PlayerMatchStats struct {
	ID              string    `json:"id"`
	PlayerID        string    `json:"player_id"`
	MatchID         string    `json:"match_id"`
	MinutesPlayed   int       `json:"minutes_played"`
	Goals           int       `json:"goals"`
	Assists         int       `json:"assists"`
	Passes          int       `json:"passes"`
	PassAccuracy    float64   `json:"pass_accuracy"` // percentage
	Shots           int       `json:"shots"`
	ShotsOnTarget   int       `json:"shots_on_target"`
	Tackles         int       `json:"tackles"`
	Interceptions   int       `json:"interceptions"`
	Fouls           int       `json:"fouls"`
	YellowCards     int       `json:"yellow_cards"`
	RedCards        int       `json:"red_cards"`
	DistanceCovered float64   `json:"distance_covered"` // in kilometers
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type PlayerMatchStatsRepository interface {
	Create(stats *PlayerMatchStats) error
	GetByID(id string) (*PlayerMatchStats, error)
	Update(stats *PlayerMatchStats) error
	Delete(id string) error
	ListByPlayerID(playerID string) ([]*PlayerMatchStats, error)
	ListByMatchID(matchID string) ([]*PlayerMatchStats, error)
	GetPlayerSeasonStats(playerID string, season string) (*PlayerSeasonStats, error)
}

type PlayerSeasonStats struct {
	PlayerID        string  `json:"player_id"`
	Season          string  `json:"season"`
	MatchesPlayed   int     `json:"matches_played"`
	MinutesPlayed   int     `json:"minutes_played"`
	Goals           int     `json:"goals"`
	Assists         int     `json:"assists"`
	PassAccuracy    float64 `json:"pass_accuracy"`
	ShotsOnTarget   int     `json:"shots_on_target"`
	TacklesPerGame  float64 `json:"tackles_per_game"`
	YellowCards     int     `json:"yellow_cards"`
	RedCards        int     `json:"red_cards"`
	DistanceCovered float64 `json:"distance_covered"`
} 