package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	LinkId      primitive.ObjectID `bson:"linkId" json:"linkId"`
	UtmSource   string             `bson:"utmSource,omitempty" json:"utmSource"`
	UtmMedium   string             `bson:"utmMedium,omitempty" json:"utmMedium"`
	UtmCampaign string             `bson:"utmCampaign,omitempty" json:"utmCampaign"`
	UtmTerm     string             `bson:"utmTerm,omitempty" json:"utmTerm"`
	UtmContent  string             `bson:"utmContent,omitempty" json:"utmContent"`
	CreatedAt   primitive.DateTime `bson:"createdAt" json:"createdAt"`
}
