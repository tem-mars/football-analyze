package service

import (
	"football-analytics/internal/domain"
	"time"

	"github.com/google/uuid"
)

// TeamService is interface for business logic about team
type TeamService interface {
	CreateTeam(name, country, league, logo string) (*domain.Team, error)
	GetTeamByID(id string) (*domain.Team, error)
	UpdateTeam(id, name, country, league, logo string) (*domain.Team, error)
	DeleteTeam(id string) error
	ListTeams() ([]*domain.Team, error)
}

type teamService struct {
	teamRepo domain.TeamRepository
}

// NewTeamService create instance of TeamService
func NewTeamService(teamRepo domain.TeamRepository) TeamService {
	return &teamService{
		teamRepo: teamRepo,
	}
}

// CreateTeam create new team
func (s *teamService) CreateTeam(name, country, league, logo string) (*domain.Team, error) {
	team := &domain.Team{
		ID:        uuid.New().String(),
		Name:      name,
		Country:   country,
		League:    league,
		Logo:      logo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.teamRepo.Create(team); err != nil {
		return nil, err
	}

	return team, nil
}

// GetTeamByID get team by id
func (s *teamService) GetTeamByID(id string) (*domain.Team, error) {
	return s.teamRepo.GetByID(id)
}

// UpdateTeam update team
func (s *teamService) UpdateTeam(id, name, country, league, logo string) (*domain.Team, error) {
	team, err := s.teamRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	team.Name = name
	team.Country = country
	team.League = league
	team.Logo = logo
	team.UpdatedAt = time.Now()

	if err := s.teamRepo.Update(team); err != nil {
		return nil, err
	}

	return team, nil
}

// DeleteTeam 
func (s *teamService) DeleteTeam(id string) error {
	return s.teamRepo.Delete(id)
}

// ListTeams
func (s *teamService) ListTeams() ([]*domain.Team, error) {
	return s.teamRepo.List()
} 