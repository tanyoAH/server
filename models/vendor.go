package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Vendor struct {
	Id          bson.ObjectId `json:"id" bson:"id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	// TODO - figure out tripadvisor integration
	TripAdvisorUrl string    `json:"tripAdvisorUrl" bson:"trip_advisor_url"`
	ThumbnailUrl   string    `json:"thumbnailUrl" bson:"thumbnail_url"`
	CreatedAt      time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt      time.Time `json:"updatedAt" bson:"updated_at"`
}
