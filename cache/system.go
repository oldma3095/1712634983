package cache

import (
	"github.com/oldma3095/1712634983/tools/software"
	"github.com/shirou/gopsutil/v3/net"
	"time"
)

type SystemInfo struct {
	Cpu      Cpu             `json:"cpu"`
	Ram      Ram             `json:"ram"`
	Disk     Disk            `json:"disk"`
	Nvidia   []NvidiaSmiInfo `json:"nvidia"`
	IntelGPU IntelGPU        `json:"intelGPU"`
	NetSpeed NetSpeed        `json:"netSpeed"`
	Software Software        `json:"software"`
}

type Software struct {
	SystemHostname string `json:"systemHostname"` // 主机名
	SystemVersion  string `json:"systemVersion"`  // 系统版本
	SystemRuntime  string `json:"systemRuntime"`  // 系统运行时间
	SystemModel    string `json:"systemModel"`    // 系统型号
	KernelVersion  string `json:"kernelVersion"`  // 内核版本
	MAC            string `json:"mac"`            // MAC地址
}

type Cpu struct {
	Name        string    `json:"name"`
	Cpus        []float64 `json:"cpus"`
	Cores       int32     `json:"cores"`
	Temperature uint32    `json:"temperature"`
}

type Ram struct {
	UsedMb      int32 `json:"usedMb"`
	TotalMb     int32 `json:"totalMb"`
	UsedPercent int32 `json:"usedPercent"`
}

type Disk struct {
	UsedMb      int32 `json:"usedMb"`
	UsedGb      int32 `json:"usedGb"`
	TotalMb     int32 `json:"totalMb"`
	TotalGb     int32 `json:"totalGb"`
	UsedPercent int32 `json:"usedPercent"`
}

type NvidiaSmiInfo struct {
	Index             string `json:"index"`
	Name              string `json:"name"`              // 名称
	TemperatureGPU    string `json:"temperatureGPU"`    // 温度
	UtilizationGPU    string `json:"utilizationGPU"`    // gpu使用率
	UtilizationMemory string `json:"utilizationMemory"` // 显存使用率
	MemoryTotal       string `json:"memoryTotal"`       // 显存总量
	MemoryFree        string `json:"memoryFree"`        // 显存可用
	MemoryUsed        string `json:"memoryUsed"`        // 显存已使用
	CudaVersion       string `json:"cudaVersion"`       // cuda可用最高版本
	PowerDraw         string `json:"powerDraw"`         // 功耗
	DriverVersion     string `json:"driverVersion"`     // 驱动版本
}

type IntelGPU struct {
	Period             string   `json:"period"`             // 时间
	RequestedFrequency string   `json:"requestedFrequency"` // 请求频率
	ActualFrequency    string   `json:"actualFrequency"`    // 实际频率
	Interrupts         string   `json:"interrupts"`         // 中断请求
	Rc6                string   `json:"rc6"`                // rc6
	Engines            []Engine `json:"engines"`
}

type Engine struct {
	Name    string  `json:"name"`
	Busy    float32 `json:"busy"`    // 繁忙
	Sema    float32 `json:"sema"`    // 信号量
	Wait    float32 `json:"wait"`    // 等待
	BusyStr string  `json:"busyStr"` // 繁忙
	SemaStr string  `json:"semaStr"` // 信号量
	WaitStr string  `json:"waitStr"` // 等待
}

type NetSpeed struct {
	SendSpeed float32 `json:"sendSpeed"` // 发送速度
	RecvSpeed float32 `json:"recvSpeed"` // 接收速度
}

var systemInfoCK = "system_info"
var lastSentBytes uint64
var lastRecvBytes uint64
var lastTime int64
var lastSentSpeed float32
var lastRecvSpeed float32

func GetSystemInfo() (infos SystemInfo) {
	load, b := Cache.Get(systemInfoCK)
	if b && load != nil {
		infos = load.(SystemInfo)
	}
	return infos
}

