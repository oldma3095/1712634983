package slave

import "github.com/oldma3095/1712634983/cache"

type Base struct {
	UUID      string `json:"uuid"`
	Timestamp int64  `json:"timestamp"`
}

type Handle struct {
	Base
	Init bool `json:"init"` // 首次连接初始化
}

type HandlePredict struct {
	Base
	Code int32              `json:"code"`
	Msg  string             `json:"msg"`
	Data cache.NiuNiuResult `json:"data"`
}

type HandleRecord struct {
	Base
	RecordFileUrl string `json:"recordFileUrl"` // 录制后返回的视频地址
	Msg           string `json:"msg"`           // 错误信息
}
