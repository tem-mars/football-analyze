package postgres

import (
	"football-analytics/internal/domain"
	"time"

	"github.com/jmoiron/sqlx"
)

type playerRepository struct {
	db *sqlx.DB
}

// NewPlayerRepository create repository for Player data
func NewPlayerRepository(db *sqlx.DB) domain.PlayerRepository {
	return &playerRepository{
		db: db,
	}
}

// Create add new Player data
func (r *playerRepository) Create(player *domain.Player) error {
	query := `
		INSERT INTO players (id, name, position, team_id, number, birthday, height, weight, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	
	_, err := r.db.Exec(
		query,
		player.ID,
		player.Name,
		player.Position,
		player.TeamID,
		player.Number,
		player.Birthday,
		player.Height,
		player.Weight,
		player.CreatedAt,
		player.UpdatedAt,
	)
	
	return err
}

func (r *playerRepository) GetByID(id string) (*domain.Player, error) {
	query := `
		SELECT id, name, position, team_id, number, birthday, height, weight, created_at, updated_at
		FROM players
		WHERE id = $1
	`
	
	var player domain.Player
	err := r.db.Get(&player, query, id)
	if err != nil {
		return nil, err
	}
	
	return &player, nil
}

func (r *playerRepository) Update(player *domain.Player) error {
	query := `
		UPDATE players
		SET name = $1, position = $2, team_id = $3, number = $4, birthday = $5, 
			height = $6, weight = $7, updated_at = $8
		WHERE id = $9
	`
	
	player.UpdatedAt = time.Now()
	
	_, err := r.db.Exec(
		query,
		player.Name,
		player.Position,
		player.TeamID,
		player.Number,
		player.Birthday,
		player.Height,
		player.Weight,
		player.UpdatedAt,
		player.ID,
	)
	
	return err
}

func (r *playerRepository) Delete(id string) error {
	query := `DELETE FROM players WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *playerRepository) List() ([]*domain.Player, error) {
	query := `
		SELECT id, name, position, team_id, number, birthday, height, weight, created_at, updated_at
		FROM players
	`
	
	var players []*domain.Player
	err := r.db.Select(&players, query)
	if err != nil {
		return nil, err
	}
	
	return players, nil
} 