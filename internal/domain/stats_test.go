package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPlayerStats_Calculate(t *testing.T) {
	stats := PlayerStats{
		PlayerID:    1,
		MatchID:     1,
		Minutes:     90,
		Goals:       2,
		Assists:     1,
		Shots:       5,
		ShotsOnTarget: 3,
		Passes:      50,
		PassesCompleted: 45,
		Date:        time.Now(),
	}

	stats.Calculate()

	assert.Equal(t, float64(60), stats.PassAccuracy)
	assert.Equal(t, float64(60), stats.ShotAccuracy)
}

func TestPlayerStats_Validate(t *testing.T) {
	tests := []struct {
		name    string
		stats   PlayerStats
		wantErr bool
	}{
		{
			name: "valid stats",
			stats: PlayerStats{
				PlayerID:    1,
				MatchID:     1,
				Minutes:     90,
				Goals:       2,
				Assists:     1,
				Shots:       5,
				ShotsOnTarget: 3,
				Passes:      50,
				PassesCompleted: 45,
			},
			wantErr: false,
		},
		{
			name: "invalid stats - negative minutes",
			stats: PlayerStats{
				PlayerID:    1,
				MatchID:     1,
				Minutes:     -1,
			},
			wantErr: true,
		},
		{
			name: "invalid stats - shots on target > total shots",
			stats: PlayerStats{
				PlayerID:    1,
				MatchID:     1,
				Minutes:     90,
				Shots:       2,
				ShotsOnTarget: 3,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.stats.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}