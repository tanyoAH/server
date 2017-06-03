package twsproto

type ErrorPayload struct {
	RequestId string `json:"reqId"`
	Error     string `json:"error"`
}
