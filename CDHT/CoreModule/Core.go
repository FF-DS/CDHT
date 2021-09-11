package CoreModule


import (
    "cdht/Applications/TestApplications"
    "cdht/Applications/CDHTNetworkTools"
    "cdht/RoutingModule"
    "cdht/ReportModule"
    "cdht/NetworkTools"
    "cdht/API"
)


type Core struct {
	config   		        *Configuration
	
	apiCommunication   	    *API.ApiCommunication
	logManager  		    *ReportModule.Logge
	internalNetworkTools    *NetworkTools.NetworkTool
	routingTableInfo   	    *RoutingModule.RoutingTable
	cdhtNetworkTools   	    *CDHTNetworkTools.CDHTNetworkTool
}



func (core *Core) InitalizeApiCommunication(){
	apiComm := API.ApiCommunication{
        ChannelSize:  core.config.Application_Channel_Size,
        PORT:  core.config.Application_Connecting_Port,
    }
    apiComm.Init()
	
	core.apiCommunication = &apiComm
}


func (core *Core) InitalizeLogManager(){
	logMngr := ReportModule.Logger{}
	logMngr.Configure( core.config.GetLogConfiguration() )
    logMngr.Init()
	
	core.logManager = &logMngr
}


func (core *Core) InitalizeNetworkTools(){
	netTools := NetworkTools.NetworkTool{ Logger: core.logManager }
    netTools.Init( core.config.Network_Tool_Channel_Size )

	core.internalNetworkTools = &netTools
}


func (core *Core) InitalizeRoutingTable(){
	routingTableInfo := RoutingModule.RoutingTable{ 
		Applications: core.apiCommunication.Application, 
		Logger: core.logManager, 
		NetworkTools: core.internalNetworkTools.NetworkToolPackets, 
		RoutingUpdateDelay: 1, 
		SuccessorsTableLength: 5, 
		RemoteNodeAddr: core.config.Remote_Node_Address,
	}
	
	core.routingTableInfo = &routingTableInfo

	// [CREATE RING] or [JOIN RING]
	if core.config.Application_Mode == MODE_CREATE_RING {
		core.routingTableInfo.CreateRing()    
	}else if core.config.Application_Mode == MODE_JOIN_RING {
		core.routingTableInfo.RunNode()    
	}
}


func (core *Core) InitalizeCDHTNetworkTools(){
	
}


func (core *Core) InitalizeTestApplicationTool(){
	
}