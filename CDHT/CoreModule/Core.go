package CoreModule


import (
	"cdht/Applications/TestApplications"
    "cdht/Applications/CDHTNetworkTools"
    "cdht/RoutingModule"
    "cdht/ReportModule"
    "cdht/NetworkTools"
    "cdht/Config"
    "cdht/API"
	"sync"
	"time"
	"fmt"
)


type Core struct {
	ConfigMngr				*Config.Config
	Config   		        *Config.Configuration
	
	ApiCommunication   	    *API.ApiCommunication
	LogManager  		    *ReportModule.Logger
	InternalNetworkTools    *NetworkTools.NetworkTool
	RoutingTableInfo   	    *RoutingModule.RoutingTable
	CdhtNetworkTools   	    *CDHTNetworkTools.CDHTNetworkTool
	CdhtTestApplication     *TestApplications.TestApplication
	RemoteNode              *RoutingModule.NodeRPC
}




// # --------------- [INITALIZATION!!] --------------- #

func (core *Core) InitalizeApiCommunication(){
	apiComm := API.ApiCommunication{
        ChannelSize:  core.Config.Application_Channel_Size,
        PORT:  core.Config.Application_Connecting_Port,
    }
    apiComm.Init()
	
	core.ApiCommunication = &apiComm
}


func (core *Core) InitalizeLogManager(){
	logMngr := ReportModule.Logger{}
	logMngr.Configure( core.Config.GetLogConfiguration() )
    logMngr.Init()
	
	core.LogManager = &logMngr
	core.ConfigMngr.Logger = &logMngr
}


func (core *Core) InitalizeNetworkTools(){
	netTools := NetworkTools.NetworkTool{ Logger: core.LogManager }
    netTools.Init( core.Config.Network_Tool_Channel_Size )

	core.InternalNetworkTools = &netTools
}


func (core *Core) ConnectWithNode(){
	remoteNode := RoutingModule.NodeRPC{ Node_address : core.Config.Remote_Node_Address }
	remoteNode.Connect()
	core.RemoteNode = &remoteNode
}



func (core *Core) InitalizeRoutingTable(){
	routingTableInfo := RoutingModule.RoutingTable{ 
		Applications: core.ApiCommunication.Application, 
		Logger: core.LogManager, 
		NetworkTools: core.InternalNetworkTools.NetworkToolPackets, 
		RoutingUpdateDelay: core.Config.Routing_Update_Delay, 
		SuccessorsTableLength: core.Config.Successors_Table_Length, 
		NodePort: core.Config.Node_Port,
		IP_address: core.Config.Node_IP,
		M: core.Config.GetNodeM(),
		Node_id: core.Config.GetNodeID(),
		FingerTableLength: core.Config.Finger_Table_Length,
		JumpSpacing: core.Config.Jump_Spacing,
	}
	
	core.RoutingTableInfo = &routingTableInfo

	// [CREATE RING] or [JOIN RING]
	if core.Config.Application_Mode == Config.MODE_CREATE_RING {
		core.RoutingTableInfo.CreateRing()    

	}else if core.Config.Application_Mode == Config.MODE_JOIN_RING {
		core.ConnectWithNode()
		core.RoutingTableInfo.RunNode( core.RemoteNode )    

	}else if core.Config.Application_Mode == Config.MODE_REPLICA_NODE {
		core.ConnectWithNode()
		core.RoutingTableInfo.InitReplicaRoutingTable( core.RemoteNode )    
	}
}


func (core *Core) InitalizeCDHTNetworkTools(){
	if core.RoutingTableInfo.NodeInfo().IP_address == "" {
		time.Sleep(time.Second)
		core.InitalizeCDHTNetworkTools()
		return 
	}
	cdhtTools := CDHTNetworkTools.CDHTNetworkTool{
        AppServerIP: core.RoutingTableInfo.NodeInfo().IP_address,
        AppServerPort: core.Config.Application_Connecting_Port,
        PingToolListeningPort: core.Config.CDHT_Ping_Tool_Listening_Port,
        ReadCommandDelay: core.Config.CDHT_API_Communication_Delay,
        ChannelSize: core.Config.CDHT_Command_Channel_Size,
        NodeId: core.RoutingTableInfo.NodeInfo().Node_id,
        NodeAddress: core.RoutingTableInfo.NodeInfo().IP_address + ":" + core.Config.Node_Port,
        URLGetCommandFromServer: core.Config.CDHT_URL_Get_Command_From_Server,
        URLSendCommandResult: core.Config.CDHT_URL_Send_Command_Result,
    }
    cdhtTools.Init()
	core.CdhtNetworkTools = &cdhtTools
}



