package master

type Base struct {
	CanRecordNow bool  `json:"canRecordNow"`
	Timestamp    int64 `json:"timestamp"`
}
