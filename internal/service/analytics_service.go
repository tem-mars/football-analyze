package service

import (
	"football-analytics/internal/domain"
	"time"
)

type analyticsService struct {
	playerStatsRepo domain.PlayerMatchStatsRepository
	playerRepo      domain.PlayerRepository
	matchRepo       domain.MatchRepository
}

// NewAnalyticsService create instance of AnalyticsService
func NewAnalyticsService(
	playerStatsRepo domain.PlayerMatchStatsRepository,
	playerRepo domain.PlayerRepository,
	matchRepo domain.MatchRepository,
) domain.AnalyticsService {
	return &analyticsService{
		playerStatsRepo: playerStatsRepo,
		playerRepo:      playerRepo,
		matchRepo:       matchRepo,
	}
}

// CalculatePlayerPerformance calculate player performance in a specific time range
func (s *analyticsService) CalculatePlayerPerformance(playerID string, timeRange string) (*domain.PerformanceMetrics, error) {
	// get player data
	player, err := s.playerRepo.GetByID(playerID)
	if err != nil {
		return nil, err
	}

	// set time range
	var startDate, endDate time.Time
	endDate = time.Now()

	switch timeRange {
	case "week":
		startDate = endDate.AddDate(0, 0, -7)
	case "month":
		startDate = endDate.AddDate(0, -1, 0)
	case "season":
		// assume season start in August
		year := endDate.Year()
		if endDate.Month() < 8 {
			year--
		}
		startDate = time.Date(year, 8, 1, 0, 0, 0, 0, time.UTC)
	default: // "all"
		startDate = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	}

	// get match data in a specific time range
	matches, err := s.matchRepo.ListByDateRange(startDate, endDate)
	if err != nil {
		return nil, err
	}

	// create map of match IDs
	matchIDs := make(map[string]bool)
	for _, match := range matches {
		matchIDs[match.ID] = true
	}

	// get player stats
	allStats, err := s.playerStatsRepo.ListByPlayerID(playerID)
	if err != nil {
		return nil, err
	}

	// filter stats in a specific time range
	var filteredStats []*domain.PlayerMatchStats
	for _, stat := range allStats {
		if matchIDs[stat.MatchID] {
			filteredStats = append(filteredStats, stat)
		}
	}

	// calculate performance
	metrics := &domain.PerformanceMetrics{
		PlayerID: playerID,
	}

	if len(filteredStats) == 0 {
		return metrics, nil
	}

	var totalMinutes, totalGoals, totalAssists, totalPasses, totalShots, totalShotsOnTarget int
	var totalTackles, totalInterceptions int
	var totalPassAccuracy, totalDistance float64

	for _, stat := range filteredStats {
		totalMinutes += stat.MinutesPlayed
		totalGoals += stat.Goals
		totalAssists += stat.Assists
		totalPasses += stat.Passes
		totalShots += stat.Shots
		totalShotsOnTarget += stat.ShotsOnTarget
		totalTackles += stat.Tackles
		totalInterceptions += stat.Interceptions
		totalPassAccuracy += stat.PassAccuracy
		totalDistance += stat.DistanceCovered
	}

	// calculate average and ratio
	matchCount := float64(len(filteredStats))
	minutesPlayed := float64(totalMinutes)

	if minutesPlayed > 0 {
		metrics.GoalsPerMinute = float64(totalGoals) / minutesPlayed
		metrics.AssistsPerMinute = float64(totalAssists) / minutesPlayed
	}

	metrics.PassAccuracy = totalPassAccuracy / matchCount

	if totalShots > 0 {
		metrics.ShotAccuracy = float64(totalShotsOnTarget) / float64(totalShots)
	}

	metrics.DefensiveEfficiency = float64(totalTackles+totalInterceptions) / matchCount
	metrics.Stamina = totalDistance / matchCount

	// calculate overall rating (example calculation)
	metrics.OverallRating = (
		metrics.GoalsPerMinute*100 +
			metrics.AssistsPerMinute*50 +
			metrics.PassAccuracy*0.3 +
			metrics.ShotAccuracy*0.2 +
			metrics.DefensiveEfficiency*0.1 +
			metrics.Stamina*0.1) / 6

	return metrics, nil
}

// ComparePlayerPerformance compare player performance of multiple players
func (s *analyticsService) ComparePlayerPerformance(playerIDs []string) (map[string]*domain.PerformanceMetrics, error) {
	result := make(map[string]*domain.PerformanceMetrics)

	for _, playerID := range playerIDs {
		metrics, err := s.CalculatePlayerPerformance(playerID, "season")
		if err != nil {
			return nil, err
		}
		result[playerID] = metrics
	}

	return result, nil
}

