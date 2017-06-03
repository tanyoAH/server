package models

import "gopkg.in/mgo.v2/bson"

type BasicUserResponse struct {
	Id         bson.ObjectId `json:"id"`
	FullName   string        `json:"fullName"`
	Headline   string        `json:"headline"`
	ProfileUrl string        `json:"profileUrl"`
	From       string        `json:"from"`
	Age        int64         `json:"age"`
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
