package tws

import "github.com/tanyoAH/tanyo-server/twsproto"

type TWSHandler func(*Session, *twsproto.Request)

type TWSNotificationHandler func(*Session, *twsproto.Notification)

type TWSVNotificationHandlers map[string]TWSNotificationHandler

type TWSHandlers map[string]TWSHandler

var handlers TWSHandlers = map[string]TWSHandler{
// TODO
}

var notificationHandlers TWSVNotificationHandlers = map[string]TWSNotificationHandler{
// TODO handle inbound notifications
}
