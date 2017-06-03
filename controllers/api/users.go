package api

import (
	"github.com/tanyoAH/tanyo-server/models"
	"github.com/tanyoAH/tanyo-server/utils"
	"net/http"
)

func V0_GetUserProfile(w http.ResponseWriter, r *http.Request) {
	userId := utils.GetMuxPathIds(r)["userId"]
	if !userId.Valid() {
		utils.JSONBadRequestError(w, "Invalid user id", "")
		return
	}

	user := models.User{Id: userId}
	err := user.FindById()
	if err != nil {
		utils.JSONNotFoundError(w, "User not found", "")
		return
	}

	utils.JSONSuccess(w, user, "Successfully returned user")
}
