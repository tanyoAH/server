package twsproto

import "net/http"

func SendNotificationWithError(w chan MessageWrapper, event, errMsg string, statusCode int64) {
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}
	if errMsg == "" {
		errMsg = "An error occurred"
	}
	w <- MessageWrapper{Type: NotificationMessageType, Payload: Notification{Event: event, Error: errMsg, Status: statusCode}}
}

func SendNotification(w chan MessageWrapper, event string, data interface{}) {
	w <- MessageWrapper{Type: NotificationMessageType, Payload: Notification{Event: event, Data: data, Status: http.StatusOK}}
}
