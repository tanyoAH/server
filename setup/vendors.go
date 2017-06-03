package setup

import "github.com/tanyoAH/tanyo-server/models"

func setupVendors() error {
	dolphinDiscovery := models.Vendor{
		Id:             "593332f371d675aff38c53b6",
		Name:           "Dolphin Discovery",
		Description:    "",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=3911P70&d=5432465&aidSuffix=tvrm&partner=Viator",
		ThumbnailUrl:   "https://cache-graphicslib.viator.com/graphicslib/media/02/sunset-photo_21102082-622x408.jpg",
	}
	err := dolphinDiscovery.Create()
	if err != nil {
		return err
	}

	cancunVacationExperts := models.Vendor{
		Id:             "5933332c71d675aff38c53b7",
		Name:           "Cancun Vacation Experts",
		Description:    "",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=3206CUNATVZIP&d=150807&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://cache-graphicslib.viator.com/graphicslib/media/62/zip-line-fun-photo_16698978-622x408.jpg",
	}
	err = cancunVacationExperts.Create()
	if err != nil {
		return err
	}

	ecoColors := models.Vendor{
		Id:             "5933333771d675aff38c53b8",
		Name:           "EcoColors",
		Description:    "",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=3467WHALE&d=150807&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://media-cdn.tripadvisor.com/media/photo-f/0f/70/fb/c6/come-and-swim-with-the.jpg",
	}
	err = ecoColors.Create()
	if err != nil {
		return err
	}

	cancunAdventure := models.Vendor{
		Id:             "5933334171d675aff38c53b9",
		Name:           "Cancun Adventure",
		Description:    "",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=7460P2&d=150807&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://media-cdn.tripadvisor.com/media/photo-s/0f/73/a1/da/cancun-adventures-sunset.jpg",
	}
	err = cancunAdventure.Create()
	if err != nil {
		return err
	}

	totalSnorkel := models.Vendor{
		Id:             "5933334a71d675aff38c53ba",
		Name:           "Total Snorkel",
		Description:    "",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=32277P1&d=7171248&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://media-cdn.tripadvisor.com/media/photo-f/0f/6a/e9/c9/photo1jpg.jpg",
	}
	err = totalSnorkel.Create()
	if err != nil {
		return err
	}

	oceanTours := models.Vendor{
		Id:             "5933335271d675aff38c53bb",
		Name:           "Ocean Tours",
		Description:    "",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=12861P5&d=153186&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://media-cdn.tripadvisor.com/media/photo-s/0f/6c/1f/8d/5-38-largejpg.jpg",
	}
	err = oceanTours.Create()
	if err != nil {
		return err
	}

	amigoTours := models.Vendor{
		Id:             "5933335a71d675aff38c53bc",
		Name:           "Amigo Tours",
		Description:    "",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=5885TULUM&d=150807&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://media-cdn.tripadvisor.com/media/photo-f/0f/66/f9/1e/photo3jpg.jpg",
	}
	err = amigoTours.Create()
	if err != nil {
		return err
	}

	sambaCatamarans := models.Vendor{
		Id:             "5933336371d675aff38c53bd",
		Name:           "Samba Catamarans",
		Description:    "",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=19720P1&d=150807&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://media-cdn.tripadvisor.com/media/photo-f/0f/6b/86/20/catamaran-frontal-view.jpg",
	}
	err = sambaCatamarans.Create()
	if err != nil {
		return err
	}

	aquaWorld := models.Vendor{
		Id:             "5933336e71d675aff38c53be",
		Name:           "AquaWorld",
		Description:    "",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=5248P8&d=150807&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://media-cdn.tripadvisor.com/media/photo-s/0f/71/0f/5d/photo2jpg.jpg",
	}
	err = aquaWorld.Create()
	if err != nil {
		return err
	}

	olympusTours := models.Vendor{
		Id:             "5933337671d675aff38c53bf",
		Name:           "Olympus Tours",
		Description:    "",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=2895CANLOBSTER&d=150807&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://cache-graphicslib.viator.com/graphicslib/media/14/leaving-the-dock-photo_18685972-622x408.jpg",
	}
	err = olympusTours.Create()
	if err != nil {
		return err
	}

	pirateShowCancun := models.Vendor{
		Id:             "5933337e71d675aff38c53c0",
		Name:           "Pirate Show Cancun Jolly Roger",
		Description:    "",
		TripAdvisorUrl: "https://en.tripadvisor.com.hk/AttractionProductDetail?product=19334P1&d=150807&aidSuffix=xsell&partner=Viator",
		ThumbnailUrl:   "https://cache-graphicslib.viator.com/graphicslib/media/5f/pirates-in-uproar-photo_13602655-622x408.jpg",
	}
	err = pirateShowCancun.Create()
	if err != nil {
		return err
	}

	return nil
}
