package client

import (
	"context"
	"fmt"
	"go_poker/grpc/cache"
	commonApi "go_poker/grpc/protos/common/api"
	"time"
)

func (client *Clients) PushClientInfoToMaster() {
	var err error
	defer func() {
		if err != nil {
			client.exit <- err.Error()
		}
	}()
	stream, err := client.apiServiceClient.ClientInfo(context.Background())
	if err != nil {
		return
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-stream.Context().Done():
			err = fmt.Errorf("服务端主动断开")
			return
		case <-ticker.C:
			systemInfo := cache.GetSystemInfo()

			var nvidiaInfos []*commonApi.Nvidia
			for _, nvidiaSmiInfo := range systemInfo.Nvidia {
				nvidiaInfos = append(nvidiaInfos, &commonApi.Nvidia{
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

			var intelGPUEngines []*commonApi.Engine
			for _, engine := range systemInfo.IntelGPU.Engines {
				intelGPUEngines = append(intelGPUEngines, &commonApi.Engine{
					Name:    engine.Name,
					Busy:    engine.Busy,
					Sema:    engine.Sema,
					Wait:    engine.Wait,
					BusyStr: engine.BusyStr,
					SemaStr: engine.SemaStr,
					WaitStr: engine.WaitStr,
				})
			}

			info := commonApi.ClientInfoReq{
				Cpu: &commonApi.Cpu{
					Cpus:        systemInfo.Cpu.Cpus,
					Cores:       systemInfo.Cpu.Cores,
					Name:        systemInfo.Cpu.Name,
					Temperature: systemInfo.Cpu.Temperature,
				},
				Ram: &commonApi.Ram{
					UsedMb:      systemInfo.Ram.UsedMb,
					TotalMb:     systemInfo.Ram.TotalMb,
					UsedPercent: systemInfo.Ram.UsedPercent,
				},
				Disk: &commonApi.Disk{
					UsedMb:      systemInfo.Disk.UsedMb,
					UsedGb:      systemInfo.Disk.UsedGb,
					TotalMb:     systemInfo.Disk.TotalMb,
					TotalGb:     systemInfo.Disk.TotalGb,
					UsedPercent: systemInfo.Disk.UsedPercent,
				},
				Nvidia: nvidiaInfos,
				IntelGPU: &commonApi.IntelGPU{
					Period:             systemInfo.IntelGPU.Period,
					RequestedFrequency: systemInfo.IntelGPU.RequestedFrequency,
					ActualFrequency:    systemInfo.IntelGPU.ActualFrequency,
					Interrupts:         systemInfo.IntelGPU.Interrupts,
					Rc6:                systemInfo.IntelGPU.Rc6,
					Engines:            intelGPUEngines,
				},
				NetSpeed: &commonApi.NetSpeed{
					SendSpeed: systemInfo.NetSpeed.SendSpeed,
					RecvSpeed: systemInfo.NetSpeed.RecvSpeed,
				},
				Software: &commonApi.Software{
					SystemHostname: systemInfo.Software.SystemHostname,
					SystemVersion:  systemInfo.Software.SystemVersion,
					SystemRuntime:  systemInfo.Software.SystemRuntime,
					SystemModel:    systemInfo.Software.SystemModel,
					KernelVersion:  systemInfo.Software.KernelVersion,
					Mac:            systemInfo.Software.MAC,
				},
			}
			err = stream.SendMsg(&info)
			if err != nil {
				return
			}
		}
	}
}
