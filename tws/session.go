package tws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/tanyoAH/tanyo-server/twsproto"
	"net/http"
	"time"
)

type Session struct {
	IsClosed     bool
	connection   *websocket.Conn
	WriteChannel chan twsproto.MessageWrapper
	WSENV        *State
}

func (sess *Session) Initialize(conn *websocket.Conn, s *State) {
	sess.connection = conn
	sess.WSENV = s
	sess.WriteChannel = make(chan twsproto.MessageWrapper)
}

func (sess *Session) Close() error {
	if sess.IsClosed {
		return nil
	}
	sess.IsClosed = true
	// TODO any other cleanup?
	return sess.connection.Close()
}

func (sess *Session) Reader() {
	defer sess.Close()
	sess.connection.SetReadLimit(maxMessageSize)
	sess.connection.SetReadDeadline(time.Now().Add(pongWait))
	sess.connection.SetPongHandler(func(string) error { sess.connection.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		if sess.IsClosed || sess.connection == nil {
			break
		}
		msgType, msg, err := sess.connection.ReadMessage()
		if err != nil || msgType == websocket.CloseMessage {
			break
		}
		if msgType == websocket.PingMessage || msgType == websocket.PongMessage {
			continue
		}
		if msgType != websocket.TextMessage {
			sess.WriteChannel <- twsproto.MessageWrapper{Type: twsproto.ErrorMessageType, Payload: twsproto.ErrorPayloadSansRequest{Error: ErrInvalidMessageType.Error()}}
			continue
		}
		msgWrapper := twsproto.MessageWrapper{}
		err = json.Unmarshal(msg, &msgWrapper)
		if err != nil {
			sess.WriteChannel <- twsproto.MessageWrapper{Type: twsproto.ErrorMessageType, Payload: twsproto.ErrorPayloadSansRequest{Error: ErrInvalidJSON.Error()}}
			continue
		}
		switch msgWrapper.Type {
		// TODO: Accept potential response messages in the future -- not setup to handle those yet though
		case twsproto.RequestMessageType:
			reqWrapper := twsproto.RequestMessageWrapper{}
			err = json.Unmarshal(msg, &reqWrapper)
			if err != nil {
				sess.WriteChannel <- twsproto.MessageWrapper{Type: twsproto.ErrorMessageType, Payload: twsproto.ErrorPayloadSansRequest{Error: ErrInvalidJSON.Error()}}
				continue
			}
			if !twsproto.IsValidId(reqWrapper.Payload.Id) {
				sess.WriteChannel <- twsproto.MessageWrapper{Type: twsproto.ErrorMessageType, Payload: twsproto.ErrorPayloadSansRequest{Error: ErrInvalidId.Error()}}
				continue
			}
			handlerFunc, found := handlers[reqWrapper.Payload.Method]
			if !found {
				sess.WriteChannel <- twsproto.MessageWrapper{Type: twsproto.ResponseMessageType, Payload: twsproto.ErrorPayload{RequestId: reqWrapper.Payload.Id, Error: ErrMethodNotFound.Error()}}
				continue
			}
			go handlerFunc(sess, &reqWrapper.Payload)
		case twsproto.NotificationMessageType:
			notiWrapper := twsproto.NotificationMessageWrapper{}
			err = json.Unmarshal(msg, &notiWrapper)
			if err != nil {
				sess.WriteChannel <- twsproto.MessageWrapper{Type: twsproto.ErrorMessageType, Payload: twsproto.ErrorPayloadSansRequest{Error: ErrInvalidJSON.Error()}}
				continue
			}
			handlerFunc, found := notificationHandlers[notiWrapper.Payload.Event]
			if !found {
				sess.WriteChannel <- twsproto.MessageWrapper{Type: twsproto.ErrorMessageType, Payload: twsproto.ErrorPayloadSansRequest{Error: ErrInvalidJSON.Error()}}
				continue
			}
			go handlerFunc(sess, &notiWrapper.Payload)
		default:
			sess.WriteChannel <- twsproto.MessageWrapper{Type: twsproto.ErrorMessageType, Payload: twsproto.ErrorPayloadSansRequest{Error: ErrInvalidMessageType.Error()}}
			continue
		}
	}
}

func (sess *Session) Writer() {
	pingTicker := time.NewTicker(pingPeriod)
	defer func() {
		pingTicker.Stop()
		sess.Close()
	}()
	for {
		select {
		case msg := <-sess.WriteChannel:
			if sess.IsClosed || sess.connection == nil {
				return
			}
			if msg.Type == twsproto.ResponseMessageType {
				if resp, ok := msg.Payload.(twsproto.Response); ok {
					if resp.Status == 0 {
						resp.Status = http.StatusOK
						msg = twsproto.MessageWrapper{Type: msg.Type, Payload: resp}
					}
				}
			}
			bytes, err := json.Marshal(&msg)
			if err != nil {
				Log.WithField("error", err).Error("Error marshalling json response")
				return
			}
			sess.connection.SetWriteDeadline(time.Now().Add(writeWait))
			if err := sess.connection.WriteMessage(websocket.TextMessage, bytes); err != nil {
				return
			}
		case <-pingTicker.C:
			if sess.IsClosed || sess.connection == nil {
				return
			}
			sess.connection.SetWriteDeadline(time.Now().Add(writeWait))
			if err := sess.connection.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
