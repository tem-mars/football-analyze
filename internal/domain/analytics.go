package domain

type PerformanceMetrics struct {
	PlayerID           string  `json:"player_id"`
	GoalsPerMinute     float64 `json:"goals_per_minute"`
	AssistsPerMinute   float64 `json:"assists_per_minute"`
	PassAccuracy       float64 `json:"pass_accuracy"`
	ShotAccuracy       float64 `json:"shot_accuracy"`
	DefensiveEfficiency float64 `json:"defensive_efficiency"`
	Stamina            float64 `json:"stamina"`
	OverallRating      float64 `json:"overall_rating"`
}

type AnalyticsService interface {
	CalculatePlayerPerformance(playerID string, timeRange string) (*PerformanceMetrics, error)
	ComparePlayerPerformance(playerIDs []string) (map[string]*PerformanceMetrics, error)
	GetPlayerProgressOverTime(playerID string, startDate, endDate string) ([]*PerformanceMetrics, error)
	GetTeamPerformanceByPosition(teamID string) (map[string][]*PerformanceMetrics, error)
} 