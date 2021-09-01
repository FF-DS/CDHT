package ReportModule


import (
	"math/big"
	"fmt"
)
// # --------------------------- eunm --------------------------- # //

const (
	LOG_OPERATION_STATUS_SUCCESS 	 string = "SUCCESS"
	LOG_OPERATION_STATUS_FAILED      string = "FIALED"

	LOG_TYPE_ROUTING_TABLE           string = "TYPE_ROUTING_TABLE"
	LOG_TYPE_NODE_INFORMATION        string = "TYPE_NODE_INFORMATION"
	LOG_TYPE_NETWORK_TOOL_FORWARDING string = "TYPE_NETWORK_TOOL"
	LOG_TYPE_APP_SERVICE             string = "TYPE_APP_SERVICE"


	LOG_LOCATION_TYPE_INCOMMING      string = "LOCATION_TYPE_INCOMMING"
	LOG_LOCATION_TYPE_LEAVING        string = "LOCATION_TYPE_LEAVING"
	LOG_LOCATION_TYPE_SELF           string = "LOCATION_TYPE_SELF"


	URL_SEND_ROUTE_TABLE_LOG         string = "https://cdht-monitoring-api.herokuapp.com/route-log"
	URL_SEND_NODE_INFO_LOG           string = "https://cdht-monitoring-api.herokuapp.com/node-log"
	URL_SEND_NETWOR_TOOL_LOG         string = "https://cdht-monitoring-api.herokuapp.com/net-log"
	URL_SEND_CONFIG_TOOL_LOG         string = "https://cdht-monitoring-api.herokuapp.com/config-log"
)


// #--------------------------------- Log Object ---------------------------------# //

type Log struct {
	Type  string
	OperationStatus  string
	LogLocation  string
	NodeId *big.Int
	NodeAddress string
	LogBody interface{}
}



// #--------------------------------- Opreation on log ---------------------------------# //

func  (log *Log) GetLogName() string {
	return log.Type
}



// #--------------------------------- string log ---------------------------------# //
func  (log *Log) ToString() string {
	str := "---------------- Log Data ----------------\n"  
	str += fmt.Sprintf(" [+] Operation Type : %s\n", log.Type )
	str += fmt.Sprintf(" [+] Operation Status : %s\n", log.OperationStatus )
	str += fmt.Sprintf(" [+] Node Id : %s\n", log.NodeId.String() )
	str += fmt.Sprintf(" [+] Node Address : %s\n", log.NodeAddress ) 
	str += fmt.Sprintf(" [+] Log Body : %s\n", log.LogBody )
	str += "------------------------------------------\n"  
	return str;
}