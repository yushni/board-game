package services

import (
	"board-game/app/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreCollector_CollectAndUpdate(t *testing.T) {
	internalErr := errors.New("internal test error")

	table := map[string]struct {
		service   ScoreCollector
		argScores []models.Score
		expErr    error
	}{
		"can't add point: repo error": {
			service: ScoreCollector{scoresRepo: ScoresRepoMock{
				addPoints: func(scores []models.Score) error {
					return internalErr
				},
			}},
			expErr: internalErr,
		},
	}

	for name, tt := range table {
		tt := tt
		t.Run(name, func(t *testing.T) {
			gotErr := tt.service.CollectAndUpdate(tt.argScores)

			assert.Equal(t, tt.expErr, gotErr)
		})
	}
}

type ScoresRepoMock struct {
	addPoints func(scores []models.Score) error
}

func (m ScoresRepoMock) AddPoints(scores []models.Score) error {
	return m.addPoints(scores)
}
