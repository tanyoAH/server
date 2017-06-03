package startup

import (
	"net/http"

	"github.com/tanyoAH/tanyo-server/config"
	"github.com/tanyoAH/tanyo-server/controllers"
	// Import xwsenv for it's init routine that's relied upon for startup
	_ "github.com/NXC-Solutions/xwsenv"
	"github.com/gorilla/handlers"
	"github.com/tanyoAH/tanyo-server/models" /* MySQL */
	"os"

	"github.com/tanyoAH/tanyo-server/tws"
)

var Log = config.Conf.GetLogger()

var CorsHandler = handlers.CORS(handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}), handlers.AllowCredentials(), handlers.AllowedHeaders([]string{"content-type"}), handlers.AllowedOrigins([]string{config.Conf.Home, config.Conf.Origin}))

func StartServer(withTWS bool) {
	Log.Info("Server Loaded")

	err := models.Setup()
	if err != nil {
		Log.WithField("error", err).Fatal("Couldn't setup models")
	}

	if withTWS {
		Log.Info("Initializing WSENV environment(s)")
		ws, err := tws.CreateAndSetupNewState()
		if err != nil {
			Log.WithField("error", err).Fatal("Couldn't setup WSENV Eevironment(s)")
		}
		go startTWSServer(ws)
	}

	Log.WithField("hostname", config.Conf.ApiUrl).Info("API Server starting")
	if config.Conf.IsDebugMode() {
		http.ListenAndServe(config.Conf.ApiUrl, CorsHandler(handlers.CombinedLoggingHandler(os.Stdout, controllers.CreateRouter())))
	} else {
		http.ListenAndServe(config.Conf.ApiUrl, handlers.CombinedLoggingHandler(os.Stdout, controllers.CreateRouter()))
	}
}

func startTWSServer(w *tws.State) {
	Log.WithField("hostname", config.Conf.WsUrl).Info("TWS Server starting")
	if config.Conf.IsDebugMode() {
		http.ListenAndServe(config.Conf.WsUrl, CorsHandler(handlers.CombinedLoggingHandler(os.Stdout, w.CreateRouter())))
	} else {
		http.ListenAndServe(config.Conf.WsUrl, handlers.CombinedLoggingHandler(os.Stdout, w.CreateRouter()))
	}
}
