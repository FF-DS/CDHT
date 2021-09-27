package Config

import (
	"time"
	"fmt"
)


func (config *Configuration) DefaultConfig() *Configuration {
	return &Configuration{
		Node_M  : "160", 
		Jump_Spacing : 2, 
		Routing_Update_Delay : 1,
		Successors_Table_Length : 5,

		Application_Channel_Size : 100000,

		Log_Route_Table_Delay_Value : 15, 
		Log_Node_Channel_Size_Value : 100, 
		Log_Node_Channel_Delay_Value : 15, 
		Log_Network_Channel_Size_Value : 100, 
		Log_Network_Channel_Delay_Value : 15, 
		Log_Configuration_Channel_Size_Value : 100, 
		Log_Configuration_Channel_Delay_Value : 15, 
		Log_URL_Send_Route_Table : "https://cdht-monitoring-api.herokuapp.com/logs", 
		Log_URL_Send_Node_Info : "https://cdht-monitoring-api.herokuapp.com/logs", 
		Log_URL_Send_Network_Tool : "https://cdht-monitoring-api.herokuapp.com/logs", 
		Log_URL_Send_Configuration_Tool : "https://cdht-monitoring-api.herokuapp.com/logs", 

		Network_Tool_Channel_Size : 100000,

		CDHT_API_Communication_Delay : 5,
		CDHT_Command_Channel_Size : 100000,
		CDHT_URL_Get_Command_From_Server : "https://cdht-monitoring-api.herokuapp.com/network-tools",
		CDHT_URL_Send_Command_Result : "https://cdht-monitoring-api.herokuapp.com/network-tools",

		TEST_APP_Net_Channel_Size : 100000,
		TEST_APP_AppName : "TEST_APP",
		TEST_APP_Packet_Delay :  500, // millisecond
		TEST_APP_RUNNING_TIME : 300, // minute 

		CONFIGURATION_DOWNLOAD_DELAY : 5,
		CONFIGURATION_DOWNLOAD_URL : "https://cdht-monitoring-api.herokuapp.com/configuration/current",

		REPLICATION_COUNT: 0,
	}
}



