package repositories

import (
	"board-game/app/models"
	"errors"
)

type ScoresRepo struct {
}

func NewScoresRepo() ScoresRepo {
	return ScoresRepo{}
}

func (r ScoresRepo) All() ([]models.Score, error) {
	return nil, errors.New("not implemented")
}

func (r ScoresRepo) AddPoints(scores []models.Score) error {
	return errors.New("not implement")
}
