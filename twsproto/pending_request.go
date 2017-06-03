package twsproto

type PendingRequest struct {
	Request  *Request
	RespChan chan *Response
}
