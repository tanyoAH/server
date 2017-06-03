package setup

import (
	"github.com/tanyoAH/tanyo-server/models"
	"gopkg.in/mgo.v2/bson"
)

func setupActivities() error {
	colombusLobster := models.Activity{
		VendorId:       "593332f371d675aff38c53b6",
		Name:           "Columbus Lobster Dinner Cruise in Cancun",
		Description:    "2.5-hour dinner cruise in Cancun",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=3911P70&d=5432465&aidSuffix=tvrm&partner=Viator",
		ThumbnailUrl:   "https://cache-graphicslib.viator.com/graphicslib/thumbs674x446/3911/SITours/columbus-lobster-dinner-cruise-in-cancun-in-cancun-376831.jpg",
		Price:          93,
		Location: models.MgoXY{
			X: 21.099722,
			Y: -86.765110,
		},
		IsEvening:        true,
		DurationHours:    2.5,
		Interests:        []string{"Food & Beverages", "Parties"},
		CommittedTripIds: []bson.ObjectId{},
		GroupChat:        []models.ChatMessage{},
	}
	err := colombusLobster.Create()
	if err != nil {
		return err
	}

	atvPlusZipline := models.Activity{
		VendorId:       "5933332c71d675aff38c53b7",
		Name:           "Cancun Combo Tour: ATV and Zipline with Cenote Swim",
		Description:    "Half-day zipline and ATV combo tour",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=3206CUNATVZIP&d=150807&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://cache-graphicslib.viator.com/graphicslib/thumbs674x446/3206/SITours/cancun-extreme-zipline-canopy-tour-in-cancun-43831.jpg",
		Price:          150,
		Location: models.MgoXY{
			X: 21.099722,
			Y: -86.765110,
		},
		IsEvening:        true,
		DurationHours:    4,
		Interests:        []string{"Extreme Sports"},
		CommittedTripIds: []bson.ObjectId{},
		GroupChat:        []models.ChatMessage{},
	}
	err = atvPlusZipline.Create()
	if err != nil {
		return err
	}

	swimWithWhaleSharks := models.Activity{
		VendorId:       "5933333771d675aff38c53b8",
		Name:           "Swim with Whale Sharks in Cancun: Small-Group Snorkeling Tour",
		Description:    "Whale shark snorkeling adventure off the coast of Cancun",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=3467WHALE&d=150807&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://cache-graphicslib.viator.com/graphicslib/thumbs674x446/3467/SITours/swim-with-whale-sharks-in-cancun-small-group-snorkeling-tour-in-cancun-171325.jpg",
		Price:          170,
		Location: models.MgoXY{
			X: 21.099722,
			Y: -86.765110,
		},
		IsEvening:        true,
		DurationHours:    2,
		Interests:        []string{"Extreme Sports"},
		CommittedTripIds: []bson.ObjectId{},
		GroupChat:        []models.ChatMessage{},
	}
	err = swimWithWhaleSharks.Create()
	if err != nil {
		return err
	}

	// TODO: add the rest and setup the field corrections on creation

	return nil
}
