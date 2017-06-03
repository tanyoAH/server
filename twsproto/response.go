package twsproto

type Response struct {
	ReqId  string      `json:"reqId"`
	Status int64       `json:"status"`
	Debug  string      `json:"debug,omitempty"`
	Error  string      `json:"error,omitempty"`
	Data   interface{} `json:"data"`
}
