package slave

type Handle struct {
	Code int32      `json:"code"`
	Msg  string     `json:"msg"`
	Data ResultData `json:"data"`
}
