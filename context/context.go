package context

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tanyoAH/tanyo-server/models"
	"github.com/tanyoAH/tanyo-server/utils"
	"net/http"
)

const currentUserContextKey string = "user"

type APIResponse struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Debug   string      `json:"debug,omitempty"`
}

var unauthorizedError APIResponse = APIResponse{Message: "Unauthorized", Success: false, Debug: "Token/account error"}

func AddCurrentUserToContext(r *http.Request, user models.User) context.Context {
	return context.WithValue(r.Context(), currentUserContextKey, user)
}

func GetCurrentUser(r *http.Request) (models.User, error) {
	user := r.Context().Value(currentUserContextKey)
	switch user.(type) {
	case nil:
		return models.User{}, errors.New("Current user not found")
	default:
		return user.(models.User), nil
	}
}

func GetCurrentUserAndCatchForAPI(w http.ResponseWriter, r *http.Request) (models.User, error) {
	var user models.User
	user, err := GetCurrentUser(r)
	if err != nil || user.Id == "" {
		jsonWriter(w, unauthorizedError, http.StatusUnauthorized)
		return user, err
	}
	return user, nil
}

func GetCurrentUserAndCatchForView(w http.ResponseWriter, r *http.Request) (models.User, error) {
	var user models.User
	user, err := GetCurrentUser(r)
	if err != nil || user.Id == "" {
		utils.UnauthorizedErrorView(w, r)
		return user, err
	}
	return user, nil
}

func jsonWriter(w http.ResponseWriter, d interface{}, c int) {
	//dj, err := json.MarshalIndent(d, "", "  ")
	dj, err := json.Marshal(d)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", dj)
}
