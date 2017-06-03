package models

import "time"

type Trip struct {
	LocationName   string     `json:"locationName" bson:"locationName"`
	Coordinates    MgoXY      `json:"coordinates" bson:"coordinates"`
	TimePeriod     TimePeriod `json:"timePeriod" bson:"time_period"`
	ActivityBudget float64    `json:"activityBudget" bson:"activity_budget"`
	MaxDistance    float64    `json:"maxDistance" bson:"max_distance"`
	CreatedAt      time.Time  `json:"createdAt" bson:"created_at"`
	UpdatedAt      time.Time  `json:"updatedAt" bson:"updated_at"`
}
