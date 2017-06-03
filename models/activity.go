package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Activity struct {
	TimePeriod       TimePeriod      `json:"timePeriod" bson:"time_period"`
	CommittedUserIds []bson.ObjectId `json:"committedUserIds" bson:"committed_user_ids"`
	GroupChat        []ChatMessage   `json:"groupChat" bson:"group_chat"`
	CreatedAt        time.Time       `json:"createdAt" bson:"created_at"`
	UpdatedAt        time.Time       `json:"updatedAt" bson:"updated_at"`
}
