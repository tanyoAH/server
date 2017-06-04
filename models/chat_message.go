package models

import (
	"time"
)

type ChatMessage struct {
	User      BasicUserResponse `json:"user" bson:"user"`
	Text      string            `json:"text" bson:"text"`
	CreatedAt time.Time         `json:"createdAt" bson:"created_at"`
}
