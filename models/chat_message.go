package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type ChatMessage struct {
	UserId    bson.ObjectId `json:"userId" bson:"user_id"`
	Text      string        `json:"text" bson:"text"`
	CreatedAt time.Time     `json:"createdAt" bson:"created_at"`
}
