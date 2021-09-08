package NetworkTools

import (
	"cdht/ReportModule"
	"cdht/RoutingModule"
	"cdht/Util"
)


type NetworkTool struct {
	RoutingTable *RoutingModule.RoutingTable
	node *RoutingModule.Node
    Logger *ReportModule.Logger
	NetworkToolPackets chan Util.RequestObject 
}

func (netTool *NetworkTool) Init(size int){		
	netTool.NetworkToolPackets = make(chan Util.RequestObject, size) 
}

func (netTool *NetworkTool) RunTools(){
	netTool.node = netTool.RoutingTable.NodeInfo()
	go netTool.runNetworkTools()
}


func (netTool *NetworkTool) runNetworkTools(){
	hopTools := HopCountTool{ routingTable: netTool.RoutingTable , networkTool: netTool }
	lookupTools := LookUpTool{ routingTable: netTool.RoutingTable , networkTool: netTool }
	pingTools := PingTool{ routingTable: netTool.RoutingTable , networkTool: netTool }

	for {
		req :=  <- netTool.NetworkToolPackets

		if req.AppName == "HopCountTool" {
			hopTools.RunTool( req )
		}else if req.AppName == "LookUpTool" { 
			lookupTools.RunTool( req )
		}else if req.AppName == "PingTool" { 
			pingTools.RunTool( req )
		}
	}
}


func (netTool *NetworkTool) logNetworkReport(logType string, location string, status string, logData map[string]string ) {
    netLog := ReportModule.Log {
        Type: logType,
        OperationStatus: status,
        LogLocation: location,
        NodeId: netTool.node.Node_id,
        NodeAddress: netTool.node.IP_address + ":" + netTool.node.Port,
        LogBody: logData,
    }

    netTool.Logger.NetworkToolLog(netLog)
}