// GetPlayerProgressOverTime get player progress over time
func (s *analyticsService) GetPlayerProgressOverTime(playerID string, startDateStr, endDateStr string) ([]*domain.PerformanceMetrics, error) {
	// convert date from string to time.Time
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return nil, err
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return nil, err
	}

	// get match data in a specific time range
	matches, err := s.matchRepo.ListByDateRange(startDate, endDate)
	if err != nil {
		return nil, err
	}

	// sort matches by date
	// (in this case, assume matches are sorted by date)

	// get player stats
	allStats, err := s.playerStatsRepo.ListByPlayerID(playerID)
	if err != nil {
		return nil, err
	}

	// create map of stats by match ID
	statsMap := make(map[string]*domain.PlayerMatchStats)
	for _, stat := range allStats {
		statsMap[stat.MatchID] = stat
	}

	// create progress data by interval (e.g. every 5 matches)
	var progressData []*domain.PerformanceMetrics
	const interval = 5
	var currentStats []*domain.PlayerMatchStats

	for i, match := range matches {
		if stat, ok := statsMap[match.ID]; ok {
			currentStats = append(currentStats, stat)
		}

		// calculate every interval matches or when reach last match
		if (i+1)%interval == 0 || i == len(matches)-1 {
			if len(currentStats) > 0 {
				// calculate performance from accumulated stats
				metrics := calculateMetricsFromStats(playerID, currentStats)
				progressData = append(progressData, metrics)
			}
		}
	}

	return progressData, nil
}

// GetTeamPerformanceByPosition get team performance by position
func (s *analyticsService) GetTeamPerformanceByPosition(teamID string) (map[string][]*domain.PerformanceMetrics, error) {
	// get player data in the team
	allPlayers, err := s.playerRepo.List()
	if err != nil {
		return nil, err
	}

	// filter players in the team and group by position
	playersByPosition := make(map[string][]string)
	for _, player := range allPlayers {
		if player.TeamID == teamID {
			playersByPosition[player.Position] = append(playersByPosition[player.Position], player.ID)
		}
	}

	// calculate player performance in each position
	result := make(map[string][]*domain.PerformanceMetrics)
	for position, playerIDs := range playersByPosition {
		var positionMetrics []*domain.PerformanceMetrics
		for _, playerID := range playerIDs {
			metrics, err := s.CalculatePlayerPerformance(playerID, "season")
			if err != nil {
				return nil, err
			}
			positionMetrics = append(positionMetrics, metrics)
		}
		result[position] = positionMetrics
	}

	return result, nil
}

// calculateMetricsFromStats calculate performance from stats
func calculateMetricsFromStats(playerID string, stats []*domain.PlayerMatchStats) *domain.PerformanceMetrics {
	metrics := &domain.PerformanceMetrics{
		PlayerID: playerID,
	}

	if len(stats) == 0 {
		return metrics
	}

	var totalMinutes, totalGoals, totalAssists, totalPasses, totalShots, totalShotsOnTarget int
	var totalTackles, totalInterceptions int
	var totalPassAccuracy, totalDistance float64

	for _, stat := range stats {
		totalMinutes += stat.MinutesPlayed
		totalGoals += stat.Goals
		totalAssists += stat.Assists
		totalPasses += stat.Passes
		totalShots += stat.Shots
		totalShotsOnTarget += stat.ShotsOnTarget
		totalTackles += stat.Tackles
		totalInterceptions += stat.Interceptions
		totalPassAccuracy += stat.PassAccuracy
		totalDistance += stat.DistanceCovered
	}

	// calculate average and ratio
	matchCount := float64(len(stats))
	minutesPlayed := float64(totalMinutes)

	if minutesPlayed > 0 {
		metrics.GoalsPerMinute = float64(totalGoals) / minutesPlayed
		metrics.AssistsPerMinute = float64(totalAssists) / minutesPlayed
	}

	metrics.PassAccuracy = totalPassAccuracy / matchCount

	if totalShots > 0 {
		metrics.ShotAccuracy = float64(totalShotsOnTarget) / float64(totalShots)
	}

	metrics.DefensiveEfficiency = float64(totalTackles+totalInterceptions) / matchCount
	metrics.Stamina = totalDistance / matchCount

	// calculate overall rating
	metrics.OverallRating = (
		metrics.GoalsPerMinute*100 +
			metrics.AssistsPerMinute*50 +
			metrics.PassAccuracy*0.3 +
			metrics.ShotAccuracy*0.2 +
			metrics.DefensiveEfficiency*0.1 +
			metrics.Stamina*0.1) / 6

	return metrics
} 