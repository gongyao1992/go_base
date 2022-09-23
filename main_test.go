package xx_monitor

import (
	"testing"
	"xiangxin/monitor/app/pkg/setting"
)

func TestMonitor(t *testing.T)  {

	setting.Setup()

	t.Log(setting.KafkaSetting)
}