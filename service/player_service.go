package service

import (
	"context"
	"go-mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type playerService struct {
	Cl *mongo.Collection
}
type PlayerService interface {
	InsertPlayer(layer models.Player) (int64, error)
	GetAllPlayers() ([]models.Player, error)
	GetPlayer(id int64) (models.Player, error)
	UpdatePlayer(id int64, Player models.Player) (int64, error)
	DeletePlayer(id int64) (int64, error)
}

func NewPLayerService(cl *mongo.Collection) PlayerService {
	return &playerService{Cl: cl}
}
func (p *playerService) InsertPlayer(Player models.Player) (int64, error) {

	_, err := p.Cl.InsertOne(context.Background(), Player)
	if err != nil {
		return -1, err
	}
	return 1, nil
}

// get one Player from the p.DB by its Playerid
func (p *playerService) GetPlayer(id int64) (models.Player, error) {
	res := p.Cl.FindOne(context.Background(), id)
	if res.Err() != nil {
		return models.Player{}, res.Err()
	}
	player := models.Player{}
	err := res.Decode(&player)
	if err != nil {
		return models.Player{}, err
	}
	return player, nil
}

// get one Player from the p.DB by its Playerid
func (p *playerService) GetAllPlayers() ([]models.Player, error) {

	filter := bson.M{}
	cursor, er1 := p.Cl.Find(context.Background(), filter)
	if er1 != nil {
		return nil, er1
	}
	var players []models.Player
	er2 := cursor.All(context.Background(), &players)
	if er2 != nil {
		return nil, er2
	}
	return players, nil
}

// update Player in the p.DB
func (p *playerService) UpdatePlayer(id int64, Player models.Player) (int64, error) {

	filter := bson.M{"_id": Player.ID}
	update := bson.M{
		"$set": Player,
	}
	res, err := p.Cl.UpdateOne(context.Background(), filter, update)
	if res.ModifiedCount > 0 {
		return res.ModifiedCount, err
	} else if res.UpsertedCount > 0 {
		return res.UpsertedCount, err
	} else {
		return res.MatchedCount, err
	}
}

// delete Player in the p.DB
func (p *playerService) DeletePlayer(id int64) (int64, error) {

	filter := bson.M{"_id": id}
	res, err := p.Cl.DeleteOne(context.Background(), filter)
	if res == nil || err != nil {
		return 0, err
	}
	return res.DeletedCount, err
}
