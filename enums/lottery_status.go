package enums

//1: "开始下注", 2: "倒计时", 3: "封盘", 4: "核对账单", 5: "等待开奖", 6: "开奖结果", 7: "洗牌",

const (
	LotteryStatusPre       = iota
	LotteryStatusStart     // 开始下注
	LotteryStatusCountdown // 倒计时
	LotteryStatusStop      // 封盘
	LotteryStatusChecking  // 核对账单
	LotteryStatusOpen      // 等待开奖
	LotteryStatusResults   // 开奖
	LotteryStatusReload    // 结束
)
