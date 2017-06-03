package api

import (
	"github.com/tanyoAH/tanyo-server/context"
	"github.com/tanyoAH/tanyo-server/models"
	"github.com/tanyoAH/tanyo-server/utils"
	"net/http"
)

func V0_GetMyProfile(w http.ResponseWriter, r *http.Request) {
	user, err := context.GetCurrentUser(r)
	if err != nil {
		utils.JSONForbiddenError(w, "Invalid user", "")
		return
	}

	utils.JSONSuccess(w, user, "Successfully returned user")
}

func V0_GetMyTrips(w http.ResponseWriter, r *http.Request) {
	user, err := context.GetCurrentUser(r)
	if err != nil {
		utils.JSONForbiddenError(w, "Invalid user", "")
		return
	}

	resp, err := models.GetMyTrips(user.Id)
	if err != nil {
		utils.JSONNotFoundError(w, "Trip not found", "")
		return
	}

	utils.JSONSuccess(w, resp, "Successfully returned trip")
}
