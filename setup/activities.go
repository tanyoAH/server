package setup

import (
	"github.com/tanyoAH/tanyo-server/models"
	"gopkg.in/mgo.v2/bson"
)

func setupActivities() error {

	colombusLobster := models.Activity{
		VendorId:       bson.ObjectIdHex("593332f371d675aff38c53b6"),
		Name:           "Columbus Lobster Dinner Cruise in Cancun",
		Description:    "2.5-hour dinner cruise in Cancun",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=3911P70&d=5432465&aidSuffix=tvrm&partner=Viator",
		ThumbnailUrl:   "https://cache-graphicslib.viator.com/graphicslib/thumbs674x446/3911/SITours/columbus-lobster-dinner-cruise-in-cancun-in-cancun-376831.jpg",
		Price:          95,
		Location: models.MgoXY{
			X: 21.099722,
			Y: -86.765110,
		},
		IsEvening:        false,
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
		VendorId:       bson.ObjectIdHex("5933332c71d675aff38c53b7"),
		Name:           "Cancun Combo Tour: ATV and Zipline with Cenote Swim",
		Description:    "Half-day zipline and ATV combo tour",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=3206CUNATVZIP&d=150807&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://cache-graphicslib.viator.com/graphicslib/thumbs674x446/3206/SITours/cancun-extreme-zipline-canopy-tour-in-cancun-43831.jpg",
		Price:          143,
		Location: models.MgoXY{
			X: 21.099722,
			Y: -86.765110,
		},
		IsEvening:        false,
		DurationHours:    4,
		Interests:        []string{"Extreme Sports", "Diving"},
		CommittedTripIds: []bson.ObjectId{},
		GroupChat:        []models.ChatMessage{},
	}
	err = atvPlusZipline.Create()
	if err != nil {
		return err
	}

	swimWithWhaleSharks := models.Activity{
		VendorId:       bson.ObjectIdHex("5933333771d675aff38c53b8"),
		Name:           "Swim with Whale Sharks in Cancun: Small-Group Snorkeling Tour",
		Description:    "Whale shark snorkeling adventure off the coast of Cancun",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=3467WHALE&d=150807&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://cache-graphicslib.viator.com/graphicslib/thumbs674x446/3467/SITours/swim-with-whale-sharks-in-cancun-small-group-snorkeling-tour-in-cancun-171325.jpg",
		Price:          175,
		Location: models.MgoXY{
			X: 21.099722,
			Y: -86.765110,
		},
		IsEvening:        false,
		DurationHours:    2,
		Interests:        []string{"Extreme Sports", "Diving"},
		CommittedTripIds: []bson.ObjectId{},
		GroupChat:        []models.ChatMessage{},
	}
	err = swimWithWhaleSharks.Create()
	if err != nil {
		return err
	}

	tulumCave := models.Activity{
		VendorId:       bson.ObjectIdHex("5933334171d675aff38c53b9"),
		Name:           "Private Tour: Tulum and Cave Adventure from Cancun",
		Description:    "Full-day private tour of Tulum and Kantun-Chi from Cancun",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=7460P2&d=150807&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://cache-graphicslib.viator.com/graphicslib/thumbs674x446/7460/SITours/private-tour-tulum-and-cave-adventure-from-cancun-in-cancun-193568.jpg",
		Price:          143,
		Location: models.MgoXY{
			X: 21.099722,
			Y: -86.765110,
		},
		IsEvening:        false,
		DurationHours:    5,
		Interests:        []string{"Extreme Sports", "Caving"},
		CommittedTripIds: []bson.ObjectId{},
		GroupChat:        []models.ChatMessage{},
	}
	err = tulumCave.Create()
	if err != nil {
		return err
	}

	reefShipwreck := models.Activity{
		VendorId:       bson.ObjectIdHex("5933334a71d675aff38c53ba"),
		Name:           "Reef and Shipwreck Snorkeling Tour in Cancun",
		Description:    "Reef and Shipwreck Snorkeling Tour",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=32277P1&d=7171248&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://cache-graphicslib.viator.com/graphicslib/thumbs674x446/32277/SITours/reef-and-shipwreck-snorkeling-tour-in-cancun-in-cancun-304076.jpg",
		Price:          57,
		Location: models.MgoXY{
			X: 21.099722,
			Y: -86.765110,
		},
		IsEvening:        false,
		DurationHours:    2,
		Interests:        []string{"Extreme Sports", "Diving"},
		CommittedTripIds: []bson.ObjectId{},
		GroupChat:        []models.ChatMessage{},
	}
	err = reefShipwreck.Create()
	if err != nil {
		return err
	}

	comboTour := models.Activity{
		VendorId:       bson.ObjectIdHex("5933335271d675aff38c53bb"),
		Name:           "3-in-1 Discovery Combo Tour: Tulum Ruins, Reef Snorkeling Plus Cenote and Caves",
		Description:    "7-hour Riviera Maya tour",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=12861P5&d=153186&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://cache-graphicslib.viator.com/graphicslib/thumbs674x446/12861/SITours/3-in-1-discovery-combo-tour-tulum-ruins-reef-snorkeling-plus-cenote-in-cancun-446203.jpg",
		Price:          116,
		Location: models.MgoXY{
			X: 21.099722,
			Y: -86.765110,
		},
		IsEvening:        false,
		DurationHours:    7,
		Interests:        []string{"Diving", "Extreme Sports", "Caving", "Museum"},
		CommittedTripIds: []bson.ObjectId{},
		GroupChat:        []models.ChatMessage{},
	}
	err = comboTour.Create()
	if err != nil {
		return err
	}

	earlyAccess := models.Activity{
		VendorId:       bson.ObjectIdHex("5933335a71d675aff38c53bc"),
		Name:           "Viator Exclusive: Early Access to Tulum Ruins with an Archeologist",
		Description:    "Early-access Tulum tour with a private archaeologist — a Viator Exclusive",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=5885TULUM&d=150807&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://cache-graphicslib.viator.com/graphicslib/thumbs674x446/5885/SITours/viator-exclusive-early-access-to-tulum-ruins-with-an-archeologist-in-cancun-149589.jpg",
		Price:          28,
		Location: models.MgoXY{
			X: 21.099722,
			Y: -86.765110,
		},
		IsEvening:        false,
		DurationHours:    6,
		Interests:        []string{"Museum"},
		CommittedTripIds: []bson.ObjectId{},
		GroupChat:        []models.ChatMessage{},
	}
	err = earlyAccess.Create()
	if err != nil {
		return err
	}

	islaContoy := models.Activity{
		VendorId:       bson.ObjectIdHex("5933332c71d675aff38c53b7"),
		Name:           "Isla Contoy Day Trip: Snorkeling at Ixlache Reef and Eco-Paradise Tour",
		Description:    "Snorkeling at Ixlache Reef, followed by a guided tour and free time on Contoy Island, from Cancun",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=3206CONTOY&d=150807&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://cache-graphicslib.viator.com/graphicslib/thumbs674x446/3206/SITours/isla-contoy-day-trip-snorkeling-at-ixlache-reef-and-eco-paradise-tour-in-cancun-158394.jpg",
		Price:          105,
		Location: models.MgoXY{
			X: 21.099722,
			Y: -86.765110,
		},
		IsEvening:        false,
		DurationHours:    2,
		Interests:        []string{"Extreme Sports", "Diving"},
		CommittedTripIds: []bson.ObjectId{},
		GroupChat:        []models.ChatMessage{},
	}
	err = islaContoy.Create()
	if err != nil {
		return err
	}

	hookDinner := models.Activity{
		VendorId:       bson.ObjectIdHex("5933332c71d675aff38c53b7"),
		Name:           "Captain Hook Dinner Cruise in Cancun",
		Description:    "Captain Hook Pirate Dinner Cruise in Cancun!",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=3206CAPHOOK&d=150807&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://cache-graphicslib.viator.com/graphicslib/thumbs674x446/3206/SITours/captain-hook-dinner-cruise-in-cancun-in-cancun-182885.jpg",
		Price:          47,
		Location: models.MgoXY{
			X: 21.099722,
			Y: -86.765110,
		},
		IsEvening:        true,
		DurationHours:    3,
		Interests:        []string{"Food & Beverages", "Parties"},
		CommittedTripIds: []bson.ObjectId{},
		GroupChat:        []models.ChatMessage{},
	}
	err = hookDinner.Create()
	if err != nil {
		return err
	}

	chichenItza := models.Activity{
		VendorId:       bson.ObjectIdHex("5933332c71d675aff38c53b7"),
		Name:           "Chichen Itza Day Trip from Cancun",
		Description:    "Day trip to Chichen Itza from Cancun",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=3206CHI_01&d=150807&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://cache-graphicslib.viator.com/graphicslib/thumbs674x446/3206/SITours/chichen-itza-day-trip-from-cancun-in-cancun-39959.jpg",
		Price:          74,
		Location: models.MgoXY{
			X: 21.099722,
			Y: -86.765110,
		},
		IsEvening:        false,
		DurationHours:    4,
		Interests:        []string{"Bars", "Shopping"},
		CommittedTripIds: []bson.ObjectId{},
		GroupChat:        []models.ChatMessage{},
	}
	err = chichenItza.Create()
	if err != nil {
		return err
	}

	islasTour := models.Activity{
		VendorId:       bson.ObjectIdHex("5933332c71d675aff38c53b7"),
		Name:           "Isla Contoy and Isla Mujeres Tour with Snorkeling from Cancun or Playa del Carmen",
		Description:    "Full-day catamaran trip to Isla Contoy and Isla Mujeres",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=3206P63&d=150807&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://cache-graphicslib.viator.com/graphicslib/thumbs674x446/3206/SITours/isla-contoy-and-isla-mujeres-tour-with-snorkeling-from-cancun-or-in-cancun-396711.jpg",
		Price:          116,
		Location: models.MgoXY{
			X: 21.099722,
			Y: -86.765110,
		},
		IsEvening:        false,
		DurationHours:    6,
		Interests:        []string{"Extreme Sports", "Diving"},
		CommittedTripIds: []bson.ObjectId{},
		GroupChat:        []models.ChatMessage{},
	}
	err = islasTour.Create()
	if err != nil {
		return err
	}

	islaSailing := models.Activity{
		VendorId:       bson.ObjectIdHex("5933336371d675aff38c53bd"),
		Name:           "Isla Mujeres Sailing and Snorkel Cruise from Cancun",
		Description:    "7-hour sailing and snorkeling cruise to Isla Mujeres from Cancun",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=19720P1&d=150807&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://cache-graphicslib.viator.com/graphicslib/thumbs674x446/19720/SITours/isla-mujeres-sailing-and-snorkel-cruise-from-cancun-in-cancun-237428.jpg",
		Price:          79,
		Location: models.MgoXY{
			X: 21.099722,
			Y: -86.765110,
		},
		IsEvening:        false,
		DurationHours:    7,
		Interests:        []string{"Extreme Sports"},
		CommittedTripIds: []bson.ObjectId{},
		GroupChat:        []models.ChatMessage{},
	}
	err = islaSailing.Create()
	if err != nil {
		return err
	}

	romanticDinner := models.Activity{
		VendorId:       bson.ObjectIdHex("5933336e71d675aff38c53be"),
		Name:           "Cancun Romantic Dinner Cruise",
		Description:    "Embark on a unique, romantic Mississippi-style cruise",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=5248P8&d=150807&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://cache-graphicslib.viator.com/graphicslib/thumbs674x446/5248/SITours/cancun-romantic-dinner-cruise-in-cancun-457629.jpg",
		Price:          94,
		Location: models.MgoXY{
			X: 21.099722,
			Y: -86.765110,
		},
		IsEvening:        true,
		DurationHours:    3,
		Interests:        []string{"Food & Beverages", "Parties"},
		CommittedTripIds: []bson.ObjectId{},
		GroupChat:        []models.ChatMessage{},
	}
	err = romanticDinner.Create()
	if err != nil {
		return err
	}

	lobsterDinnerCruise := models.Activity{
		VendorId:       bson.ObjectIdHex("5933337671d675aff38c53bf"),
		Name:           "Lobster Dinner Cruise from Cancun",
		Description:    "Enjoy a relaxing dinner cruise on the Nichupte Lagoon from Cancun, Mexico",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=2895CANLOBSTER&d=150807&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://cache-graphicslib.viator.com/graphicslib/thumbs674x446/2895/SITours/lobster-dinner-cruise-from-cancun-in-cancun-115641.jpg",
		Price:          94,
		Location: models.MgoXY{
			X: 21.099722,
			Y: -86.765110,
		},
		IsEvening:        true,
		DurationHours:    3,
		Interests:        []string{"Food & Beverages", "Parties"},
		CommittedTripIds: []bson.ObjectId{},
		GroupChat:        []models.ChatMessage{},
	}
	err = lobsterDinnerCruise.Create()
	if err != nil {
		return err
	}

	jollyRodger := models.Activity{
		VendorId:       bson.ObjectIdHex("5933337e71d675aff38c53c0"),
		Name:           "Jolly Roger Pirate Night Show and Dinner",
		Description:    "3.5-hour pirate experience in Cancun",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=19334P1&d=150807&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://cache-graphicslib.viator.com/graphicslib/thumbs674x446/19334/SITours/jolly-roger-pirate-night-show-and-dinner-in-cancun-393034.jpg",
		Price:          101,
		Location: models.MgoXY{
			X: 21.099722,
			Y: -86.765110,
		},
		IsEvening:        true,
		DurationHours:    3.5,
		Interests:        []string{"Food & Beverages", "Parties"},
		CommittedTripIds: []bson.ObjectId{},
		GroupChat:        []models.ChatMessage{},
	}
	err = jollyRodger.Create()
	if err != nil {
		return err
	}

	return nil
}
