package api

import (
	"fmt"
	"go_poker/grpc/cache"
	commonApi "go_poker/grpc/protos/common/api"
	"go_poker/grpc/tools"
	"log"
	"time"
)

type CommonServer struct {
	commonApi.UnimplementedApiServiceServer
}

func (s *CommonServer) ClientInfo(stream commonApi.ApiService_ClientInfoServer) error {
	ip, errGRPCClientIP := tools.GRPCClientIP(stream.Context())
	if errGRPCClientIP != nil {
		return errGRPCClientIP
	}
	log.Println("ServerInfoGreeter", ip)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-stream.Context().Done():
			return fmt.Errorf("客户端主动断开")
		case <-ticker.C:
			info, err := stream.Recv()
			if err != nil {
				return err
			}

			var nvidiaInfos []cache.NvidiaSmiInfo
			for _, nvidiaInfo := range info.Nvidia {
				nvidiaInfos = append(nvidiaInfos, cache.NvidiaSmiInfo{
					Index:             nvidiaInfo.Index,
					Name:              nvidiaInfo.Name,
					TemperatureGPU:    nvidiaInfo.TemperatureGPU,
					UtilizationGPU:    nvidiaInfo.UtilizationGPU,
					UtilizationMemory: nvidiaInfo.UtilizationMemory,
					MemoryTotal:       nvidiaInfo.MemoryTotal,
					MemoryFree:        nvidiaInfo.MemoryFree,
					MemoryUsed:        nvidiaInfo.MemoryUsed,
					CudaVersion:       nvidiaInfo.CudaVersion,
					PowerDraw:         nvidiaInfo.PowerDraw,
					DriverVersion:     nvidiaInfo.DriverVersion,
				})
			}

			var engines []cache.Engine
			for _, engine := range info.IntelGPU.Engines {
				engines = append(engines, cache.Engine{
					Name:    engine.Name,
					Busy:    engine.Busy,
					Sema:    engine.Sema,
					Wait:    engine.Wait,
					BusyStr: engine.BusyStr,
					SemaStr: engine.SemaStr,
					WaitStr: engine.WaitStr,
				})
			}

			serverInfo := cache.SystemInfo{
				Cpu: cache.Cpu{
					Name:        info.Cpu.Name,
					Cpus:        info.Cpu.Cpus,
					Cores:       info.Cpu.Cores,
					Temperature: info.Cpu.Temperature,
				},
				Ram: cache.Ram{
					UsedMb:      info.Ram.UsedMb,
					TotalMb:     info.Ram.TotalMb,
					UsedPercent: info.Ram.UsedPercent,
				},
				Disk: cache.Disk{
					UsedMb:      info.Disk.UsedMb,
					UsedGb:      info.Disk.UsedGb,
					TotalMb:     info.Disk.TotalMb,
					TotalGb:     info.Disk.TotalGb,
					UsedPercent: info.Disk.UsedPercent,
				},
				Nvidia: nvidiaInfos,
				IntelGPU: cache.IntelGPU{
					Period:             info.IntelGPU.Period,
					RequestedFrequency: info.IntelGPU.RequestedFrequency,
					ActualFrequency:    info.IntelGPU.ActualFrequency,
					Interrupts:         info.IntelGPU.Interrupts,
					Rc6:                info.IntelGPU.Rc6,
					Engines:            engines,
				},
				NetSpeed: cache.NetSpeed{
					SendSpeed: info.NetSpeed.SendSpeed,
					RecvSpeed: info.NetSpeed.RecvSpeed,
				},
				Software: cache.Software{
					SystemHostname: info.Software.SystemHostname,
					SystemVersion:  info.Software.SystemVersion,
					SystemRuntime:  info.Software.SystemRuntime,
					SystemModel:    info.Software.SystemModel,
					KernelVersion:  info.Software.KernelVersion,
					MAC:            info.Software.Mac,
				},
			}

			log.Printf("%+v\n", serverInfo)
		}
	}
}
