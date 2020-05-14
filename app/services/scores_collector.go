package services

import "board-game/app/models"

type (
	ScoresRepo interface {
		AddPoints(scores []models.Score) error
	}

	ScoreCollector struct {
		scoresRepo ScoresRepo
	}
)

func NewScoreCollector(r ScoresRepo) ScoreCollector {
	return ScoreCollector{
		scoresRepo: r,
	}
}

func (c ScoreCollector) CollectAndUpdate(scores []models.Score) error {
	return c.scoresRepo.AddPoints([]models.Score{})
}
