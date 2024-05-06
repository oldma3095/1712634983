package master

type Handle struct {
	Method string `json:"method"`           // create delete update get
	Record bool   `json:"record,omitempty"` // 是否录制视频
}
