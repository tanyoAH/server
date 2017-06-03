package models

import (
	"github.com/asaskevich/govalidator"
)

type CreateTripRequest struct {
	LocationName   string     `json:"locationName" valid:"length(1|250)"`
	Coordinates    MgoXY      `json:"coordinates"`
	TimePeriod     TimePeriod `json:"timePeriod"`
	ActivityBudget float64    `json:"activityBudget" valid:"required"`
	MaxDistance    float64    `json:"maxDistance" valid:"required"`
}

func (req *CreateTripRequest) Parameters() error {
	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		return err
	}
	return nil
}
