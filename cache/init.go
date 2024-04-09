package cache

import (
	cache "github.com/patrickmn/go-cache"
	"time"
)

var Cache *cache.Cache

func init() {
	// 初始化缓存
	Cache = cache.New(time.Hour*24, time.Minute*10)
	go func() {
		ticker := time.NewTicker(time.Second * 2)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				handleSystemInfo()
			}
		}
	}()
}
