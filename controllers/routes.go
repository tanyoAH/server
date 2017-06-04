package controllers

import (
	"github.com/gorilla/mux"
	"github.com/tanyoAH/tanyo-server/config"
	"github.com/tanyoAH/tanyo-server/controllers/api"
	"github.com/tanyoAH/tanyo-server/middleware"
	"github.com/tanyoAH/tanyo-server/tws"
	"net/http"
)

var Log = config.Conf.GetLogger()

func CreateRouter(ws *tws.State) http.Handler {
	router := mux.NewRouter()

	apiV0Router := router.PathPrefix("/api/v0").Subrouter()
	apiV0Router.HandleFunc("/", middleware.Use(api.V0_API, middleware.RequireUserForAPI)).Methods("GET")
	apiV0Router.HandleFunc("/me", middleware.Use(api.V0_GetMyProfile, middleware.RequireUserForAPI)).Methods("GET")
	apiV0Router.HandleFunc("/me/trips", middleware.Use(api.V0_GetMyTrips, middleware.RequireUserForAPI)).Methods("GET")

	apiV0Router.HandleFunc("/users/{userId}", middleware.Use(api.V0_GetUserProfile, middleware.RequireUserForAPI)).Methods("GET")

	apiV0Router.HandleFunc("/trips", middleware.Use(api.V0_CreateTrip, middleware.RequireUserForAPI)).Methods("POST")
	apiV0Router.HandleFunc("/trips/{tripId}", middleware.Use(api.V0_GetTrip, middleware.RequireUserForAPI)).Methods("GET")
	apiV0Router.HandleFunc("/trips/{tripId}/recommendations", middleware.Use(api.V0_GetActivityRecommendationsForTrip, middleware.RequireUserForAPI)).Methods("GET")

	apiV0Router.HandleFunc("/trips/{tripId}/activities/{activityId}", middleware.Use(api.V0_GetActivity, middleware.RequireUserForAPI)).Methods("GET")
	apiV0Router.HandleFunc("/trips/{tripId}/activities/{activityId}/commitments", middleware.Use(api.V0_CommitToActivity, middleware.RequireUserForAPI)).Methods("POST")

	marketingSiteFS := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	router.PathPrefix("/static/").Handler(marketingSiteFS).Methods("GET")

	return middleware.Use(router.ServeHTTP, middleware.GetContext, middleware.RecoverAndLog)
}
