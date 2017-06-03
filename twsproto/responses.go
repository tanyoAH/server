package twsproto

import "net/http"

func RespondError(w chan MessageWrapper, r *Request, errMsg string, statusCode int64) {
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}
	if errMsg == "" {
		errMsg = "An error occurred"
	}
	w <- MessageWrapper{Type: ResponseMessageType, Payload: Response{ReqId: r.Id, Error: errMsg, Status: statusCode}}
}

func RespondSuccess(w chan MessageWrapper, r *Request, data interface{}) {
	w <- MessageWrapper{Type: ResponseMessageType, Payload: Response{ReqId: r.Id, Data: data, Status: http.StatusOK}}
}
