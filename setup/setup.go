package setup

import (
	"github.com/tanyoAH/tanyo-server/config"
	"github.com/tanyoAH/tanyo-server/models"
)

var Log = config.Conf.GetLogger()

func StartSetup(resetDatabase bool) {
	err := models.Setup()
	if err != nil {
		Log.WithField("error", err).Fatal("Couldn't setup models")
	}

	Log.Info("Models loaded")

	if resetDatabase {
		if config.Conf.IsProductionMode() {
			Log.Fatal("DB Reset is prohibited in production. Make sure you're in debug mode to reset the db.")
		}
		Log.Info("Resetting database")
		err = models.WipeDatabase()
		if err != nil {
			Log.WithField("error", err).Fatal("Couldn't wipe database for reset")
		}
		Log.Info("Finished resetting database")
	}

	err = insertData()
	if err != nil {
		Log.WithField("error", err).Fatal("Couldn't insert data into database")
	}
	Log.Infof("Finished setting up database")
}

func insertData() error {
	err := setupUsers()
	if err != nil {
		return err
	}

	err = setupVendors()
	if err != nil {
		return err
	}

	err = setupActivities()
	if err != nil {
		return err
	}

	return nil
}
