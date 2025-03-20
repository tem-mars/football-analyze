package service

import (
	"errors"
	"football-analytics/internal/domain"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPlayerRepository is mock for PlayerRepository
type MockPlayerRepository struct {
	mock.Mock
}

func (m *MockPlayerRepository) Create(player *domain.Player) error {
	args := m.Called(player)
	return args.Error(0)
}

func (m *MockPlayerRepository) GetByID(id string) (*domain.Player, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Player), args.Error(1)
}

func (m *MockPlayerRepository) Update(player *domain.Player) error {
	args := m.Called(player)
	return args.Error(0)
}

func (m *MockPlayerRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockPlayerRepository) List() ([]*domain.Player, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Player), args.Error(1)
}

func TestCreatePlayer(t *testing.T) {
	mockRepo := new(MockPlayerRepository)
	service := NewPlayerService(mockRepo)

	name := "Test Player"
	position := "Forward"
	teamID := uuid.New().String()
	number := 10
	birthday := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	height := 180.5
	weight := 75.0

	// set behavior of mock
	mockRepo.On("Create", mock.AnythingOfType("*domain.Player")).Return(nil)

	// call service
	player, err := service.CreatePlayer(name, position, teamID, number, birthday, height, weight)

	// check result
	assert.NoError(t, err)
	assert.NotNil(t, player)
	assert.Equal(t, name, player.Name)
	assert.Equal(t, position, player.Position)
	assert.Equal(t, teamID, player.TeamID)
	assert.Equal(t, number, player.Number)
	assert.Equal(t, birthday, player.Birthday)
	assert.Equal(t, height, player.Height)
	assert.Equal(t, weight, player.Weight)

	// check mock is called as expected
	mockRepo.AssertExpectations(t)
}

func TestGetPlayerByID(t *testing.T) {
	mockRepo := new(MockPlayerRepository)
	service := NewPlayerService(mockRepo)

	playerID := uuid.New().String()
	expectedPlayer := &domain.Player{
		ID:       playerID,
		Name:     "Test Player",
		Position: "Midfielder",
	}

	// case found data
	mockRepo.On("GetByID", playerID).Return(expectedPlayer, nil)
	player, err := service.GetPlayerByID(playerID)
	assert.NoError(t, err)
	assert.Equal(t, expectedPlayer, player)

	// case not found data
	notFoundID := uuid.New().String()
	mockRepo.On("GetByID", notFoundID).Return(nil, errors.New("player not found"))
	player, err = service.GetPlayerByID(notFoundID)
	assert.Error(t, err)
	assert.Nil(t, player)

	mockRepo.AssertExpectations(t)
}

func TestUpdatePlayer(t *testing.T) {
	mockRepo := new(MockPlayerRepository)
	service := NewPlayerService(mockRepo)

	playerID := uuid.New().String()
	existingPlayer := &domain.Player{
		ID:       playerID,
		Name:     "Old Name",
		Position: "Old Position",
		TeamID:   "old-team-id",
		Number:   9,
		Birthday: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Height:   175.0,
		Weight:   70.0,
	}

	newName := "New Name"
	newPosition := "New Position"
	newTeamID := uuid.New().String()
	newNumber := 10
	newBirthday := time.Date(1991, 2, 2, 0, 0, 0, 0, time.UTC)
	newHeight := 180.0
	newWeight := 75.0

	// set behavior of mock
	mockRepo.On("GetByID", playerID).Return(existingPlayer, nil)
	mockRepo.On("Update", mock.AnythingOfType("*domain.Player")).Return(nil)

	// call service
	updatedPlayer, err := service.UpdatePlayer(playerID, newName, newPosition, newTeamID, newNumber, newBirthday, newHeight, newWeight)

	// check result
	assert.NoError(t, err)
	assert.NotNil(t, updatedPlayer)
	assert.Equal(t, playerID, updatedPlayer.ID)
	assert.Equal(t, newName, updatedPlayer.Name)
	assert.Equal(t, newPosition, updatedPlayer.Position)
	assert.Equal(t, newTeamID, updatedPlayer.TeamID)
	assert.Equal(t, newNumber, updatedPlayer.Number)
	assert.Equal(t, newBirthday, updatedPlayer.Birthday)
	assert.Equal(t, newHeight, updatedPlayer.Height)
	assert.Equal(t, newWeight, updatedPlayer.Weight)

	mockRepo.AssertExpectations(t)
}

func TestDeletePlayer(t *testing.T) {
	mockRepo := new(MockPlayerRepository)
	service := NewPlayerService(mockRepo)

	playerID := uuid.New().String()

	// set behavior of mock
	mockRepo.On("Delete", playerID).Return(nil)

	// call service
	err := service.DeletePlayer(playerID)

	// check result
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestListPlayers(t *testing.T) {
	mockRepo := new(MockPlayerRepository)
	service := NewPlayerService(mockRepo)

	expectedPlayers := []*domain.Player{
		{ID: uuid.New().String(), Name: "Player 1"},
		{ID: uuid.New().String(), Name: "Player 2"},
	}

	// set behavior of mock
	mockRepo.On("List").Return(expectedPlayers, nil)

	// call service
	players, err := service.ListPlayers()

	// check result
	assert.NoError(t, err)
	assert.Equal(t, expectedPlayers, players)
	mockRepo.AssertExpectations(t)
} 