package utils

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

var ErrInvalidBSONObjectId = errors.New("Invalid bson object id")

func GetBSONObjectId(id string) (bson.ObjectId, error) {
	if !bson.IsObjectIdHex(id) {
		return "", ErrInvalidBSONObjectId
	}
	return bson.ObjectIdHex(id), nil
}

func BSONObjectArrayToStringArray(arr []bson.ObjectId) []string {
	resp := make([]string, len(arr))
	for ind := range arr {
		resp[ind] = arr[ind].String()
	}
	return resp
}
