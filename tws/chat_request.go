package tws

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type ChatRequest struct {
	ActivityId string `json:"activityId"`
	Text       string `json:"text"`
}

func (req *ChatRequest) Parameters() error {
	if bson.IsObjectIdHex(req.ActivityId) {
		return errors.New("Invalid activity id")
	}
	return nil
}
