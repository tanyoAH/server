package models

import "gopkg.in/mgo.v2/bson"

type BasicActivityResponse struct {
	Id           bson.ObjectId `json:"id"`
	ThumbnailUrl string        `json:"thumbnailUrl"`
	TimePeriod   TimePeriod    `json:"timePeriod"`
}
