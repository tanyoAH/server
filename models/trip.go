package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Trip struct {
	Id             bson.ObjectId `json:"id" bson:"id"`
	UserId         bson.ObjectId `json:"-" bson:"user_id"`
	LocationName   string        `json:"locationName" bson:"locationName"`
	Coordinates    MgoXY         `json:"coordinates" bson:"coordinates"`
	TimePeriod     TimePeriod    `json:"timePeriod" bson:"time_period"`
	ActivityBudget float64       `json:"activityBudget" bson:"activity_budget"`
	MaxDistance    float64       `json:"maxDistance" bson:"max_distance"`
	CreatedAt      time.Time     `json:"createdAt" bson:"created_at"`
	UpdatedAt      time.Time     `json:"updatedAt" bson:"updated_at"`
}

func (trip *Trip) updateTS() {
	trip.UpdatedAt = time.Now()
}

func (trip *Trip) Create() error {
	trip.Id = bson.NewObjectId()
	trip.CreatedAt = time.Now()
	trip.updateTS()
	return tripsC.Insert(trip)
}

func (trip *Trip) CreateFromRequest(req *CreateTripRequest) error {
	trip.LocationName = req.LocationName
	trip.Coordinates = req.Coordinates
	trip.TimePeriod = req.TimePeriod
	trip.ActivityBudget = req.ActivityBudget
	trip.MaxDistance = req.MaxDistance
	return trip.Create()
}

func (trip *Trip) FindById() error {
	return tripsC.FindId(trip.Id).One(trip)
}

func (trip *Trip) GetDetailedResponse() (*DetailedTripResponse, error) {
	err := trip.FindById()
	if err != nil {
		return nil, err
	}

	u := User{Id: trip.UserId}
	err = u.FindById()
	if err != nil {
		return nil, err
	}

	activities, err := GetCommittedActivitiesForTrip(trip.Id)
	if err != nil {
		return nil, err
	}

	return &DetailedTripResponse{
		Trip:       *trip,
		User:       *u.ConvertToBasicUserResponse(),
		Activities: activities,
	}, nil
}
