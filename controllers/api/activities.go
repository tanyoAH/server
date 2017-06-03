package api

import (
	"github.com/tanyoAH/tanyo-server/context"
	"github.com/tanyoAH/tanyo-server/models"
	"github.com/tanyoAH/tanyo-server/utils"
	"net/http"
)

func V0_GetActivity(w http.ResponseWriter, r *http.Request) {
	user, err := context.GetCurrentUser(r)
	if err != nil {
		utils.JSONForbiddenError(w, "Invalid user", "")
		return
	}
	activityId := utils.GetMuxPathIds(r)["activityId"]
	if !activityId.Valid() {
		utils.JSONBadRequestError(w, "Invalid activity id", "")
		return
	}
	tripId := utils.GetMuxPathIds(r)["tripId"]
	if !tripId.Valid() {
		utils.JSONBadRequestError(w, "Invalid trip id", "")
		return
	}

	activity := models.Activity{Id: activityId}
	resp, err := activity.GetDetailedResponse(user.Id, tripId)
	if err != nil {
		utils.JSONNotFoundError(w, "Activity not found", "")
		return
	}

	utils.JSONSuccess(w, resp, "Successfully returned activity")
}

func V0_CommitToActivity(w http.ResponseWriter, r *http.Request) {
	user, err := context.GetCurrentUser(r)
	if err != nil {
		utils.JSONForbiddenError(w, "Invalid user", "")
		return
	}
	activityId := utils.GetMuxPathIds(r)["activityId"]
	if !activityId.Valid() {
		utils.JSONBadRequestError(w, "Invalid activity id", "")
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

	activity := models.Activity{Id: activityId}
	err = activity.UpdateCommitted(tripId)
	if err != nil {
		utils.JSONNotFoundError(w, "Activity not found", "")
		return
	}

	resp, err := activity.GetDetailedResponse(user.Id, trip.Id)
	if err != nil {
		utils.JSONNotFoundError(w, "Activity not found", "")
		return
	}

	utils.JSONSuccess(w, resp, "Successfully committed to activity")
}
