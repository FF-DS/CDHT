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

	URLSendRouteTableLogValue string
	URLSendNodeInfoLogValue string
	URLSendNetworkToolLogValue string
	URLSendConfigToolLogValue string

}

func (conf *LogConfig) RouteTableDelay() time.Duration {
	if conf.RouteTableDelayValue == 0 {
		return 15
	}
	return conf.RouteTableDelayValue
}


func (conf *LogConfig) NodeChanSize() int {
	if conf.NodeChanSizeValue == 0 {
		return 100
	}
	return conf.NodeChanSizeValue
}

func (conf *LogConfig) NodeChanDelay() time.Duration {
	if conf.NodeChanDelayValue == 0 {
		return 15
	}
	return conf.NodeChanDelayValue
}

func (conf *LogConfig) NetChanSize() int {
	if conf.NetChanSizeValue == 0 {
		return 100
	}
	return conf.NetChanSizeValue
}

func (conf *LogConfig) NetChanDelay() time.Duration {
	if conf.NetChanDelayValue == 0 {
		return 15
	}
	return conf.NetChanDelayValue
}

func (conf *LogConfig) ConfigChanSize() int {
	if conf.ConfigChanSizeValue == 0 {
		return 100
	}
	return conf.ConfigChanSizeValue
}

func (conf *LogConfig) ConfigChanDelay() time.Duration {
	if conf.ConfigChanDelayValue == 0 {
		return 15
	}
	return conf.ConfigChanDelayValue
}



// url links
func (conf *LogConfig) URLSendRouteTableLog() string {
	if conf.URLSendRouteTableLogValue == "" {
		return URL_SEND_ROUTE_TABLE_LOG
	}
	return conf.URLSendRouteTableLogValue
}

func (conf *LogConfig) URLSendNodeInfoLog() string {
	if conf.URLSendNodeInfoLogValue == "" {
		return URL_SEND_NODE_INFO_LOG
	}
	return conf.URLSendNodeInfoLogValue
}

func (conf *LogConfig) URLSendNetworkToolLog() string {
	if conf.URLSendNetworkToolLogValue == "" {
		return URL_SEND_NETWORK_TOOL_LOG
	}
	return conf.URLSendNetworkToolLogValue
}

func (conf *LogConfig) URLSendConfigToolLog() string {
	if conf.URLSendConfigToolLogValue == "" {
		return URL_SEND_CONFIG_TOOL_LOG
	}
	return conf.URLSendConfigToolLogValue
}