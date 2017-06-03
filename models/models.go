package models

import (
	"github.com/oschwald/geoip2-golang"
	"github.com/tanyoAH/tanyo-server/config"
	"gopkg.in/mgo.v2"
	"time"
)

var Log = config.Conf.GetLogger()
var db *mgo.Database
var usersC, tripsC, vendorsC, activitiesC *mgo.Collection
var geodb *geoip2.Reader

func Setup() error {
	mgoSess, err := mgo.DialWithTimeout(config.Conf.DBConnectionUrl, time.Second*3)
	if err != nil {
		return err
	}
	db = mgoSess.DB(config.Conf.DBName)

	usersC = db.C("users")
	tripsC = db.C("trips")
	vendorsC = db.C("vendors")
	activitiesC = db.C("activities")

	geodb, err = geoip2.Open(config.Conf.GeoLiteMMDBPath)
	if err != nil {
		return err
	}

	return nil
}

func WipeDatabase() error {
	// TODO
	return nil
}
