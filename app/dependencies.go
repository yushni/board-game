package app

import (
	"board-game/app/repositories"
	"board-game/app/services"
)

type Dependencies struct {
	SoresRepo   repositories.ScoresRepo
	PointsAdder services.ScorePublisher
}

func InitDependencies() *Dependencies {
	return &Dependencies{
		SoresRepo: repositories.NewScoresRepo(),
		//PointsAdder: service.NewPointsAdder(),
	}
}
