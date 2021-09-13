package Config

func (config *Configuration) CopyConfigData(RemoteConfig Configuration){
	if config.ConfigID > RemoteConfig.ConfigID {
		return
	}
	
	// node
	config.ConfigID = RemoteConfig.ConfigID
	config.Node_M  =  RemoteConfig.Node_M
	config.Jump_Spacing = RemoteConfig.Jump_Spacing
	config.Finger_Table_Length = RemoteConfig.Finger_Table_Length
	config.Routing_Update_Delay = RemoteConfig.Routing_Update_Delay
	config.Successors_Table_Length = RemoteConfig.Successors_Table_Length

	// application
	config.Application_Channel_Size = RemoteConfig.Application_Channel_Size

	// log configuration
	config.Log_Route_Table_Delay_Value = RemoteConfig.Log_Route_Table_Delay_Value
	config.Log_Node_Channel_Size_Value = RemoteConfig.Log_Node_Channel_Size_Value
	config.Log_Node_Channel_Delay_Value = RemoteConfig.Log_Node_Channel_Delay_Value
	config.Log_Network_Channel_Size_Value = RemoteConfig.Log_Network_Channel_Size_Value
	config.Log_Network_Channel_Delay_Value = RemoteConfig.Log_Network_Channel_Delay_Value
	config.Log_Configuration_Channel_Size_Value = RemoteConfig.Log_Configuration_Channel_Size_Value
	config.Log_Configuration_Channel_Delay_Value = RemoteConfig.Log_Configuration_Channel_Delay_Value
	config.Log_URL_Send_Route_Table = RemoteConfig.Log_URL_Send_Route_Table
	config.Log_URL_Send_Node_Info = RemoteConfig.Log_URL_Send_Node_Info
	config.Log_URL_Send_Network_Tool = RemoteConfig.Log_URL_Send_Network_Tool
	config.Log_URL_Send_Configuration_Tool = RemoteConfig.Log_URL_Send_Configuration_Tool

	// Network tool [Internal]
	config.Network_Tool_Channel_Size = RemoteConfig.Network_Tool_Channel_Size

	// CDHT Network tool [External]
	config.CDHT_Ping_Tool_Listening_Port = RemoteConfig.CDHT_Ping_Tool_Listening_Port
	config.CDHT_API_Communication_Delay = RemoteConfig.CDHT_API_Communication_Delay
	config.CDHT_Command_Channel_Size = RemoteConfig.CDHT_Command_Channel_Size
	config.CDHT_URL_Get_Command_From_Server = RemoteConfig.CDHT_URL_Get_Command_From_Server
	config.CDHT_URL_Send_Command_Result = RemoteConfig.CDHT_URL_Send_Command_Result

	// TEST Application
	config.TEST_APP_Net_Channel_Size = RemoteConfig.TEST_APP_Net_Channel_Size
	config.TEST_APP_AppName = RemoteConfig.TEST_APP_AppName
	config.TEST_APP_Packet_Delay = RemoteConfig.TEST_APP_Packet_Delay 
	config.TEST_APP_RUNNING_TIME = RemoteConfig.TEST_APP_RUNNING_TIME 
}