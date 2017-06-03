package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Vendor struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	// TODO - figure out tripadvisor integration
	TripAdvisorUrl string    `json:"tripAdvisorUrl" bson:"trip_advisor_url"`
	ThumbnailUrl   string    `json:"thumbnailUrl" bson:"thumbnail_url"`
	CreatedAt      time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt      time.Time `json:"updatedAt" bson:"updated_at"`
}

func (vendor *Vendor) Create() error {
	if !vendor.Id.Valid() {
		vendor.Id = bson.NewObjectId()
	}
	return vendorsC.Insert(vendor)
}

func (vendor *Vendor) FindById() error {
	return vendorsC.FindId(vendor.Id).One(vendor)
}
