package controllers

import (
	"github.com/gorilla/mux"
	"github.com/tanyoAH/tanyo-server/config"
	"github.com/tanyoAH/tanyo-server/controllers/api"
	"net/http"
)

var Log = config.Conf.GetLogger()

func CreateRouter() http.Handler {
	router := mux.NewRouter()

	apiV0Router := router.PathPrefix("/api/v0").Subrouter()
	apiV0Router.HandleFunc("/", Use(api.V0_API, RequireUserForAPI)).Methods("GET")
	apiV0Router.HandleFunc("/me/trips", Use(api.V0_GetMyTrips, RequireUserForAPI)).Methods("GET")

	apiV0Router.HandleFunc("/trips", Use(api.V0_CreateTrip, RequireUserForAPI)).Methods("POST")
	apiV0Router.HandleFunc("/trips/{tripId}", Use(api.V0_GetTrip, RequireUserForAPI)).Methods("GET")
	apiV0Router.HandleFunc("/trips/{tripId}/recommendations", Use(api.V0_GetActivityRecommendationsForTrip, RequireUserForAPI)).Methods("GET")

	apiV0Router.HandleFunc("/trips/{tripId}/activities/{activityId}", Use(api.V0_GetActivity, RequireUserForAPI)).Methods("GET")
	apiV0Router.HandleFunc("/trips/{tripId}/activities/{activityId}/commitments", Use(api.V0_CommitToActivity, RequireUserForAPI)).Methods("GET")

	marketingSiteFS := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	router.PathPrefix("/static/").Handler(marketingSiteFS).Methods("GET")

	return Use(router.ServeHTTP, GetContext, RecoverAndLog)
}

// `Use` allows us to stack middleware to process the request
// Example taken from https://github.com/gorilla/mux/pull/36#issuecomment-25849172
func Use(handler http.HandlerFunc, mid ...func(http.Handler) http.HandlerFunc) http.HandlerFunc {
	for _, m := range mid {
		handler = m(handler)
	}
	return handler
}