func (config *Configuration) CopyConfiguration(NewConfig *Configuration) {
	defaultConfig := config.DefaultConfig()

	config.Node_M = selectStr( NewConfig.Node_M, config.Node_M, defaultConfig.Node_M)
	config.Jump_Spacing = selectInt( NewConfig.Jump_Spacing, config.Jump_Spacing, defaultConfig.Jump_Spacing) 
	config.Routing_Update_Delay = selectTime( NewConfig.Routing_Update_Delay, config.Routing_Update_Delay, defaultConfig.Routing_Update_Delay)
	config.Successors_Table_Length = selectInt( NewConfig.Successors_Table_Length, config.Successors_Table_Length,  defaultConfig.Successors_Table_Length)

	config.Application_Channel_Size = selectInt( NewConfig.Application_Channel_Size, config.Application_Channel_Size, defaultConfig.Application_Channel_Size)

	config.Log_Route_Table_Delay_Value = selectTime( NewConfig.Log_Route_Table_Delay_Value, config.Log_Route_Table_Delay_Value, defaultConfig.Log_Route_Table_Delay_Value) 
	config.Log_Node_Channel_Size_Value = selectInt( NewConfig.Log_Node_Channel_Size_Value, config.Log_Node_Channel_Size_Value, defaultConfig.Log_Node_Channel_Size_Value) 
	config.Log_Node_Channel_Delay_Value = selectTime( NewConfig.Log_Node_Channel_Delay_Value, config.Log_Node_Channel_Delay_Value, defaultConfig.Log_Node_Channel_Delay_Value)
	config.Log_Network_Channel_Size_Value = selectInt( NewConfig.Log_Network_Channel_Size_Value, config.Log_Network_Channel_Size_Value, defaultConfig.Log_Network_Channel_Size_Value)
	config.Log_Network_Channel_Delay_Value = selectTime( NewConfig.Log_Network_Channel_Delay_Value, config.Log_Network_Channel_Delay_Value, defaultConfig.Log_Network_Channel_Delay_Value)
	config.Log_Configuration_Channel_Size_Value = selectInt( NewConfig.Log_Configuration_Channel_Size_Value, config.Log_Configuration_Channel_Size_Value, defaultConfig.Log_Configuration_Channel_Size_Value)
	config.Log_Configuration_Channel_Delay_Value = selectTime( NewConfig.Log_Configuration_Channel_Delay_Value, config.Log_Configuration_Channel_Delay_Value, defaultConfig.Log_Configuration_Channel_Delay_Value)

	config.Log_URL_Send_Route_Table = selectURL( NewConfig.Log_URL_Send_Route_Table, config.Log_URL_Send_Route_Table, defaultConfig.Log_URL_Send_Route_Table)
	config.Log_URL_Send_Node_Info = selectURL( NewConfig.Log_URL_Send_Node_Info, config.Log_URL_Send_Node_Info, defaultConfig.Log_URL_Send_Node_Info) 
	config.Log_URL_Send_Network_Tool = selectURL( NewConfig.Log_URL_Send_Network_Tool, config.Log_URL_Send_Network_Tool, defaultConfig.Log_URL_Send_Network_Tool) 
	config.Log_URL_Send_Configuration_Tool = selectURL( NewConfig.Log_URL_Send_Configuration_Tool, config.Log_URL_Send_Configuration_Tool, defaultConfig.Log_URL_Send_Configuration_Tool) 

	config.Network_Tool_Channel_Size = selectInt( NewConfig.Network_Tool_Channel_Size, config.Network_Tool_Channel_Size, defaultConfig.Network_Tool_Channel_Size)

	config.CDHT_API_Communication_Delay = selectTime( NewConfig.CDHT_API_Communication_Delay, config.CDHT_API_Communication_Delay, defaultConfig.CDHT_API_Communication_Delay)
	config.CDHT_Command_Channel_Size = selectInt( NewConfig.CDHT_Command_Channel_Size, config.CDHT_Command_Channel_Size, defaultConfig.CDHT_Command_Channel_Size)
	config.CDHT_URL_Get_Command_From_Server = selectURL( NewConfig.CDHT_URL_Get_Command_From_Server, config.CDHT_URL_Get_Command_From_Server, defaultConfig.CDHT_URL_Get_Command_From_Server)
	config.CDHT_URL_Send_Command_Result = selectURL( NewConfig.CDHT_URL_Send_Command_Result, config.CDHT_URL_Send_Command_Result, defaultConfig.CDHT_URL_Send_Command_Result)

	config.TEST_APP_Net_Channel_Size = selectInt( NewConfig.TEST_APP_Net_Channel_Size, config.TEST_APP_Net_Channel_Size, defaultConfig.TEST_APP_Net_Channel_Size)
	config.TEST_APP_AppName = selectStr( NewConfig.TEST_APP_AppName, config.TEST_APP_AppName, defaultConfig.TEST_APP_AppName)
	config.TEST_APP_Packet_Delay = selectTime( NewConfig.TEST_APP_Packet_Delay, config.TEST_APP_Packet_Delay, defaultConfig.TEST_APP_Packet_Delay)
	config.TEST_APP_RUNNING_TIME = selectTime( NewConfig.TEST_APP_RUNNING_TIME, config.TEST_APP_RUNNING_TIME, defaultConfig.TEST_APP_RUNNING_TIME)

	config.CONFIGURATION_DOWNLOAD_DELAY = selectTime( NewConfig.CONFIGURATION_DOWNLOAD_DELAY, config.CONFIGURATION_DOWNLOAD_DELAY, defaultConfig.CONFIGURATION_DOWNLOAD_DELAY)
	config.CONFIGURATION_DOWNLOAD_URL  = selectURL( NewConfig.CONFIGURATION_DOWNLOAD_URL, config.CONFIGURATION_DOWNLOAD_URL, defaultConfig.CONFIGURATION_DOWNLOAD_URL)

	config.REPLICATION_COUNT = NewConfig.REPLICATION_COUNT
}