func (core *Core) InitalizeTestApplicationTool(){
	testApp := TestApplications.TestApplication{
		IPAddress: core.Config.TEST_APP_IPAddress,
		Port: core.Config.TEST_APP_Port,
		UDPListenerPort: core.Config.TEST_APP_UDP_Listener_Port,
		NetChannelSize: core.Config.TEST_APP_Net_Channel_Size,
		AppName: core.Config.TEST_APP_AppName,
		PacketDelay: core.Config.TEST_APP_Packet_Delay,
	}

	testApp.Init()
	core.CdhtTestApplication = &testApp
}



func (core *Core) InitalizeConfiguration() {
	configMngr := Config.Config{}
    config := configMngr.LoadConfig()
	
	core.ConfigMngr = &configMngr
	core.Config = config
	
    go configMngr.DownloadConfiguration()
	go core.UpdateApplicationConfiguration()
}



// # --------------- [RUNNING!!] --------------- #

func (core *Core) RunNode() {
	// Initalizing apps
	core.InitalizeApiCommunication()
	core.InitalizeLogManager()
	core.InitalizeNetworkTools()


	// init routing table
	core.InitalizeRoutingTable()
    progressBar("[Init]: + Initalizing routing tables.... ", 100, nil)


	// init api communication
	core.ApiCommunication.NodeRoutingTable = core.RoutingTableInfo
	core.ApiCommunication.StartAppServer()


	// init [internal network tool]
	core.InternalNetworkTools.RoutingTable = core.RoutingTableInfo
	core.InternalNetworkTools.RunTools()

	core.RunApplications()
    progressBar("[Init]: + Initalizing application manager....", 100, nil)

	// NODE INFO
	fmt.Println("\n")
	fmt.Println(logo)
	core.RoutingTableInfo.PrintCurrentNodeInfo() 
	fmt.Println("\n")
}



func (core *Core) RunApplications() {
	if (core.Config.RUN_TCP_TEST_APPLICATION){
		core.InitalizeTestApplicationTool()
		core.CdhtTestApplication.TestAppTCP()
	}

	if (core.Config.RUN_UDP_TEST_APPLICATION){
		core.InitalizeTestApplicationTool()
		core.CdhtTestApplication.TestAppUDP()
	}

	if (core.Config.RUN_NETWORK_TEST_APPLICATION){
		core.InitalizeCDHTNetworkTools()
	}
}






// # -------------------  [START]  ------------------ #
// #==================================================#
func (core *Core) START() {
	var wg sync.WaitGroup
	
	wg.Add(1)
	go progressBar("[Init]: + Loading configuration.... ", 300, &wg)
	wg.Wait()
	core.InitalizeConfiguration()

	if !core.Config.RUN_TCP_TEST_APPLICATION && !core.Config.RUN_UDP_TEST_APPLICATION {
		core.RunNode()
		ui := TerminalUI{ CoreLink : core }
    	ui.UserUI()
	}
	core.RunApplications()

	if core.Config.RUN_TCP_TEST_APPLICATION || core.Config.RUN_UDP_TEST_APPLICATION { 
        time.Sleep(time.Minute * core.Config.TEST_APP_RUNNING_TIME)
    }
}




var logo  string = `
                                               ▄████▄  ▓█████▄  ██░ ██ ▄▄▄█████▓
                                               ▒██▀ ▀█  ▒██▀ ██▌▓██░ ██▒▓  ██▒ ▓▒
                                               ▒▓█    ▄ ░██   █▌▒██▀▀██░▒ ▓██░ ▒░
                                               ▒▓▓▄ ▄██▒░▓█▄   ▌░▓█ ░██ ░ ▓██▓ ░ 
                                               ▒ ▓███▀ ░░▒████▓ ░▓█▒░██▓  ▒██▒ ░ 
                                               ░ ░▒ ▒  ░ ▒▒▓  ▒  ▒ ░░▒░▒  ▒ ░░   
                                                 ░  ▒    ░ ▒  ▒  ▒ ░▒░ ░    ░    
                                               ░         ░ ░  ░  ░  ░░ ░  ░      
                                               ░ ░         ░     ░  ░  ░         
                                               ░         ░
`