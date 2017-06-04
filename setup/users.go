package setup

import "github.com/tanyoAH/tanyo-server/models"

func setupUsers() error {
	kimA := models.User{
		ProfileUrl:         "https://images.pexels.com/photos/4618/man-person-hat-fur.jpg?h=350&auto=compress&cs=tinysrgb",
		AccessToken:        "at_kim_a",
		FullName:           "Kim A.",
		Headline:           "Young Entrepreneur",
		Age:                23,
		TravelledCountries: []string{"Vietnam", "Singapore", "Japan"},
		From:               "Philippines",
		Interests:          []string{"Painting", "Extreme Sports", "Surfing"},
	}
	err := kimA.Create()
	if err != nil {
		return err
	}

	victorB := models.User{
		ProfileUrl:         "https://images.pexels.com/photos/127229/pexels-photo-127229.jpeg?h=350&auto=compress&cs=tinysrgb",
		AccessToken:        "at_victor_b",
		FullName:           "Victor B.",
		Headline:           "Movie Director",
		Age:                29,
		TravelledCountries: []string{"China", "France", "Italy"},
		From:               "Hong Kong",
		Interests:          []string{"Dance", "Theatre", "Kayaking"},
	}
	err = victorB.Create()
	if err != nil {
		return err
	}

	lucyC := models.User{
		ProfileUrl:         "https://images.pexels.com/photos/216983/pexels-photo-216983.jpeg?h=350&auto=compress&cs=tinysrgb",
		AccessToken:        "at_lucy_c",
		FullName:           "Lucy C.",
		Headline:           "Fashion Model",
		Age:                25,
		TravelledCountries: []string{"Korea", "Germany", "Spain"},
		From:               "United States",
		Interests:          []string{"Food & Beverage", "Shopping", "Surfing"},
	}
	err = lucyC.Create()
	if err != nil {
		return err
	}

	florenceD := models.User{
		ProfileUrl:         "https://images.pexels.com/photos/274595/pexels-photo-274595.jpeg?h=350&auto=compress&cs=tinysrgb",
		AccessToken:        "at_florence_d",
		FullName:           "Florence D.",
		Headline:           "Doctor",
		Age:                28,
		TravelledCountries: []string{"Poland", "Bolivia", "Brazil"},
		From:               "Singapore",
		Interests:          []string{"Bars", "Parties", "Kayaking"},
	}
	err = florenceD.Create()
	if err != nil {
		return err
	}

	kellyE := models.User{
		ProfileUrl:         "https://images.pexels.com/photos/218721/pexels-photo-218721.jpeg?h=350&auto=compress&cs=tinysrgb",
		AccessToken:        "at_kelly_e",
		FullName:           "Kelly E.",
		Headline:           "World Traveler",
		Age:                32,
		TravelledCountries: []string{"Thailand", "Indonesia", "Chile"},
		From:               "Canada",
		Interests:          []string{"Biking", "Surfing", "Kayaking"},
	}
	err = kellyE.Create()
	if err != nil {
		return err
	}

	jerryF := models.User{
		ProfileUrl:         "https://images.pexels.com/photos/301958/pexels-photo-301958.jpeg?h=350&auto=compress&cs=tinysrgb",
		AccessToken:        "at_jerry_f",
		FullName:           "Jerry F.",
		Headline:           "Ballet Dancer",
		Age:                35,
		TravelledCountries: []string{"Canada", "United States", "Mexico"},
		From:               "Japan",
		Interests:          []string{"Horse Riding", "Surfing", "Skiing"},
	}
	err = jerryF.Create()
	if err != nil {
		return err
	}

	paulG := models.User{
		ProfileUrl:         "https://images.pexels.com/photos/65109/pexels-photo-65109.jpeg?h=350&auto=compress&cs=tinysrgb",
		AccessToken:        "at_paul_g",
		FullName:           "Paul G.",
		Headline:           "Weird Scientist",
		Age:                36,
		TravelledCountries: []string{"Philippines", "Malaysia", "Sweden"},
		From:               "Russia",
		Interests:          []string{"Biking", "Skiing", "Kayaking"},
	}
	err = paulG.Create()
	if err != nil {
		return err
	}

	stephenH := models.User{
		ProfileUrl:         "https://images.pexels.com/photos/403936/pexels-photo-403936.jpeg?h=350&auto=compress&cs=tinysrgb",
		AccessToken:        "at_stephen_h",
		FullName:           "Stephen H.",
		Headline:           "Handsome Dude",
		Age:                38,
		TravelledCountries: []string{"Norway", "Bulgaria", "Peru"},
		From:               "Poland",
		Interests:          []string{"Drawing", "Surfing", "Theater"},
	}
	err = stephenH.Create()
	if err != nil {
		return err
	}

	daisyI := models.User{
		ProfileUrl:         "https://images.pexels.com/photos/253758/pexels-photo-253758.jpeg?h=350&auto=compress&cs=tinysrgb",
		AccessToken:        "at_daisy_i",
		FullName:           "Daisy I.",
		Headline:           "Opera Singer",
		Age:                20,
		TravelledCountries: []string{"Russia", "Hong Kong", "India"},
		From:               "Korea",
		Interests:          []string{"Kayaking", "Surfing", "Theater"},
	}
	err = daisyI.Create()
	if err != nil {
		return err
	}

	robertJ := models.User{
		ProfileUrl:         "https://images.pexels.com/photos/4618/man-person-hat-fur.jpg?h=350&auto=compress&cs=tinysrgb",
		AccessToken:        "at_robert_j",
		FullName:           "Robert J.",
		Headline:           "Artist",
		Age:                22,
		TravelledCountries: []string{"South Africa", "Argentina", "Venezuela"},
		From:               "Spain",
		Interests:          []string{"Kayaking", "Shopping", "Museum"},
	}
	err = robertJ.Create()
	if err != nil {
		return err
	}

	seanK := models.User{
		ProfileUrl:         "https://images.pexels.com/photos/269334/pexels-photo-269334.jpeg?h=350&auto=compress&cs=tinysrgb",
		AccessToken:        "at_sean_k",
		FullName:           "Sean K.",
		Headline:           "Jazz Guitarist",
		Age:                27,
		TravelledCountries: []string{"United Kingdom", "Turkey", "Finland"},
		From:               "France",
		Interests:          []string{"Theater", "Surfing", "Food & Beverage"},
	}
	err = seanK.Create()
	if err != nil {
		return err
	}

	katieL := models.User{
		ProfileUrl:         "https://images.pexels.com/photos/37533/studio-portrait-woman-face-37533.jpeg?h=350&auto=compress&cs=tinysrgb",
		AccessToken:        "at_katie_l",
		FullName:           "Katie L.",
		Headline:           "Professor",
		Age:                30,
		TravelledCountries: []string{"Nepal", "Egypt", "Saudi Arabia"},
		From:               "Mexico",
		Interests:          []string{"Sculpting", "Drawing", "Kayaking"},
	}
	err = katieL.Create()
	if err != nil {
		return err
	}

	return nil
}
