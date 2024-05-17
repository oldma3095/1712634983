package master

type ClientInfo struct {
	Base
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg"`
}
