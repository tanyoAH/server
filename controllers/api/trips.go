package api

import (
	"github.com/tanyoAH/tanyo-server/context"
	"github.com/tanyoAH/tanyo-server/models"
	"github.com/tanyoAH/tanyo-server/utils"
	"net/http"
)

func V0_CreateTrip(w http.ResponseWriter, r *http.Request) {
	user, err := context.GetCurrentUserAndCatchForAPI(w, r)
	if err != nil {
		return
	}

	reqObj := models.CreateTripRequest{}
	err = utils.JSONDecodeAndCatchForAPI(w, r, &reqObj)
	if err != nil {
		return
	}

	trip := models.Trip{
		UserId:         user.Id,
		LocationName:   reqObj.LocationName,
		Coordinates:    reqObj.Coordinates,
		TimePeriod:     reqObj.TimePeriod,
		ActivityBudget: reqObj.ActivityBudget,
		MaxDistance:    reqObj.MaxDistance,
	}
	err = trip.Create()
	if err != nil {
		utils.JSONInternalError(w, "Failed to create trip", "")
		return
	}
	resp, err := trip.GetDetailedResponse()
	if err != nil {
		utils.JSONInternalError(w, "Failed to create trip", "")
		return
	}

	utils.JSONSuccess(w, resp, "Successfully created trip")
}

func V0_GetTrip(w http.ResponseWriter, r *http.Request) {
	tripId := utils.GetMuxPathIds(r)["tripId"]
	if !tripId.Valid() {
		utils.JSONBadRequestError(w, "Invalid trip id", "")
		return
	}

	trip := models.Trip{Id: tripId}
	resp, err := trip.GetDetailedResponse()
	if err != nil {
		utils.JSONNotFoundError(w, "Trip not found", "")
		return
	}

	utils.JSONSuccess(w, resp, "Successfully returned trip")
}

func V0_GetActivityRecommendationsForTrip(w http.ResponseWriter, r *http.Request) {
	user, err := context.GetCurrentUserAndCatchForAPI(w, r)
	if err != nil {
		return
	}
	tripId := utils.GetMuxPathIds(r)["tripId"]
	if !tripId.Valid() {
		utils.JSONBadRequestError(w, "Invalid trip id", "")
		return
	}

	trip := models.Trip{Id: tripId}
	err = trip.FindById()
	if err != nil {
		utils.JSONNotFoundError(w, "Trip not found", "")
		return
	}

	resp, err := models.GetRecommendedActivitiesForTrip(user, trip)
	if err != nil {
		utils.JSONNotFoundError(w, "Activities error", "")
		return
	}

	utils.JSONSuccess(w, resp, "Successfully returned activities")
}
