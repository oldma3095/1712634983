package software

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type rawIntelGPUInfo struct {
	Period struct {
		Duration float64 `json:"duration"`
		Unit     string  `json:"unit"`
	} `json:"period"`
	Frequency struct {
		Requested float64 `json:"requested"`
		Actual    float64 `json:"actual"`
		Unit      string  `json:"unit"`
	} `json:"frequency"`
	Interrupts struct {
		Count float64 `json:"count"`
		Unit  string  `json:"unit"`
	} `json:"interrupts"`
	Rc6 struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	} `json:"rc6"`
	Engines map[string]rawEngine `json:"engines"`
}

type rawEngine struct {
	Busy float64 `json:"busy"`
	Sema float64 `json:"sema"`
	Wait float64 `json:"wait"`
	Unit string  `json:"unit"`
}

type IntelGPUInfo struct {
	Period             string `json:"period"`             // 时间
	RequestedFrequency string `json:"requestedFrequency"` // 请求频率
	ActualFrequency    string `json:"actualFrequency"`    // 实际频率
	Interrupts         string `json:"interrupts"`         // 中断请求
	Rc6                string `json:"rc6"`                // rc6
	Engines            []Engine
}

type Engine struct {
	Name    string  `json:"name"`
	Busy    float64 `json:"busy"`    // 繁忙
	Sema    float64 `json:"sema"`    // 信号量
	Wait    float64 `json:"wait"`    // 等待
	BusyStr string  `json:"busyStr"` // 繁忙
	SemaStr string  `json:"semaStr"` // 信号量
	WaitStr string  `json:"waitStr"` // 等待
}

func InitIntelGPUTopInfo() (info IntelGPUInfo, err error) {
	c := exec.Command("/bin/bash", "-c", "timeout 1 intel_gpu_top -J")
	output, _ := c.CombinedOutput()
	info, err = handleRawIntelGPUInfos(string(output))
	return
}

func handleRawIntelGPUInfos(output string) (info IntelGPUInfo, err error) {
	data := fmt.Sprintf("[%s]", output)
	var rawInfos []rawIntelGPUInfo
	err = json.Unmarshal([]byte(data), &rawInfos)
	if err != nil {
		return
	}

	var rawInfo rawIntelGPUInfo
	if len(rawInfos) > 0 {
		rawInfo = rawInfos[0]
	}

	var engines []Engine
	for name, engine := range rawInfo.Engines {
		engines = append(engines, Engine{
			Name:    name,
			Busy:    engine.Busy,
			Sema:    engine.Sema,
			Wait:    engine.Wait,
			BusyStr: fmt.Sprintf("%.2f %s", engine.Busy, engine.Unit),
			SemaStr: fmt.Sprintf("%.2f %s", engine.Sema, engine.Unit),
			WaitStr: fmt.Sprintf("%.2f %s", engine.Wait, engine.Unit),
		})
	}
	info = IntelGPUInfo{
		Period:             fmt.Sprintf("%.2f %s", rawInfo.Period.Duration, rawInfo.Period.Unit),
		RequestedFrequency: fmt.Sprintf("%.2f %s", rawInfo.Frequency.Requested, rawInfo.Frequency.Unit),
		ActualFrequency:    fmt.Sprintf("%.2f %s", rawInfo.Frequency.Actual, rawInfo.Frequency.Unit),
		Interrupts:         fmt.Sprintf("%.2f %s", rawInfo.Interrupts.Count, rawInfo.Interrupts.Unit),
		Rc6:                fmt.Sprintf("%.2f %s", rawInfo.Rc6.Value, rawInfo.Period.Unit),
		Engines:            engines,
	}

	return
}
