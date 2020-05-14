package services

import (
	"board-game/app/models"
	"context"
)

type (
	PubSubClient interface {
		// TODO: this method could be changed due to your needs
		// Publish must publish message to queue
		Publish(ctx context.Context, message []byte) error
	}

	ScorePublisher struct {
		pbClient PubSubClient
	}
)

func NewPointsAdder(pbClient PubSubClient) ScorePublisher {
	return ScorePublisher{
		pbClient: pbClient,
	}
}

func (s ScorePublisher) Publish(score models.Score) error {

}
