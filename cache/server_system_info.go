package cache

import (
	"fmt"
	"time"
)

func getServerSystemInfoCacheKey(uuid string) string {
	return fmt.Sprintf("server_system_info:%s", uuid)
}

func GetServerSystemInfo(uuid string) (infos SystemInfo) {
	key := getServerSystemInfoCacheKey(uuid)
	load, b := Cache.Get(key)
	if b && load != nil {
		infos = load.(SystemInfo)
	}
	return infos
}

func SetServerSystemInfo(info SystemInfo) {
	key := getServerSystemInfoCacheKey(info.UUID)
	Cache.Set(key, info, time.Minute)
}
