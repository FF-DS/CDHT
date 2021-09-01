package ReportModule

import (
	"time"
    "encoding/json"
    "net/http"
    "bytes"
	"fmt"
)


type Logger struct {
	routingTable Log
	nodeLogChannel []Log
	networkToolChannel  []Log
	configurationToolChannel []Log

	config LogConfig
}


func (logger *Logger) Configure(config LogConfig) {
	logger.config = config
}


func (logger *Logger) Init() {
	logger.nodeLogChannel = []Log{}
	logger.networkToolChannel =  []Log{}
	logger.configurationToolChannel = []Log{}

	go logger.reportRouteTableInfo()
	go logger.reportNodeLogChannel()
	go logger.reportNetToolInfo()
	go logger.reportConfigInfo()
}



// --------------------------- Send to the api --------------------------- //
func  (logger *Logger) RouteTableLog(logData Log) {
	logger.routingTable = logData
}

func (logger *Logger) RouteLogs() []Log {
	return []Log{ logger.routingTable }
}



func  (logger *Logger) NodeLog(logData Log) {
	if( len(logger.nodeLogChannel) >= logger.config.NodeChanSize() -2 ){
		logger.nodeLogChannel = logger.nodeLogChannel[1:]
	}
	logger.nodeLogChannel = append(logger.nodeLogChannel, logData)
}

func (logger *Logger) NodeLogs() []Log {
	return logger.nodeLogChannel
}



func  (logger *Logger) NetworkToolLog(logData Log) {
	if( len(logger.networkToolChannel) >= logger.config.NetChanSize() -2 ){
		logger.networkToolChannel = logger.networkToolChannel[1:]
	}
	logger.networkToolChannel = append(logger.networkToolChannel, logData)
}

func (logger *Logger) NetworkToolLogs() []Log {
	return logger.networkToolChannel 
}



func  (logger *Logger) ConfigToolLog(logData Log) {
	if( len(logger.configurationToolChannel) >= logger.config.ConfigChanSize() -2 ){
		logger.configurationToolChannel = logger.configurationToolChannel[1:]
	}
	logger.configurationToolChannel = append(logger.configurationToolChannel, logData)
}

func (logger *Logger) ConfigToolLogs() []Log {
	return logger.configurationToolChannel
}



// --------------------------- Send to the api --------------------------- //
func (logger *Logger) reportRouteTableInfo() {
	for {
		time.Sleep(time.Second * logger.config.RouteTableDelay() )
		
		// logs := []Log{ logger.routingTable }
		// sendLogToAPI( logs, URL_SEND_ROUTE_TABLE_LOG)
	}
}


func (logger *Logger) reportNodeLogChannel() {
	for {
		time.Sleep(time.Second * logger.config.NodeChanDelay() )
		
		// logs := logger.nodeLogChannel[:logger.config.NodeChanSize()]
		// sendLogToAPI(logs, URL_SEND_NODE_INFO_LOG)
	}
}


func (logger *Logger) reportNetToolInfo() {
	for {
		time.Sleep(time.Second * logger.config.NetChanDelay() )
		
		// logs := logger.networkToolChannel[:logger.config.NetChanSize()]
		// sendLogToAPI(logs, URL_SEND_NETWOR_TOOL_LOG)
	}
}


func (logger *Logger) reportConfigInfo() {
	for {
		time.Sleep(time.Second * logger.config.ConfigChanDelay() )
		
		// logs := logger.configurationToolChannel[:logger.config.ConfigChanSize()]
		// sendLogToAPI(logs, URL_SEND_CONFIG_TOOL_LOG)
	}
}





// --------------------------- helper method --------------------------- //

// api caller 
func sendLogToAPI(logger []Log, url_link string) bool {
    postBody, _ := json.Marshal( logger )

    responseBody := bytes.NewBuffer(postBody)
    resp, err := http.Post(url_link, "application/json", responseBody)
    
    if err != nil {
		fmt.Println("[API-LOG][Error]: +", err)
       	return false
    }
    
    defer resp.Body.Close()
	return true
}
