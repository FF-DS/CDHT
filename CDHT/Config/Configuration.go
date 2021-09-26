package Config

import (
	"cdht/ReportModule"
	"math/big"
	"strconv"
	"time"
)

const (
	MODE_CREATE_RING = "FIRST_NODE"
	MODE_JOIN_RING = "SECOND_NODE"
	MODE_REPLICA_NODE = "REPLICA_NODE"
)


type Configuration struct {
	ConfigID int
	
	// mode of the application
	Application_Mode string

	RUN_TCP_TEST_APPLICATION bool
	RUN_UDP_TEST_APPLICATION bool
	RUN_NETWORK_TEST_APPLICATION bool

	// node level & routing information
	Node_Port string
	Node_IP string
	Node_ID string
	Node_M  string
	Jump_Spacing int
	Finger_Table_Length int
	Routing_Update_Delay time.Duration
	Successors_Table_Length int

	// remote node connection  format `ip:port`
	Remote_Node_Address string


	// application
	Application_Channel_Size int
	Application_Connecting_Port string


	// log configuration
	Log_Route_Table_Delay_Value time.Duration
	Log_Node_Channel_Size_Value int
	Log_Node_Channel_Delay_Value time.Duration
	
	Log_Network_Channel_Size_Value int
	Log_Network_Channel_Delay_Value time.Duration

	Log_Configuration_Channel_Size_Value int
	Log_Configuration_Channel_Delay_Value time.Duration

	Log_URL_Send_Route_Table string
	Log_URL_Send_Node_Info string
	Log_URL_Send_Network_Tool string
	Log_URL_Send_Configuration_Tool string


	// Network tool [Internal]
	Network_Tool_Channel_Size int


	// CDHT Network tool [External]
	CDHT_Ping_Tool_Listening_Port string
	CDHT_API_Communication_Delay time.Duration
	CDHT_Command_Channel_Size int
	CDHT_URL_Get_Command_From_Server string
	CDHT_URL_Send_Command_Result string


	// TEST Application
	TEST_APP_IPAddress string
	TEST_APP_Port string
	TEST_APP_UDP_Listener_Port string
	TEST_APP_Net_Channel_Size int
	TEST_APP_AppName string
	TEST_APP_Packet_Delay  time.Duration 
	TEST_APP_RUNNING_TIME  time.Duration

	CONFIGURATION_DOWNLOAD_DELAY  time.Duration
	CONFIGURATION_DOWNLOAD_URL string

	REPLICATION_COUNT int
}


func (config *Configuration) GetLogConfiguration() ReportModule.LogConfig {
	return ReportModule.LogConfig {
		RouteTableDelayValue: config.Log_Route_Table_Delay_Value,
		NodeChanSizeValue: config.Log_Node_Channel_Size_Value,
		NodeChanDelayValue: config.Log_Node_Channel_Delay_Value,
		
		NetChanSizeValue: config.Log_Network_Channel_Size_Value,
		NetChanDelayValue: config.Log_Network_Channel_Delay_Value,
		
		ConfigChanSizeValue: config.Log_Configuration_Channel_Size_Value,
		ConfigChanDelayValue: config.Log_Configuration_Channel_Delay_Value,

		URLSendRouteTableLogValue: config.Log_URL_Send_Route_Table,
		URLSendNodeInfoLogValue: config.Log_URL_Send_Node_Info,
		URLSendNetworkToolLogValue: config.Log_URL_Send_Network_Tool,
		URLSendConfigToolLogValue: config.Log_URL_Send_Configuration_Tool,
	}
}

func (config *Configuration) GetNodeM() *big.Int {
	M, ok := new(big.Int).SetString(config.Node_M, 10)
    if !ok { 
		return nil
	}

	config.Finger_Table_Length, _ = strconv.Atoi(config.Node_M)
	return M
}

func (config *Configuration) GetNodeID() *big.Int {
	ID, ok := new(big.Int).SetString(config.Node_ID, 10)
    if !ok { 
		return nil
	}

	return ID
}