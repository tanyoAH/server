package twsproto

type MessageWrapper struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

func GenerateMessageWrapperForResponse(reqId string) (MessageWrapper, Response) {
	return MessageWrapper{Type: ResponseMessageType}, Response{ReqId: reqId}
}

func GenerateMessageWrapperForRequest() (MessageWrapper, Request) {
	return MessageWrapper{Type: RequestMessageType}, Request{Id: GenerateRandomId()}
}

type RequestMessageWrapper struct {
	Type    string  `json:"type"`
	Payload Request `json:"payload"`
}

type ResponseMessageWrapper struct {
	Type    string   `json:"type"`
	Payload Response `json:"payload"`
}

type NotificationMessageWrapper struct {
	Type    string       `json:"type"`
	Payload Notification `json:"payload"`
}
