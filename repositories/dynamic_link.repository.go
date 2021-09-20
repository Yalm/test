package repositories

import (
	"github.com/yalm/cloud-messaging/models"
	"github.com/yalm/cloud-messaging/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var dynamicLinkRepository DynamicLinkRepository

type connection struct {
	database *mongo.Database
	repo     *mongo.Collection
}

type DynamicLinkRepository interface {
	FindByHostnameAndSuffix(host string, suffix string) (*models.DynamicLink, error)
}

func createDynamicLinkRepository(database *mongo.Database) DynamicLinkRepository {
	collection := database.Collection("dynamiclinks")
	dynamicLinkRepository = &connection{database: database, repo: collection}
	return dynamicLinkRepository
}

func (connection *connection) FindByHostnameAndSuffix(host string, suffix string) (*models.DynamicLink, error) {
	var dynamicLink models.DynamicLink

	ctx, cancel := utils.InitContext()
	defer cancel()

	err := connection.repo.FindOne(
		ctx,
		bson.M{"suffix.customSuffix": suffix, "dynamicLinkInfo.host": host, "status": 1}).Decode(&dynamicLink)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &dynamicLink, nil
}
