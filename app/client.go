package app

import (
	"context"
	"go-mongo/app/config"
	"go-mongo/handler"
	"go-mongo/service"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}
func NewClient(ctx context.Context, config config.Config) (*handler.PlayerHandler, error) {
	client, ctx, _, err := connect("mongodb://localhost:27017")
	if err != nil {
		return nil, err
	}
	playerService := service.NewPLayerService((*mongo.Collection)(client))
	playerHandler := handler.NewPlayerHandler(playerService)
	return playerHandler, nil
}
