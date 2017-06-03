package models

import "gopkg.in/mgo.v2/bson"

type BasicTripResponse struct {
	Id             bson.ObjectId `json:"id" bson:"_id"`
	LocationName   string        `json:"locationName" bson:"location_name"`
	TimePeriod     TimePeriod    `json:"timePeriod" bson:"time_period"`
	ActivityBudget float64       `json:"activityBudget" bson:"activity_budget"`
	MaxDistance    float64       `json:"maxDistance" bson:"max_distance"`
}

func GetMyTrips(userId bson.ObjectId) ([]BasicTripResponse, error) {
	rawTrips := []Trip{}
	err := tripsC.Find(bson.M{"user_id": userId}).All(&rawTrips)
	if err != nil {
		return nil, err
	}

	return TripsToBasicTripResponse(rawTrips), nil
}

func TripsToBasicTripResponse(rawTrips []Trip) []BasicTripResponse {
	trips := make([]BasicTripResponse, len(rawTrips))
	for ind, rt := range rawTrips {
		trips[ind] = BasicTripResponse{
			Id:             rt.Id,
			LocationName:   rt.LocationName,
			TimePeriod:     rt.TimePeriod,
			ActivityBudget: rt.ActivityBudget,
			MaxDistance:    rt.MaxDistance,
		}
	}
	return trips
}
