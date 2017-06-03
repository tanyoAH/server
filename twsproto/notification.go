package twsproto

type Notification struct {
	Event  string      `json:"event"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error,omitempty"`
	Status int64       `json:"status"`
}
