package domain

import (
    "time"
)

type Player struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    Position  string    `json:"position"`
    TeamID    string    `json:"team_id"`
    Number    int       `json:"number"`
    Birthday  time.Time `json:"birthday"`
    Height    float64   `json:"height"`
    Weight    float64   `json:"weight"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type PlayerRepository interface {
    Create(player *Player) error
    GetByID(id string) (*Player, error)
    Update(player *Player) error
    Delete(id string) error
    List() ([]*Player, error)
}