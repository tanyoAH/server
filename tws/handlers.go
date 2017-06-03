package tws

import "github.com/tanyoAH/tanyo-server/twsproto"

type WSENVHandler func(*Session, *twsproto.Request)

type WSENVNotificationHandler func(*Session, *twsproto.Notification)

type WSENVNotificationHandlers map[string]WSENVNotificationHandler

type WSENVHandlers map[string]WSENVHandler

var handlers WSENVHandlers = map[string]WSENVHandler{
// TODO
}

var notificationHandlers WSENVNotificationHandlers = map[string]WSENVNotificationHandler{
// TODO handle inbound notifications
}
