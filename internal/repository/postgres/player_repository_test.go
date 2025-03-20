package postgres

import (
	"football-analytics/internal/domain"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PlayerRepositoryTestSuite struct {
	suite.Suite
	db         *sqlx.DB
	repository domain.PlayerRepository
	teamID     string
}

func (s *PlayerRepositoryTestSuite) SetupSuite() {
	db, err := NewConnection("postgres://postgres:postgres@localhost:5432/football_analytics?sslmode=disable")
	assert.NoError(s.T(), err)
	s.db = db
	
	s.repository = NewPlayerRepository(s.db)

	teamID := uuid.New().String()
	_, err = s.db.Exec(`
		INSERT INTO teams (id, name, country, league, logo) 
		VALUES ($1, $2, $3, $4, $5)
	`, teamID, "Test Team", "Test Country", "Test League", "test-logo.png")
	assert.NoError(s.T(), err)
	s.teamID = teamID
}

func (s *PlayerRepositoryTestSuite) TearDownSuite() {
	_, err := s.db.Exec("DELETE FROM players")
	assert.NoError(s.T(), err)
	_, err = s.db.Exec("DELETE FROM teams")
	assert.NoError(s.T(), err)

	// disconnect
	s.db.Close()
}

func (s *PlayerRepositoryTestSuite) TestCreatePlayer() {
	player := &domain.Player{
		ID:        uuid.New().String(),
		Name:      "Test Player",
		Position:  "Forward",
		TeamID:    s.teamID,
		Number:    10,
		Birthday:  time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Height:    180.5,
		Weight:    75.0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.repository.Create(player)
	assert.NoError(s.T(), err)

	var count int
	err = s.db.Get(&count, "SELECT COUNT(*) FROM players WHERE id = $1", player.ID)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 1, count)
}

func (s *PlayerRepositoryTestSuite) TestGetPlayerByID() {
	player := &domain.Player{
		ID:        uuid.New().String(),
		Name:      "Get Test Player",
		Position:  "Midfielder",
		TeamID:    s.teamID,
		Number:    8,
		Birthday:  time.Date(1992, 5, 15, 0, 0, 0, 0, time.UTC),
		Height:    175.0,
		Weight:    70.0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := s.db.Exec(`
		INSERT INTO players (id, name, position, team_id, number, birthday, height, weight, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`, player.ID, player.Name, player.Position, player.TeamID, player.Number, player.Birthday,
		player.Height, player.Weight, player.CreatedAt, player.UpdatedAt)
	assert.NoError(s.T(), err)

	result, err := s.repository.GetByID(player.ID)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), result)
	assert.Equal(s.T(), player.ID, result.ID)
	assert.Equal(s.T(), player.Name, result.Name)
	assert.Equal(s.T(), player.Position, result.Position)
}

func TestPlayerRepositorySuite(t *testing.T) {
	suite.Run(t, new(PlayerRepositoryTestSuite))
} 