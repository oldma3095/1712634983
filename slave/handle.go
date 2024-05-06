package slave

import "github.com/oldma3095/1712634983/cache"

type Handle struct {
	Code int32              `json:"code"`
	Msg  string             `json:"msg"`
	Data cache.NiuNiuResult `json:"data"`
	Init bool               `json:"init"` // 首次连接初始化
}
