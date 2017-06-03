package utils

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func GetMuxPathIds(r *http.Request) map[string]bson.ObjectId {
	vars := mux.Vars(r)
	decodedVars := make(map[string]bson.ObjectId)
	for k, v := range vars {
		if bson.IsObjectIdHex(v) {
			decodedVars[k] = bson.ObjectIdHex(v)
		}
	}
	return decodedVars
}
