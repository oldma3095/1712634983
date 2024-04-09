package software

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"testing"
)

func TestCpu(t *testing.T) {
	info, err := cpu.Info()
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Logf("%+v\n", info)
}

func TestData(t *testing.T) {
	str := `
{
	"period": {
		"duration": 1000.175547,
		"unit": "ms"
	},
	"frequency": {
		"requested": 1301.771478,
		"actual": 1151.797805,
		"unit": "MHz"
	},
	"interrupts": {
		"count": 520.908556,
		"unit": "irq/s"
	},
	"rc6": {
		"value": 45.313833,
		"unit": "%"
	},
	"engines": {
		"Render/3D/0": {
			"busy": 0.000000,
			"sema": 0.000000,
			"wait": 0.000000,
			"unit": "%"
		},
		"Blitter/0": {
			"busy": 0.000000,
			"sema": 0.000000,
			"wait": 0.000000,
			"unit": "%"
		},
		"Video/0": {
			"busy": 39.320112,
			"sema": 0.000000,
			"wait": 0.000000,
			"unit": "%"
		},
		"Video/1": {
			"busy": 11.719689,
			"sema": 0.000000,
			"wait": 0.000000,
			"unit": "%"
		},
		"VideoEnhance/0": {
			"busy": 0.000000,
			"sema": 0.000000,
			"wait": 0.000000,
			"unit": "%"
		}
	}
},
{
	"period": {
		"duration": 1000.202160,
		"unit": "ms"
	},
	"frequency": {
		"requested": 1288.739468,
		"actual": 1180.761297,
		"unit": "MHz"
	},
	"interrupts": {
		"count": 535.891664,
		"unit": "irq/s"
	},
	"rc6": {
		"value": 48.217324,
		"unit": "%"
	},
	"engines": {
		"Render/3D/0": {
			"busy": 0.000000,
			"sema": 0.000000,
			"wait": 0.000000,
			"unit": "%"
		},
		"Blitter/0": {
			"busy": 0.000000,
			"sema": 0.000000,
			"wait": 0.000000,
			"unit": "%"
		},
		"Video/0": {
			"busy": 48.918471,
			"sema": 0.000000,
			"wait": 0.000000,
			"unit": "%"
		},
		"Video/1": {
			"busy": 0.000000,
			"sema": 0.000000,
			"wait": 0.000000,
			"unit": "%"
		},
		"VideoEnhance/0": {
			"busy": 0.000000,
			"sema": 0.000000,
			"wait": 0.000000,
			"unit": "%"
		}
	}
}
`
	info, err := handleRawIntelGPUInfos(str)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Logf("%+v", info)

}
