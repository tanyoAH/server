package twsproto

type Request struct {
	Id     string      `json:"id"`
	Method string      `json:"method"`
	Data   interface{} `json:"data"`
}
