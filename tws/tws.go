package tws

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/tanyoAH/tanyo-server/config"
	"github.com/tanyoAH/tanyo-server/context"
	"github.com/tanyoAH/tanyo-server/middleware"
	"github.com/tanyoAH/tanyo-server/twsproto"
	"github.com/tanyoAH/tanyo-server/utils"
	"sync"
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
	activeUsersMutex *sync.Mutex
	activeUsers      map[string]*Session
}

func CreateAndSetupNewState() (*State, error) {
	state := State{
		activeUsersMutex: &sync.Mutex{},
		activeUsers:      map[string]*Session{},
	}
	return &state, nil
}

func (state *State) CreateRouter() http.Handler {
	router := mux.NewRouter()
	router = router.StrictSlash(true)

	router.HandleFunc("/v0/ws", serveWS_V0(state)).Methods("GET")

	return middleware.Use(router.ServeHTTP, middleware.GetContext, middleware.RecoverAndLog)
}

func (state *State) AddSession(userId string, session *Session) {
	state.activeUsers[userId] = session
}

func (state *State) DropSession(userId string) {
	state.activeUsersMutex.Lock()
	defer state.activeUsersMutex.Unlock()
	delete(state.activeUsers, userId)
}

func (state *State) NotifyUsers(ids []string, event string, data interface{}) {
	for _, id := range ids {
		state.activeUsersMutex.Lock()
		if s, ok := state.activeUsers[id]; ok {
			twsproto.SendNotification(s.WriteChannel, event, data)
		}
		state.activeUsersMutex.Unlock()
	}
}

func serveWS_V0(se *State) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := context.GetCurrentUser(r)
		if err != nil {
			utils.JSONForbiddenError(w, "Invalid user", "")
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		sess := Session{
			User: user,
		}
		sess.Initialize(conn, se)
		go sess.Writer()
		se.AddSession(user.Id.String(), &sess)
		sess.Reader()
		se.DropSession(user.Id.String())
	}
}
