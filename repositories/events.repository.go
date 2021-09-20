package repositories

import (
	"github.com/yalm/cloud-messaging/models"
	"github.com/yalm/cloud-messaging/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

var eventRepository EventRepository

type EventRepository interface {
	InsertOne(event *models.Event) (*models.Event, error)
}

func createEventRepository(database *mongo.Database) EventRepository {
	collection := database.Collection("events")
	eventRepository = &connection{database: database, repo: collection}
	return eventRepository
}

func (connection *connection) InsertOne(event *models.Event) (*models.Event, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()

	_, err := connection.repo.InsertOne(ctx, event)

	if err != nil {
		return nil, err
	}

	return event, nil
}
