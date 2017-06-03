package api

import (
	"github.com/tanyoAH/tanyo-server/config"
	"github.com/tanyoAH/tanyo-server/context"
	"github.com/tanyoAH/tanyo-server/utils"
	"net/http"
)

var Log = config.Conf.GetLogger()

func V0_API(w http.ResponseWriter, r *http.Request) {
	_, err := context.GetCurrentUserAndCatchForAPI(w, r)
	if err != nil {
		return
	}

	utils.JSONSuccess(w, nil, "Server healthy")
}
