package models

import "gopkg.in/mgo.v2/bson"

type BasicActivityResponse struct {
	Id               bson.ObjectId   `json:"id" bson:"_id"`
	Name             string          `json:"name" bson:"name"`
	Price            float64         `json:"price" bson:"price"`
	ThumbnailUrl     string          `json:"thumbnailUrl" bson:"thumbnail_url"`
	TimePeriod       TimePeriod      `json:"timePeriod" bson:"time_period"`
	CommittedTripIds []bson.ObjectId `json:"-" bson:"committed_trip_ids"`
	PeopleGoingCount int             `json:"peopleGoingCount" bson:"-"`
}

func (resp *BasicActivityResponse) Format() {
	resp.PeopleGoingCount = len(resp.CommittedTripIds)
}
