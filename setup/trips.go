package setup

import (
	"github.com/tanyoAH/tanyo-server/models"
	"time"
)

func setupTrips() error {
	userA := models.User{AccessToken: "at_kim_a"}
	userB := models.User{AccessToken: "at_victor_b"}
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
		ActivityBudget: 330,
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

	bangkok := models.Trip{
		UserId:         userA.Id,
		LocationName:   "S31 Sukhumvit Hotel Soi Sukhumvit 31, Khlong Tan Nuea, Watthana, Bangkok, Thailand",
		TimePeriod:     models.TimePeriod{},
		ActivityBudget: 100,
		MaxDistance:    15,
		Coordinates: models.MgoXY{
			X: 13.733162,
			Y: 100.566089,
		},
	}
	err = bangkok.Create()
	if err != nil {
		return err
	}

	seoul := models.Trip{
		UserId:         userB.Id,
		LocationName:   "InterContinental Seoul Coex, Bongeunsa-ro, Seoul, South Korea",
		TimePeriod:     models.TimePeriod{},
		ActivityBudget: 200,
		MaxDistance:    5,
		Coordinates: models.MgoXY{
			X: 37.512939,
			Y: 127.057746,
		},
	}
	err = seoul.Create()
	if err != nil {
		return err
	}

	hk := models.Trip{
		UserId:         userB.Id,
		LocationName:   "W hotel, Austin Road West, West Kowloon",
		TimePeriod:     models.TimePeriod{},
		ActivityBudget: 150,
		MaxDistance:    5,
		Coordinates: models.MgoXY{
			X: 22.304804,
			Y: 114.160235,
		},
	}
	err = hk.Create()
	if err != nil {
		return err
	}

	singapore := models.Trip{
		UserId:         userB.Id,
		LocationName:   "Four Seaons, 190 Orchard Blvd, Singapore 248646",
		TimePeriod:     models.TimePeriod{},
		ActivityBudget: 130,
		MaxDistance:    5,
		Coordinates: models.MgoXY{
			X: 1.305149,
			Y: 103.828516,
		},
	}
	err = singapore.Create()
	if err != nil {
		return err
	}

	upcoming := models.Trip{
		LocationName:   "The Ritz-Carlton, Cancun, Calle Retorno del Rey #36, Zona Hotelera, 77500 Cancún, Q.R., Mexico",
		TimePeriod:     models.TimePeriod{},
		ActivityBudget: 600,
		MaxDistance:    5,
		Coordinates: models.MgoXY{
			X: 1.305149,
			Y: 103.828516,
		},
	}

	users, err := models.GetAllUsers()
	if err != nil {
		return err
	}
	for _, u := range users {
		if u.Id.String() == userA.Id.String() || u.Id.String() == userB.Id.String() {
			continue
		}
		i := models.Trip{
			LocationName: upcoming.LocationName,
			TimePeriod: models.TimePeriod{
				Start: time.Now().AddDate(0, 0, -1),
				End:   time.Now().AddDate(0, 0, 10),
			},
			ActivityBudget: 600,
			MaxDistance:    100,
			Coordinates: models.MgoXY{
				X: 1.305149,
				Y: 103.828516,
			},
		}
		err = i.Create()
		if err != nil {
			return err
		}
	}

	return nil
}
