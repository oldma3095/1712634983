package master

type Result struct {
	Base
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg"`
}
