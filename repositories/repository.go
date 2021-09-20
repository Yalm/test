package repositories

import "go.mongodb.org/mongo-driver/mongo"

type repository struct {
	dynamicLinkRepository DynamicLinkRepository
	eventRepository       EventRepository
}

type RepoRegistry interface {
	GetDynamicLinkRepository() DynamicLinkRepository
	GetEventRepository() EventRepository
}

func Create(database *mongo.Database) RepoRegistry {
	return repository{
		dynamicLinkRepository: createDynamicLinkRepository(database),
		eventRepository:       createEventRepository(database),
	}
}

func (r repository) GetDynamicLinkRepository() DynamicLinkRepository {
	return r.dynamicLinkRepository
}

func (r repository) GetEventRepository() EventRepository {
	return r.eventRepository
}
