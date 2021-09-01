package ReportModule

import (
	"time"
)

type LogConfig struct {
	RouteTableDelayValue time.Duration

	NodeChanSizeValue int
	NodeChanDelayValue time.Duration
	
	NetChanSizeValue int
	NetChanDelayValue time.Duration
	
	ConfigChanSizeValue int
	ConfigChanDelayValue time.Duration
}

func (conf *LogConfig) RouteTableDelay() time.Duration {
	if conf.RouteTableDelayValue == 0 {
		return 300
	}
	return conf.RouteTableDelayValue
}


func (conf *LogConfig) NodeChanSize() int {
	if conf.NodeChanSizeValue == 0 {
		return 100000
	}
	return conf.NodeChanSizeValue
}

func (conf *LogConfig) NodeChanDelay() time.Duration {
	if conf.NodeChanDelayValue == 0 {
		return 300
	}
	return conf.NodeChanDelayValue
}

func (conf *LogConfig) NetChanSize() int {
	if conf.NetChanSizeValue == 0 {
		return 100000
	}
	return conf.NetChanSizeValue
}

func (conf *LogConfig) NetChanDelay() time.Duration {
	if conf.NetChanDelayValue == 0 {
		return 300
	}
	return conf.NetChanDelayValue
}

func (conf *LogConfig) ConfigChanSize() int {
	if conf.ConfigChanSizeValue == 0 {
		return 100000
	}
	return conf.ConfigChanSizeValue
}

func (conf *LogConfig) ConfigChanDelay() time.Duration {
	if conf.ConfigChanDelayValue == 0 {
		return 300
	}
	return conf.ConfigChanDelayValue
}