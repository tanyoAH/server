package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Activity struct {
	Id          bson.ObjectId `json:"id" bson:"id"`
	VendorId    bson.ObjectId `json:"vendorId" bson:"vendorId"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	// TODO - figure out tripadvisor integration
	TripAdvisorUrl   string          `json:"tripAdvisorUrl" bson:"trip_advisor_url"`
	Price            float64         `json:"price" bson:"price"`
	ThumbnailUrl     string          `json:"thumbnailUrl" bson:"thumbnail_url"`
	TimePeriod       TimePeriod      `json:"timePeriod" bson:"time_period"`
	CommittedTripIds []bson.ObjectId `json:"committedTripIds" bson:"committed_trip_ids"`
	GroupChat        []ChatMessage   `json:"groupChat" bson:"group_chat"`
	CreatedAt        time.Time       `json:"createdAt" bson:"created_at"`
	UpdatedAt        time.Time       `json:"updatedAt" bson:"updated_at"`
}

func GetCommittedActivitiesForTrip(tripId bson.ObjectId) ([]BasicActivityResponse, error) {
	respArr := []BasicActivityResponse{}
	err := activitiesC.Find(bson.M{"committed_trip_ids": bson.M{"$elemMatch": bson.M{"$eq": tripId}}}).All(&respArr)
	if err != nil {
		return nil, err
	}
	return respArr, nil
}
