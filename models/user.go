package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id          bson.ObjectId `json:"id" bson:"id"`
	AccessToken string        `json:"-" bson:"access_token"`
	FullName    string        `json:"fullName" bson:"full_name"`
	ProfileUrl  string        `json:"profileUrl" bson:"profile_url"`
	Interests   []string      `json:"interests" bson:"interests"`
	// TODO oAuth profile link?, payment info
}

func (user *User) FindByAccessToken() error {
	return usersC.Find(bson.M{"access_token": user.AccessToken}).One(user)
}