func handleSystemInfo() SystemInfo {
	initCpu, _ := software.InitCPU()
	initRAM, _ := software.InitRAM()
	initDisk, _ := software.InitDisk()
	initNvidiaSmiInfos, _ := software.InitNvidiaSmiInfos()
	initIntelGPUInfo, _ := software.InitIntelGPUTopInfo()
	softwareInfo := software.InitSoftwareInfo()

	// nvidia gpu
	var nvidiaInfos []NvidiaSmiInfo
	for _, nvidiaSmiInfo := range initNvidiaSmiInfos {
		nvidiaInfos = append(nvidiaInfos, NvidiaSmiInfo{
			Index:             nvidiaSmiInfo.Index,
			Name:              nvidiaSmiInfo.Name,
			TemperatureGPU:    nvidiaSmiInfo.TemperatureGPU,
			UtilizationGPU:    nvidiaSmiInfo.UtilizationGPU,
			UtilizationMemory: nvidiaSmiInfo.UtilizationMemory,
			MemoryTotal:       nvidiaSmiInfo.MemoryTotal,
			MemoryFree:        nvidiaSmiInfo.MemoryFree,
			MemoryUsed:        nvidiaSmiInfo.MemoryUsed,
			CudaVersion:       nvidiaSmiInfo.CudaVersion,
			PowerDraw:         nvidiaSmiInfo.PowerDraw,
			DriverVersion:     nvidiaSmiInfo.DriverVersion,
		})
	}

	// inter gpu
	var intelGPUEngines []Engine
	for _, engine := range initIntelGPUInfo.Engines {
		intelGPUEngines = append(intelGPUEngines, Engine{
			Name:    engine.Name,
			Busy:    float32(engine.Busy),
			Sema:    float32(engine.Sema),
			Wait:    float32(engine.Wait),
			BusyStr: engine.BusyStr,
			SemaStr: engine.SemaStr,
			WaitStr: engine.WaitStr,
		})
	}

	// 网络速率
	var netSpeed NetSpeed
	timeNow := time.Now().Unix()
	// net
	counters, _ := net.IOCounters(false)
	currentSent := counters[0].BytesSent
	currentRecv := counters[0].BytesRecv
	t := timeNow - lastTime
	if t == 0 {
		netSpeed.SendSpeed = lastSentSpeed
		netSpeed.RecvSpeed = lastRecvSpeed
	} else {
		sendSpeed := float32((currentSent - lastSentBytes) / uint64(t))
		recvSpeed := float32((currentRecv - lastRecvBytes) / uint64(t))
		netSpeed.SendSpeed = sendSpeed
		netSpeed.RecvSpeed = recvSpeed
		lastSentSpeed = sendSpeed
		lastRecvSpeed = recvSpeed
	}
	lastTime = timeNow
	lastSentBytes = currentSent
	lastRecvBytes = currentRecv

	info := SystemInfo{
		Cpu: Cpu{
			Name:        initCpu.Name,
			Cpus:        initCpu.Cpus,
			Cores:       int32(initCpu.Cores),
			Temperature: initCpu.Temperature,
		},
		Ram: Ram{
			UsedMb:      int32(initRAM.UsedMB),
			TotalMb:     int32(initRAM.TotalMB),
			UsedPercent: int32(initRAM.UsedPercent),
		},
		Disk: Disk{
			UsedMb:      int32(initDisk.UsedMB),
			UsedGb:      int32(initDisk.UsedGB),
			TotalMb:     int32(initDisk.TotalMB),
			TotalGb:     int32(initDisk.TotalGB),
			UsedPercent: int32(initDisk.UsedPercent),
		},
		Nvidia: nvidiaInfos,
		IntelGPU: IntelGPU{
			Period:             initIntelGPUInfo.Period,
			RequestedFrequency: initIntelGPUInfo.RequestedFrequency,
			ActualFrequency:    initIntelGPUInfo.ActualFrequency,
			Interrupts:         initIntelGPUInfo.Interrupts,
			Rc6:                initIntelGPUInfo.Rc6,
			Engines:            intelGPUEngines,
		},
		NetSpeed: netSpeed,
		Software: Software{
			SystemHostname: softwareInfo.SystemHostname,
			SystemVersion:  softwareInfo.SystemVersion,
			SystemRuntime:  softwareInfo.SystemRuntime,
			SystemModel:    softwareInfo.SystemModel,
			KernelVersion:  softwareInfo.KernelVersion,
			MAC:            softwareInfo.MAC,
		},
	}
	Cache.Set(systemInfoCK, info, time.Minute)
	return info
}
