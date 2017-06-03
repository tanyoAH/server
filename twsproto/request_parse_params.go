package twsproto

import (
	"encoding/json"
	"net/http"
	"reflect"
)

const (
	WSNotificationErrorEvent = "error"
)

func JSONDecodeAndNotify(w chan MessageWrapper, r *Notification, outStruct interface{}) error {
	bytes, err := json.Marshal(&r.Data)
	if err != nil {
		SendNotificationWithError(w, WSNotificationErrorEvent, "Invalid JSON", http.StatusBadRequest)
		return err
	}
	err = json.Unmarshal(bytes, &outStruct)
	if err != nil {
		SendNotificationWithError(w, WSNotificationErrorEvent, "Invalid JSON", http.StatusBadRequest)
		return err
	}
	if !isCheckableRequest(outStruct) {
		return nil
	}
	method := reflect.ValueOf(outStruct).MethodByName("Parameters").Interface().(func() error)
	err = method()
	if err != nil {
		SendNotificationWithError(w, WSNotificationErrorEvent, "Invalid JSON", http.StatusBadRequest)
		return err
	}
	return nil
}

func JSONDecodeAndCatch(w chan MessageWrapper, r *Request, outStruct interface{}) error {
	bytes, err := json.Marshal(&r.Data)
	if err != nil {
		RespondError(w, r, "Invalid JSON", http.StatusBadRequest)
		return err
	}
	err = json.Unmarshal(bytes, &outStruct)
	if err != nil {
		RespondError(w, r, "Invalid JSON", http.StatusBadRequest)
		return err
	}
	if !isCheckableRequest(outStruct) {
		return nil
	}
	method := reflect.ValueOf(outStruct).MethodByName("Parameters").Interface().(func() error)
	err = method()
	if err != nil {
		RespondError(w, r, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

type CheckableRequest interface {
	Parameters() error
}

func isCheckableRequest(checkAgainst interface{}) bool {
	// TODO: a nil reader or checkAgainst should throw an error. This is just a temp fix to stop the panic
	reader := reflect.TypeOf((*CheckableRequest)(nil)).Elem()
	if reader == nil || checkAgainst == nil {
		return false
	}
	return reflect.TypeOf(checkAgainst).Implements(reader)
}
