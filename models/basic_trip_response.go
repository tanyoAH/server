package models

import "gopkg.in/mgo.v2/bson"

type BasicTripResponse struct {
	Id             bson.ObjectId `json:"id"`
	LocationName   string        `json:"locationName"`
	TimePeriod     TimePeriod    `json:"timePeriod"`
	ActivityBudget float64       `json:"activityBudget"`
	MaxDistance    float64       `json:"maxDistance"`
}

func GetMyTrips(userId bson.ObjectId) ([]BasicTripResponse, error) {
	rawTrips := []Trip{}
	err := tripsC.Find(bson.M{"user_id": userId}).All(rawTrips)
	if err != nil {
		return nil, err
	}

	return TripsToBasicTripResponse(rawTrips), nil
}

func TripsToBasicTripResponse(rawTrips []Trip) []BasicTripResponse {
	trips := make([]BasicTripResponse, len(rawTrips))
	for ind, rt := range rawTrips {
		rawTrips[ind] = rt
	}
	return trips
}
