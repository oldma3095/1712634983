package req

type Result struct {
	Code int32      `json:"code"`
	Msg  string     `json:"msg"`
	Data ResultData `json:"data"`
}

type ResultData struct {
	ID            uint     `json:"id"`
	SaveTime      int64    `json:"saveTime"`  // 保存时间戳
	Image         string   `json:"image"`     // 图片地址
	RawImage      string   `json:"rawImage"`  // 图片地址
	Video         string   `json:"video"`     // 录制视频地址
	VideoSize     int64    `json:"videoSize"` // 视频大小
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
