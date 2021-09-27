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
	ConfigID int `json:"configuration_id" bson:"configuration_id"`
	
	// mode of the application
	Application_Mode string `json:"application_mode" bson:"application_mode"`

	RUN_TCP_TEST_APPLICATION bool `json:"run_tcp_application" bson:"run_tcp_application"`
	RUN_UDP_TEST_APPLICATION bool `json:"run_ud_test_application" bson:"run_ud_test_application"`
	RUN_NETWORK_TEST_APPLICATION bool `json:"run_network_test_application" bson:"run_network_test_application"`

	// node level & routing information
	Node_Port string `json:"node_port" bson:"node_port"`
	Node_IP string `json:"node_ip" bson:"node_ip"`
	Node_ID string `json:"node_id" bson:"node_id"`
	Node_M  string `json:"node_m" bson:"node_m"`
	Jump_Spacing int `json:"jump_spacing" bson:"jump_spacing"`
	Finger_Table_Length int `json:"finger_table_length" bson:"finger_table_length"`
	Routing_Update_Delay time.Duration `json:"routing_update_delay" bson:"routing_update_delay"`
	Successors_Table_Length int `json:"succesor_table_length" bson:"succesor_table_length"`

	// remote node connection  format `ip:port`
	Remote_Node_Address string `json:"remote_node_address" bson:"remote_node_address"`


	// application
	Application_Channel_Size int `json:"application_channel_size" bson:"application_channel_size"`
	Application_Connecting_Port string `json:"application_connecting_node" bson:"application_connecting_node"`


	// log configuration
	Log_Route_Table_Delay_Value time.Duration `json:"log_route_table_delay_value" bson:"log_route_table_delay_value"`
	Log_Node_Channel_Size_Value int `json:"log_node_channel_size_value" bson:"log_node_channel_size_value"`
	Log_Node_Channel_Delay_Value time.Duration `json:"log_node_channel_delay_value" bson:"log_node_channel_delay_value"`
	
	Log_Network_Channel_Size_Value int `json:"log_network_channel_size_value" bson:"log_network_channel_size_value"`
	Log_Network_Channel_Delay_Value time.Duration `json:"log_network_channel_delay_value" bson:"log_network_channel_delay_value"`

	Log_Configuration_Channel_Size_Value int `json:"log_configuration_channel_size_value" bson:"log_configuration_channel_size_value"`
	Log_Configuration_Channel_Delay_Value time.Duration `json:"log_configuration_channel_delay_value" bson:"log_configuration_channel_delay_value"`

	Log_URL_Send_Route_Table string `json:"log_url_send_route_table" bson:"log_url_send_route_table"`
	Log_URL_Send_Node_Info string `json:"log_url_node_info" bson:"log_url_node_info"`
	Log_URL_Send_Network_Tool string `json:"log_url_send_network_tool" bson:"log_url_send_network_tool"`
	Log_URL_Send_Configuration_Tool string `json:"log_url_send_configuration_tool" bson:"log_url_send_configuration_tool"`


	// Network tool [Internal]
	Network_Tool_Channel_Size int `json:"network_tool_channel_size" bson:"network_tool_channel_size"`


	// CDHT Network tool [External]
	CDHT_Ping_Tool_Listening_Port string `json:"cdht_ping_tool_listening_port" bson:"cdht_ping_tool_listening_port"`
	CDHT_API_Communication_Delay time.Duration `json:"cdht_api_communication_delay" bson:"cdht_api_communication_delay"`
	CDHT_Command_Channel_Size int `json:"cdht_command_channel_size" bson:"cdht_command_channel_size"`
	CDHT_URL_Get_Command_From_Server string `json:"cdht_url_get_command_from_server" bson:"cdht_url_get_command_from_server"`
	CDHT_URL_Send_Command_Result string `json:"cdht_url_send_command_result" bson:"cdht_url_send_command_result"`


	// TEST Application
	TEST_APP_IPAddress string `json:"test_app_ipaddress" bson:"test_app_ipaddress"`
	TEST_APP_Port string `json:"test_app_port" bson:"test_app_port"`
	TEST_APP_UDP_Listener_Port string `json:"test_app_udp_listner_port" bson:"test_app_udp_listner_port"`
	TEST_APP_Net_Channel_Size int `json:"test_app_net_channel_size" bson:"test_app_net_channel_size"`
	TEST_APP_AppName string `json:"test_app_appname" bson:"test_app_appname"`
	TEST_APP_Packet_Delay  time.Duration  `json:"test_app_packet_delay" bson:"test_app_packet_delay"`
	TEST_APP_RUNNING_TIME  time.Duration `json:"test_app_running_time" bson:"test_app_running_time"`

	CONFIGURATION_DOWNLOAD_DELAY  time.Duration `json:"configuration_download_delay" bson:"configuration_download_delay"`
	CONFIGURATION_DOWNLOAD_URL string `json:"configuration_download_url" bson:"configuration_download_url"`

	REPLICATION_COUNT int `json:"replication_count" bson:"replication_count"`
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