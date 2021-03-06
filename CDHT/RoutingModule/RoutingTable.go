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
    RoutingUpdateDelay time.Duration
    SuccessorsTableLength int

    Applications map[string](chan Util.RequestObject)
    NetworkTools chan Util.RequestObject

    Logger *ReportModule.Logger
	node *Node

	ReplicationCount int
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
        NetworkTools : routingTable.NetworkTools,
        SuccessorsTableLength : routingTable.SuccessorsTableLength,
        ReplicationCount : routingTable.ReplicationCount,
    }

    routingTable.node.createRing()
    
    go routingTable.runStablization()
}



func (routingTable *RoutingTable) RunNode(remoteNode *NodeRPC) {
    routingTable.node = &Node{
        Port : routingTable.NodePort,
        JumpSpacing : routingTable.JumpSpacing,
        FingerTableLength : routingTable.FingerTableLength,
        Applications : routingTable.Applications,
        Logger: routingTable.Logger,
        Node_id : routingTable.Node_id,
        IP_address : routingTable.IP_address,
        M : routingTable.M,
        NetworkTools : routingTable.NetworkTools,
        SuccessorsTableLength : routingTable.SuccessorsTableLength,
        ReplicationCount : routingTable.ReplicationCount,
    }

    routingTable.node.join( remoteNode )
    
    go routingTable.runStablization()
}


func (routingTable *RoutingTable) InitReplicaRoutingTable(remoteNode *NodeRPC) {
    routingTable.node = &Node{
        Port : routingTable.NodePort,
        Applications : routingTable.Applications,
        Logger: routingTable.Logger,
        IP_address : routingTable.IP_address,
        NetworkTools : routingTable.NetworkTools,
        SuccessorsTableLength : routingTable.SuccessorsTableLength,
    }

    routingTable.node.makeReplicaOf(remoteNode)

    go routingTable.runStablization()
}

// # ------------------------ [END] Main Node INIT ----------------------- #



// # -------------------------------- INIT ------------------------------- #

func (routingTable *RoutingTable) runStablization() {
	for {
		time.Sleep(time.Second * routingTable.RoutingUpdateDelay)

        if routingTable.node.NodeState == NODE_STATE_ACTIVE {
            routingTable.node.checkPredecessor()
            routingTable.node.checkSeccessors()
            routingTable.node.stablize()
            routingTable.node.fixFinger()
            // replica info
            routingTable.node.currentNodeReplicaInfo()

        }else{
            routingTable.node.updateReplicaInfo()
        }
        // logging
        routingTable.node.logRoutingTableReport()
	}
}
// # ----------------------------- [END] INIT ---------------------------- #







// # --------------------------- Routing Functionalities -------------------------- #

// [CORE-MODULE]
func (routingTable *RoutingTable) ForwardPacket(req Util.RequestObject) Util.RequestObject {
    successor := NodeRPC{ NodeTraversalLogs: []NodeRPC{} }
    start := time.Now()
    err := routingTable.node.FindSuccessor( req.ReceiverNodeId, &successor)

    if err == nil && checkNode(&successor) != nil{
        rtt := time.Since(start)
        _, resp := successor.ResolvePacket( req )
        latency := time.Since(start)
        // ---- logging ----- //
        routingTable.node.logNodeReport( 
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
    routingTable.node.logNodeReport( 
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
    routingTable.node.LookUP( nodeId, &successor)
    return successor
}


// [CORE-MODULE]
func (routingTable *RoutingTable) PrintRoutingInfo(){
    routingTable.node.currentSuccessorTableInfo()
	routingTable.node.currentFingerTableInfo()
    routingTable.node.currentSuccessorsInfo()
}
// # ------------------------ [END] Routing Functionalities ----------------------- #


// # ------------------------ Routing Info [Exported] ----------------------- #
func (routingTable *RoutingTable) NodeInfo() *Node{
   return routingTable.node
}

func (routingTable *RoutingTable) PrintCurrentNodeInfo() {
    routingTable.node.currentNodeInfo()
}

func (routingTable *RoutingTable) PrintUpdatedNodeInfo() {
    routingTable.node.updatedNodeInfo()
}

func (routingTable *RoutingTable) PrintCurrentReplicaInfo() {
    routingTable.node.currentReplicaInfo()
}

func (routingTable *RoutingTable) PrintRemoteReplicaInfo() {
    routingTable.node.remoteReplicaInfo()
}

