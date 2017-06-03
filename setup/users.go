package setup

import "github.com/tanyoAH/tanyo-server/models"

func setupUsers() error {
	kimA := models.User{
		AccessToken:        "at_kim_a",
		FullName:           "Kim A.",
		Headline:           "Young Entrepreneur",
		Age:                23,
		TravelledCountries: []string{"Vietnam", "Singapore", "Japan"},
		From:               "Philippines",
		Interests:          []string{"Painting", "Sculpting", "Surfing"},
	}
	err := kimA.Create()
	if err != nil {
		return err
	}

	victorB := models.User{
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

	// TODO the rest

	return nil
}
