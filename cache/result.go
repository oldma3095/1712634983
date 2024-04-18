package cache

import "fmt"

type NiuNiuResult struct {
	Id            uint64   `json:"id"`
	SaveTime      int64    `json:"saveTime"`
	Image         string   `json:"image"`
	RawImage      string   `json:"rawImage"`
	Video         string   `json:"video"`
	VideoSize     int64    `json:"videoSize"`
	SaveTimeStr   string   `json:"saveTimeStr"`
	Flag          []string `json:"flag"`
	Banker        []string `json:"banker"`
	Player1       []string `json:"player1"`
	Player2       []string `json:"player2"`
	Player3       []string `json:"player3"`
	Other         []string `json:"other"`
	SimpleFlag    string   `json:"simpleFlag"`
	SimpleBanker  string   `json:"simpleBanker"`
	SimplePlayer1 string   `json:"simplePlayer1"`
	SimplePlayer2 string   `json:"simplePlayer2"`
	SimplePlayer3 string   `json:"simplePlayer3"`
	SimpleOther   string   `json:"simpleOther"`
}

func getGameDataCacheKey() string {
	return fmt.Sprintf("game_data")
}

func GetNiuNiuResultData() (data NiuNiuResult) {
	key := getGameDataCacheKey()
	get, b := Cache.Get(key)
	if b && get != nil {
		data = get.(NiuNiuResult)
		return
	}
	return
}

func SetNiuNiuResultData(data NiuNiuResult) {
	key := getGameDataCacheKey()
	Cache.Set(key, data, -1)
}
