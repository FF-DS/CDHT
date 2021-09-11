package CoreModule

import (

)

const (
	MODE_CREATE_RING = "FIRST_NODE"
	MODE_JOIN_RING = "SECOND_NODE"
	MODE_UDP_TEST_APPLICATION = "TEST_APPLICATION"
	MODE_TCP_TEST_APPLICATION = "TEST_APPLICATION"
)


type Configuration struct {
	// mode of the application
	Application_Mode string


	// node level information
	Node_Port string
	Node_IP string
	Node_ID string
	Node_M  string
	Jump_Spacing int
	FingerTableLength int


	// remote node connection  format `ip:port`
	Remote_Node_Address string


	// routing 
	RoutingUpdate_Delay int
	Successors_Table_Length int


	// application
	Application_Channel_Size int
	Application_Connecting_Port string


	// log configuration
	Log_Route_Table_Delay_Value int
	Log_Node_Channel_Size_Value int
	Log_Node_Channel_Delay_Value int
	
	Log_Network_Channel_Size_Value int
	Log_Network_Channel_Delay_Value int

	Log_Configuration_Channel_Size_Value int
	Log_Configuration_Channel_Delay_Value int

	Log_URL_Send_Route_Table string
	Log_URL_Send_Node_Info string
	Log_URL_Send_Network_Tool string
	Log_URL_Send_Configuration_Tool string


	// Network tool [Internal]
	Network_Tool_Channel_Size int


	// CDHT Network tool [External]
	CDHT_Ping_Tool_Listening_Port string
	CDHT_API_Communication_Delay int
	CDHT_Command_Channel_Size int
	CDHT_URL_Get_Command_From_Server string
	CDHT_URL_Send_Command_Result string
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