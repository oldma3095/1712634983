package software

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/host"
	gnet "github.com/shirou/gopsutil/v3/net"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type SoftwareInfo struct {
	SystemHostname string `json:"systemHostname"` // 主机名
	SystemVersion  string `json:"systemVersion"`  // 系统版本
	SystemRuntime  string `json:"systemRuntime"`  // 系统运行时间
	SystemModel    string `json:"systemModel"`    // 系统型号
	KernelVersion  string `json:"kernelVersion"`  // 内核版本
	MAC            string `json:"mac"`            // MAC地址
}

func InitSoftwareInfo() (info SoftwareInfo) {
	info.SystemHostname, _ = os.Hostname()
	info.SystemVersion = systemVersion()
	info.SystemRuntime = systemRuntime()
	info.SystemModel = systemModel()
	info.KernelVersion = kernelVersion()
	info.MAC = mac()
	return
}

func systemVersion() string {
	c := exec.Command("uname", "-a")
	output, err := c.CombinedOutput()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}

func systemRuntime() string {
	bootTime, _ := host.BootTime()
	if bootTime > 0 {
		raw := time.Since(time.Unix(int64(bootTime), 0)).String()
		split := strings.Split(raw, ".")
		if len(split) > 0 {
			return split[0] + "s"
		} else {
			return raw
		}
	}
	return ""
}

func systemModel() string {
	return runtime.GOARCH
}

func kernelVersion() string {
	c := exec.Command("uname", "-r")
	output, err := c.CombinedOutput()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}

func mac() (macStr string) {
	interfaces, err := gnet.Interfaces()
	if err != nil {
		return ""
	}

	for _, stat := range interfaces {
		checkStr := fmt.Sprintf("%s", stat.Flags)
		if stat.HardwareAddr == "" || strings.Contains(stat.Name, "docker") {
			continue
		}
		if strings.Contains(checkStr, "up") && !strings.Contains(checkStr, "loopback") {
			if macStr == "" {
				macStr = stat.HardwareAddr
			} else {
				macStr += "," + stat.HardwareAddr
			}
		}
	}
	return
}
