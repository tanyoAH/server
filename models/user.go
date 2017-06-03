package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id                 bson.ObjectId `json:"id" bson:"_id"`
	AccessToken        string        `json:"-" bson:"access_token"`
	FullName           string        `json:"fullName" bson:"full_name"`
	Headline           string        `json:"headline" bson:"headline"`
	ProfileUrl         string        `json:"profileUrl" bson:"profile_url"`
	From               string        `json:"from" bson:"from"`
	Age                int64         `json:"age" bson:"age"`
	TravelledCountries []string      `json:"travelledCountries" bson:"travelled_countries"`
	Interests          []string      `json:"interests" bson:"interests"`
	// TODO oAuth profile link?, payment info
}

func (user *User) Create() error {
	if !user.Id.Valid() {
		user.Id = bson.NewObjectId()
	}
	return usersC.Insert(user)
}

func (user *User) FindByAccessToken() error {
	return usersC.Find(bson.M{"access_token": user.AccessToken}).One(user)
}

func (user *User) FindById() error {
	return usersC.FindId(user.Id).One(user)
}

func (user *User) ConvertToBasicUserResponse() *BasicUserResponse {
	return &BasicUserResponse{
		Id:         user.Id,
		FullName:   user.FullName,
		Headline:   user.Headline,
		ProfileUrl: user.ProfileUrl,
		From:       user.From,
	}
}
