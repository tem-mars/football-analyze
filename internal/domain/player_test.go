package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPlayer_Validate(t *testing.T) {
	tests := []struct {
		name    string
		player  Player
		wantErr bool
	}{
		{
			name: "valid player",
			player: Player{
				ID:          1,
				FirstName:  "Cristiano",
				LastName:   "Ronaldo",
				Nationality: "Portugal",
				BirthDate:  time.Date(1985, 2, 5, 0, 0, 0, 0, time.UTC),
				Position:   "Forward",
				Height:     187,
				Weight:     83,
			},
			wantErr: false,
		},
		{
			name: "invalid player - empty first name",
			player: Player{
				ID:          1,
				FirstName:  "",
				LastName:   "Ronaldo",
				Nationality: "Portugal",
				Position:   "Forward",
			},
			wantErr: true,
		},
		{
			name: "invalid player - invalid position",
			player: Player{
				ID:          1,
				FirstName:  "Cristiano",
				LastName:   "Ronaldo",
				Nationality: "Portugal",
				Position:   "InvalidPosition",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.player.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPlayer_CalculateAge(t *testing.T) {
	now := time.Now()
	player := Player{
		BirthDate: now.AddDate(-25, 0, 0), // 25 years ago
	}

	assert.Equal(t, 25, player.CalculateAge())
}