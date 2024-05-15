package master

type Result struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data,omitempty"`
	Msg       string      `json:"msg"`
	Timestamp int64       `json:"timestamp"`
}
