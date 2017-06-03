package models

import "gopkg.in/mgo.v2/bson"

type BasicActivityResponse struct {
	Id           bson.ObjectId `json:"id" bson:"id"`
	Name         string        `json:"name" bson:"name"`
	Price        float64       `json:"price" bson:"price"`
	ThumbnailUrl string        `json:"thumbnailUrl" bson:"thumbnail_url"`
	TimePeriod   TimePeriod    `json:"timePeriod" bson:"time_period"`
}