func (config *Configuration) PrintConfiguration() {
    fmt.Println("---------------------------[Configuration]---------------------------\n")
	fmt.Println("      [+]: ConfigID:  ", config.ConfigID )
	fmt.Println("      [+]: Application_Mode:  ", config.Application_Mode )
	
	fmt.Println()
	fmt.Println("      [+]: RUN_TCP_TEST_APPLICATION:  ", config.RUN_NETWORK_TEST_APPLICATION )
	fmt.Println("      [+]: RUN_UDP_TEST_APPLICATION:  ", config.RUN_UDP_TEST_APPLICATION )
	fmt.Println("      [+]: RUN_NETWORK_TEST_APPLICATION:  ", config.RUN_NETWORK_TEST_APPLICATION )

	fmt.Println()
	fmt.Println("      [+]: Node_Port:  ", config.Node_Port )
	fmt.Println("      [+]: Node_IP:  ", config.Node_IP )
	fmt.Println("      [+]: Node_ID:  ", config.Node_ID )

	fmt.Println()
	fmt.Println("      [+]: Remote Node Address:  ", config.Remote_Node_Address )

	fmt.Println()
	fmt.Println("      [+]: Node_M  :  ",  config.Node_M )
	fmt.Println("      [+]: Jump_Spacing :  ", config.Jump_Spacing )
	fmt.Println("      [+]: Finger_Table_Length :  ", config.Finger_Table_Length )
	fmt.Println("      [+]: Routing_Update_Delay :  ", config.Routing_Update_Delay )
	fmt.Println("      [+]: Successors_Table_Length :  ", config.Successors_Table_Length )

	fmt.Println()
	fmt.Println("      [+]: Application_Connecting_Port :  ", config.Application_Connecting_Port )
	fmt.Println("      [+]: Application_Channel_Size :  ", config.Application_Channel_Size )

	fmt.Println()
	fmt.Println("      [+]: Log_Route_Table_Delay_Value :  ", config.Log_Route_Table_Delay_Value )
	fmt.Println("      [+]: Log_Node_Channel_Size_Value :  ", config.Log_Node_Channel_Size_Value )
	fmt.Println("      [+]: Log_Node_Channel_Delay_Value :  ", config.Log_Node_Channel_Delay_Value )
	fmt.Println("      [+]: Log_Network_Channel_Size_Value :  ", config.Log_Network_Channel_Size_Value )
	fmt.Println("      [+]: Log_Network_Channel_Delay_Value :  ", config.Log_Network_Channel_Delay_Value )
	fmt.Println("      [+]: Log_Configuration_Channel_Size_Value :  ", config.Log_Configuration_Channel_Size_Value )
	fmt.Println("      [+]: Log_Configuration_Channel_Delay_Value :  ", config.Log_Configuration_Channel_Delay_Value )
	fmt.Println("      [+]: Log_URL_Send_Route_Table :  ", config.Log_URL_Send_Route_Table )
	fmt.Println("      [+]: Log_URL_Send_Node_Info :  ", config.Log_URL_Send_Node_Info )
	fmt.Println("      [+]: Log_URL_Send_Network_Tool :  ", config.Log_URL_Send_Network_Tool )
	fmt.Println("      [+]: Log_URL_Send_Configuration_Tool :  ", config.Log_URL_Send_Configuration_Tool )

	fmt.Println()
	fmt.Println("      [+]: Network_Tool_Channel_Size :  ", config.Network_Tool_Channel_Size )
	
	fmt.Println()
	fmt.Println("      [+]: CDHT_Ping_Tool_Listening_Port :  ", config.CDHT_Ping_Tool_Listening_Port )
	fmt.Println("      [+]: CDHT_API_Communication_Delay :  ", config.CDHT_API_Communication_Delay )
	fmt.Println("      [+]: CDHT_Command_Channel_Size :  ", config.CDHT_Command_Channel_Size )
	fmt.Println("      [+]: CDHT_URL_Get_Command_From_Server :  ", config.CDHT_URL_Get_Command_From_Server )
	fmt.Println("      [+]: CDHT_URL_Send_Command_Result :  ", config.CDHT_URL_Send_Command_Result )


	fmt.Println()
	fmt.Println("      [+]: TEST_APP_IPAddress :  ",  config.TEST_APP_IPAddress )
	fmt.Println("      [+]: TEST_APP_Port :  ",  config.TEST_APP_Port ) 
	fmt.Println("      [+]: TEST_APP_UDP_Listener_Port :  ",  config.TEST_APP_UDP_Listener_Port )
	fmt.Println("      [+]: TEST_APP_AppName :  ",  config.TEST_APP_AppName )
	fmt.Println("      [+]: TEST_APP_Packet_Delay :  ",  config.TEST_APP_Packet_Delay ) 
	fmt.Println("      [+]: TEST_APP_Net_Channel_Size :  ",  config.TEST_APP_Net_Channel_Size )
	fmt.Println("      [+]: TEST_APP_RUNNING_TIME :  ",  config.TEST_APP_RUNNING_TIME ) 

	fmt.Println()
	fmt.Println("      [+]: CONFIGURATION_DOWNLOAD_DELAY :  ",  config.CONFIGURATION_DOWNLOAD_DELAY )
	fmt.Println("      [+]: CONFIGURATION_DOWNLOAD_URL :  ",  config.CONFIGURATION_DOWNLOAD_URL )
	fmt.Println()
	fmt.Println("      [+]: REPLICATION_COUNT :  ",  config.REPLICATION_COUNT )
	fmt.Println()
    fmt.Println("---------------------------------------------------------------------\n")
}




// ## -------------------  helper ------------------ ##

func selectStr(new, curr, defaultD string) string {
	if new != "" {
		return new
	}
	if curr != "" {
		return curr
	}
	return defaultD
}


func selectURL(new, curr, defaultD string) string {
	if new != "" {
		return new
	}
	if curr != "" {
		return curr
	}
	return defaultD
}


func selectInt(new, curr, defaultD int) int {
	if new != 0 {
		return new
	}
	if curr != 0 {
		return curr
	}
	return defaultD
}

func selectTime(new, curr, defaultD time.Duration) time.Duration {
	if new != 0 {
		return new
	}
	if curr != 0 {
		return curr
	}
	return defaultD
}