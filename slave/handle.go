package slave

import "github.com/oldma3095/1712634983/cache"

type Handle struct {
	Init bool   `json:"init"` // 首次连接初始化
	UUID string `json:"uuid"`
}

type HandlePredict struct {
	Code int32              `json:"code"`
	Msg  string             `json:"msg"`
	Data cache.NiuNiuResult `json:"data"`
}

type HandleRecord struct {
	RecordFileUrl string `json:"recordFileUrl"` // 录制后返回的视频地址
	Msg           string `json:"msg"`           // 错误信息
}
