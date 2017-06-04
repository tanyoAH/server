package setup

import "github.com/tanyoAH/tanyo-server/models"

func setupTrips() error {
	userA := models.User{AccessToken: "at_kim_a"}
	userB := models.User{AccessToken: "at_lucy_c"}
	err := userA.FindByAccessToken()
	if err != nil {
		return err
	}
	err = userB.FindByAccessToken()
	if err != nil {
		return err
	}

	sanFran := models.Trip{
		UserId:         userA.Id,
		LocationName:   "Hilton Hotel, 333 O'Farrell St, San Francisco, CA 94102, USA",
		TimePeriod:     models.TimePeriod{},
		ActivityBudget: 150,
		MaxDistance:    10,
		Coordinates: models.MgoXY{
			X: 37.785778,
			Y: -122.409578,
		},
	}
	err = sanFran.Create()
	if err != nil {
		return err
	}

	return nil
}
