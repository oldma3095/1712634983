package slave

import "github.com/oldma3095/1712634983/cache"

type Result struct {
	Code int32              `json:"code"`
	Msg  string             `json:"msg"`
	Data cache.NiuNiuResult `json:"data"`
}
