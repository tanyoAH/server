package models

import "gopkg.in/mgo.v2/bson"

type BasicUserResponse struct {
	Id         bson.ObjectId `json:"id"`
	FullName   string        `json:"fullName"`
	ProfileUrl string        `json:"profileUrl"`
}
