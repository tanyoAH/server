package tws

import (
	"github.com/tanyoAH/tanyo-server/config"
	"github.com/tanyoAH/tanyo-server/models"
	"github.com/tanyoAH/tanyo-server/twsproto"
	"github.com/tanyoAH/tanyo-server/utils"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func chatLogErrorIfDebugMode(err error) {
	if config.Conf.IsDebugMode() {
		Log.WithField("error", err).Error("Chat ws error")
	}
}

func chatHandler(sess *Session, r *twsproto.Request) {
	reqObj := ChatRequest{}
	err := twsproto.JSONDecodeAndCatch(sess.WriteChannel, r, &reqObj)
	if err != nil {
		return
	}

	activity := models.Activity{Id: bson.ObjectIdHex(reqObj.ActivityId)}
	err = activity.FindById()
	if err != nil {
		chatLogErrorIfDebugMode(err)
		twsproto.RespondError(sess.WriteChannel, r, "An error occurred sending the chat", http.StatusInternalServerError)
		return
	}

	cm, err := activity.AddChatMessage(*sess.User.ConvertToBasicUserResponse(), reqObj.Text)
	if err != nil {
		chatLogErrorIfDebugMode(err)
		twsproto.RespondError(sess.WriteChannel, r, "An error occurred sending the chat", http.StatusInternalServerError)
		return
	}

	userIds, err := models.GetUserIdsFromTripIdsArray(activity.CommittedTripIds)
	if err != nil {
		chatLogErrorIfDebugMode(err)
		twsproto.RespondError(sess.WriteChannel, r, "An error occurred sending the chat", http.StatusInternalServerError)
		return
	}

	payload := struct {
		ActivityId bson.ObjectId `json:"activityId"`
		models.ChatMessage
	}{ChatMessage: cm, ActivityId: activity.Id}

	sess.State.NotifyUsers(utils.BSONObjectArrayToStringArray(userIds), "activity.chat", payload)

	resp := struct {
		Ok bool `json:"ok"`
	}{Ok: true}

	twsproto.RespondSuccess(sess.WriteChannel, r, resp)
}
