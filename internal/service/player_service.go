package service

import (
	"football-analytics/internal/domain"
	"time"

	"github.com/google/uuid"
)

//interface business logic about football player
type PlayerService interface {
	CreatePlayer(name, position, teamID string, number int, birthday time.Time, height, weight float64) (*domain.Player, error)
	GetPlayerByID(id string) (*domain.Player, error)
	UpdatePlayer(id, name, position, teamID string, number int, birthday time.Time, height, weight float64) (*domain.Player, error)
	DeletePlayer(id string) error
	ListPlayers() ([]*domain.Player, error)
}

type playerService struct {
	playerRepo domain.PlayerRepository
}

// NewPlayerService create instance of PlayerService
func NewPlayerService(playerRepo domain.PlayerRepository) PlayerService {
	return &playerService{
		playerRepo: playerRepo,
	}
}

// CreatePlayer create new football player
func (s *playerService) CreatePlayer(name, position, teamID string, number int, birthday time.Time, height, weight float64) (*domain.Player, error) {
	player := &domain.Player{
		ID:        uuid.New().String(),
		Name:      name,
		Position:  position,
		TeamID:    teamID,
		Number:    number,
		Birthday:  birthday,
		Height:    height,
		Weight:    weight,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.playerRepo.Create(player); err != nil {
		return nil, err
	}

	return player, nil
}

// GetPlayerByID get football player by id
func (s *playerService) GetPlayerByID(id string) (*domain.Player, error) {
	return s.playerRepo.GetByID(id)
}

// UpdatePlayer update football player
func (s *playerService) UpdatePlayer(id, name, position, teamID string, number int, birthday time.Time, height, weight float64) (*domain.Player, error) {
	player, err := s.playerRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	player.Name = name
	player.Position = position
	player.TeamID = teamID
	player.Number = number
	player.Birthday = birthday
	player.Height = height
	player.Weight = weight
	player.UpdatedAt = time.Now()

	if err := s.playerRepo.Update(player); err != nil {
		return nil, err
	}

	return player, nil
}

// DeletePlayer delete football player
func (s *playerService) DeletePlayer(id string) error {
	return s.playerRepo.Delete(id)
}

// ListPlayers get all football player
func (s *playerService) ListPlayers() ([]*domain.Player, error) {
	return s.playerRepo.List()
}
