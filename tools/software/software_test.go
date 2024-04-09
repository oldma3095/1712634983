package software

import "testing"

func TestSoftwareInfo(t *testing.T) {
	info := InitSoftwareInfo()
	t.Logf("%+v", info)
}
