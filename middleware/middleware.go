package middleware

import (
	"github.com/tanyoAH/tanyo-server/config"
	ctx "github.com/tanyoAH/tanyo-server/context"
	"github.com/tanyoAH/tanyo-server/models"
	"github.com/tanyoAH/tanyo-server/utils"
	"net/http"
)

var Log = config.Conf.GetLogger()

func RecoverAndLog(handler http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				utils.JSONInternalError(w, "An internal server error occurred", "Please try again in a few seconds")
				Log.Error("Panic occurred in HTTP handler:", r)
			}
		}()
		handler.ServeHTTP(w, r)
	})
}

func GetContext(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var passItOn = func() {
			handler.ServeHTTP(w, r)
		}

		if at := r.Header.Get("X-ACCESS-TOKEN"); at != "" {
			if at != "" {
				user := models.User{AccessToken: at}
				err := user.FindByAccessToken()
				if err != nil {
					passItOn()
				} else {
					r = r.WithContext(ctx.AddCurrentUserToContext(r, user))
					passItOn()
				}
			} else {
				passItOn()
			}
		} else {
			passItOn()
		}
	}
}

func RequireUserForAPI(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var respondUnauthorized = func() {
			utils.JSONDetailed(w, utils.APIResponse{Message: "Unauthorized", Debug: "Invalid or missing access token header/cookie"}, http.StatusUnauthorized)
		}
		user, err := ctx.GetCurrentUser(r)
		if err != nil || user.Id == "" {
			respondUnauthorized()
			return
		} else {
			handler.ServeHTTP(w, r)
		}
	}
}

// `Use` allows us to stack middleware to process the request
// Example taken from https://github.com/gorilla/mux/pull/36#issuecomment-25849172
func Use(handler http.HandlerFunc, mid ...func(http.Handler) http.HandlerFunc) http.HandlerFunc {
	for _, m := range mid {
		handler = m(handler)
	}
	return handler
}
