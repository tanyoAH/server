package models

import "gopkg.in/mgo.v2/bson"

type BasicUserResponse struct {
	Id         bson.ObjectId `json:"id"`
	FullName   string        `json:"fullName"`
	ProfileUrl string        `json:"profileUrl"`
}

func GetBasicUserResponsesFromIdArray(ids []bson.ObjectId) ([]BasicUserResponse, error) {
	users := []User{}
	err := usersC.Find(bson.M{"_id": bson.M{"$in": ids}}).All(&users)
	if err != nil {
		return nil, err
	}
	resp := make([]BasicUserResponse, len(users))
	for ind, u := range users {
		resp[ind] = *u.ConvertToBasicUserResponse()
	}
	return resp, nil
}
