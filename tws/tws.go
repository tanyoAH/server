package tws

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/tanyoAH/tanyo-server/config"
	"github.com/tanyoAH/tanyo-server/controllers"
)

var Log = config.Conf.GetLogger()

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	// TODO Set this to a reasonable size
	maxMessageSize = 1000000
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func init() {
	if config.Conf.IsDebugMode() {
		upgrader.CheckOrigin = func(r *http.Request) bool {
			return true
		}
	} else {
		upgrader.CheckOrigin = func(r *http.Request) bool {
			if r.Header.Get("Origin") == config.Conf.Origin {
				return true
			}
			return false
		}
	}
}

type State struct {
	// TODO
}

func CreateAndSetupNewState() (*State, error) {
	state := State{}
	// TODO setup state
	return &state, nil
}

func (se *State) CreateRouter() http.Handler {
	router := mux.NewRouter()
	router = router.StrictSlash(true)

	router.HandleFunc("/v0/ws", serveWSENV_V0(se)).Methods("GET")

	return controllers.Use(router.ServeHTTP, controllers.GetContext, controllers.RecoverAndLog)
}

func serveWSENV_V0(se *State) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		sess := Session{}
		sess.Initialize(conn, se)
		go sess.Writer()
		sess.Reader()
	}
}
