package RoutingModule

import (
	"time"
    "math/big"
	"cdht/Util"
	"cdht/ReportModule"
)


type RoutingTable struct {
	RemoteNodeAddr string
	NodePort string
	IP_address string

	JumpSpacing int
	FingerTableLength int
    Node_id *big.Int
    M *big.Int

    Applications map[string](chan Util.RequestObject)
    Logger *ReportModule.Logger
	node *Node
}



// # ------------------------ Main Node INIT ----------------------- #

func (routingTable *RoutingTable) CreateRing() {
    routingTable.node = &Node{
        Port : routingTable.NodePort,
        JumpSpacing : routingTable.JumpSpacing,
        FingerTableLength : routingTable.FingerTableLength,
        Applications : routingTable.Applications,
        Logger: routingTable.Logger,
        Node_id : routingTable.Node_id,
        IP_address : routingTable.IP_address,
        M : routingTable.M,
    }

    routingTable.node.createRing()
    // [PRINT]:WILL BE REMOVED
    routingTable.node.currentNodeInfo()
    
    go routingTable.runStablization()
}



func (routingTable *RoutingTable) RunNode() {

    remoteNode := NodeRPC{ Node_address : routingTable.RemoteNodeAddr }
	remoteNode.Connect()

    // [PRINT]:WILL BE REMOVED
    remoteNode.PrintNodeInfo()

    routingTable.node = &Node{
        Port : routingTable.NodePort,
        JumpSpacing : routingTable.JumpSpacing,
        FingerTableLength : routingTable.FingerTableLength,
        Applications : routingTable.Applications,
        Logger: routingTable.Logger,
        Node_id : routingTable.Node_id,
        IP_address : routingTable.IP_address,
        M : routingTable.M,
    }

    routingTable.node.join( &remoteNode)
    // [PRINT]:WILL BE REMOVED
    routingTable.node.currentNodeInfo()
    
    go routingTable.runStablization()
}


func (routingTable *RoutingTable) runStablization() {
	for {
		time.Sleep(time.Second)
		routingTable.node.checkPredecessor()
		routingTable.node.checkSeccessors()
		routingTable.node.stablize()
		routingTable.node.fixFinger()
        // logging
        routingTable.node.LOGRoutingTableReport()
	}
}
// # ------------------------ [END] Main Node INIT ----------------------- #







// # --------------------------- Routing Functionalities -------------------------- #

// [CORE-MODULE]
func (routingTable *RoutingTable) ForwardPacket(req Util.RequestObject) Util.ResponseObject {
    successor := NodeRPC{ NodeTraversalLogs: []NodeRPC{} }
    start := time.Now()
    err := routingTable.node.FindSuccessor( req.ReceiverNodeId, &successor)

    if err == nil && checkNode(&successor) != nil{
        rtt := time.Since(start)
        _, resp := successor.ResolvePacket( req )
        latency := time.Since(start)
        // ---- logging ----- //
        routingTable.node.LOGNodeReport( 
            ReportModule.LOG_TYPE_NODE_INFORMATION,
            ReportModule.LOG_LOCATION_TYPE_LEAVING,
            ReportModule.LOG_OPERATION_STATUS_SUCCESS,
            map[string]string{
                "rtt" : rtt.String(),
                "srt" : (rtt/2).String(),
                "latency" : latency.String(),
                "app_name" : req.AppName,
            },
        )
        // ---- logging ----- //
        return resp
    }

    // ---- logging ----- //
    routingTable.node.LOGNodeReport( 
        ReportModule.LOG_TYPE_NODE_INFORMATION,
        ReportModule.LOG_LOCATION_TYPE_LEAVING,
        ReportModule.LOG_OPERATION_STATUS_FAILED,
        map[string]string{
            "rtt" : time.Since(start).String(),
            "srt" : (time.Since(start)/2).String(),
            "latency" : time.Since(start).String(),
            "app_name" : req.AppName,
        },
    )
    // ---- logging ----- //

    resp := req.GetResponseObject()
    resp.ResponseStatus = Util.PACKET_STATUS_FAILED
    return resp
}



// [CORE-MODULE]
func (routingTable *RoutingTable) LookUp(nodeId *big.Int) NodeRPC {
    successor := NodeRPC{ NodeTraversalLogs: []NodeRPC{} }
    routingTable.node.FindSuccessor( nodeId, &successor)
    return successor
}


// [CORE-MODULE]
func (routingTable *RoutingTable) PrintRoutingInfo(){
    routingTable.node.currentSuccessorTableInfo()
	routingTable.node.currentFingerTableInfo()
    routingTable.node.currentSuccessorsInfo()
}
// # ------------------------ [END] Routing Functionalities ----------------------- #


