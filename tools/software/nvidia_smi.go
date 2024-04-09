package software

import (
	"bytes"
	"encoding/csv"
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"os/exec"
)

type NvidiaSmiInfo struct {
	CudaVersion string `json:"cudaVersion"` // cuda可支持最高版本

	Index             string `json:"index"`
	Name              string `json:"name"`              // 名称
	TemperatureGPU    string `json:"temperatureGPU"`    // 温度
	UtilizationGPU    string `json:"utilizationGPU"`    // gpu使用率
	UtilizationMemory string `json:"utilizationMemory"` // 显存使用率
	MemoryTotal       string `json:"memoryTotal"`       // 显存总量
	MemoryFree        string `json:"memoryFree"`        // 显存可用
	MemoryUsed        string `json:"memoryUsed"`        // 显存已使用
	PowerDraw         string `json:"powerDraw"`         // 功耗
	DriverVersion     string `json:"driverVersion"`     // 显卡驱动
}

type NvidiaSmiLog struct {
	XMLName       xml.Name `xml:"nvidia_smi_log"`
	Text          string   `xml:",chardata"`
	Timestamp     string   `xml:"timestamp"`
	DriverVersion string   `xml:"driver_version"`
	CudaVersion   string   `xml:"cuda_version"`
	HicInfo       string   `xml:"hic_info"`
	AttachedUnits string   `xml:"attached_units"`
}

func InitNvidiaSmiInfos() (infos []NvidiaSmiInfo, err error) {
	var cudaVersion string
	cCuda := exec.Command("nvidia-smi", "-q", "-u", "-x")
	combinedOutput, err := cCuda.CombinedOutput()
	if err != nil {
		return
	}
	var nvidiaSmiLog NvidiaSmiLog
	decoder := xml.NewDecoder(bytes.NewReader(combinedOutput))
	decoder.CharsetReader = charset.NewReaderLabel
	_ = decoder.Decode(&nvidiaSmiLog)
	cudaVersion = nvidiaSmiLog.CudaVersion

	c := exec.Command(
		"nvidia-smi",
		"--query-gpu=index,gpu_name,temperature.gpu,utilization.gpu,utilization.memory,memory.total,memory.free,memory.used,power.draw,driver_version",
		//"--format=csv,noheader,nounits",
		"--format=csv,noheader",
	)
	output, err := c.CombinedOutput()
	if err != nil {
		return
	}

	csvReader := csv.NewReader(bytes.NewReader(output))
	csvReader.TrimLeadingSpace = true
	records, err := csvReader.ReadAll()
	if err != nil {
		return
	}

	for _, rows := range records {
		if len(rows) == 10 {
			infos = append(infos, NvidiaSmiInfo{
				CudaVersion:       cudaVersion,
				Index:             rows[0],
				Name:              rows[1],
				TemperatureGPU:    rows[2],
				UtilizationGPU:    rows[3],
				UtilizationMemory: rows[4],
				MemoryTotal:       rows[5],
				MemoryFree:        rows[6],
				MemoryUsed:        rows[7],
				PowerDraw:         rows[8],
				DriverVersion:     rows[9],
			})
		}
	}
	return
}
