package RoutingModule

import (
    "fmt"
	"time"
	"strconv"
    "math/big"
	"cdht/Util"
)


type RoutingTable struct {
	RingPort string
	
	RemoteNodeAddr string
	NodePort string

	JumpSpacing int
	FingerTableLength int

    Applications map[string](chan Util.RequestObject)
	node *Node
}



// # ------------------------ Main Node INIT ----------------------- #

func (routingTable *RoutingTable) CreateRing() {
    routingTable.node = &Node{
        Port : routingTable.RingPort,
        JumpSpacing : routingTable.JumpSpacing,
        FingerTableLength : routingTable.FingerTableLength,
        Applications : routingTable.Applications,
    }

	// TESTING WILL BE REMOVED //
    testChangeNodeInfo(routingTable.node)
	// END OF TESTING //

    routingTable.node.currentNodeInfo()
    routingTable.node.createRing()
    
    go routingTable.runStablization()
}



func (routingTable *RoutingTable) RunNode() {

    remoteNode := NodeRPC{ Node_address : routingTable.RemoteNodeAddr }
    // TESTING WILL BE REMOVED //
	remoteNodeAddr(&remoteNode)
	// END OF TESTING //
	remoteNode.Connect()

    printRemoteNodeInfo(&remoteNode)


    routingTable.node = &Node{
        Port : routingTable.NodePort,
        JumpSpacing : routingTable.JumpSpacing,
        FingerTableLength : routingTable.FingerTableLength,
        Applications : routingTable.Applications,
    }

	// TESTING WILL BE REMOVED //
	currentNodePort(routingTable.node)
    testChangeNodeInfo(routingTable.node)
	// END OF TESTING //

    routingTable.node.currentNodeInfo()
    routingTable.node.join( &remoteNode)
    
    go routingTable.runStablization()
}



func (routingTable *RoutingTable) runStablization() {
	for {
		time.Sleep(time.Second)
		routingTable.node.checkPredecessor()
		routingTable.node.checkSeccessors()

		routingTable.node.stablize()
		routingTable.node.currentSuccessorTableInfo()

		routingTable.node.fixFinger()
		routingTable.node.currentFingerTableInfo()

		routingTable.node.currentSuccessorsInfo()
	}
}

// # ------------------------ [END] Main Node INIT ----------------------- #





// # --------------------------- Routing Functionalities -------------------------- #

// [CORE-MODULE]
func (routingTable *RoutingTable) ForwardPacket(req Util.RequestObject) Util.ResponseObject {
    var successor NodeRPC
    err := routingTable.node.FindSuccessor( req.ReceiverNodeId, &successor)

    if err == nil && checkNode(&successor) != nil{
        _, resp := successor.ResolvePacket( req )
        return resp
    }

    resp := req.GetResponseObject()
    resp.ResponseStatus = Util.PACKET_STATUS_FAILED
    return resp
}



// # ------------------------ [END] Routing Functionalities ----------------------- #





// # ------------------------ print info ----------------------- #

func printRemoteNodeInfo(remoteNode *NodeRPC) {
    fmt.Println("-----------------Remote node Info------------------------------")
    fmt.Printf("Node ID : %s \n", remoteNode.Node_id.String())
    fmt.Printf("M       : %s \n", remoteNode.M.String())
    fmt.Printf("Address : %s \n", remoteNode.Node_address)
    fmt.Println("---------------------------------------------------------------")
}


// --------------- TEST ----------------------- //
func testChangeNodeInfo(node *Node) {
    var nodeId, m string
    fmt.Print("Enter Node Id: ")
    fmt.Scanln(&nodeId)

    fmt.Print("Enter M: ")
    fmt.Scanln(&m)

    NodeId, ok := new(big.Int).SetString(nodeId, 10)
    if !ok { fmt.Println("SetString: error node id") }

    M, ok := new(big.Int).SetString(m, 10)
    if !ok { fmt.Println("SetString: error m") }

    node.Node_id = NodeId
    node.M = M

    i, _ := strconv.Atoi(m)
    node.FingerTableLength = i
}


func remoteNodeAddr(node *NodeRPC) {
	var port string
    fmt.Print("Enter Remote Node Port: ")
    fmt.Scanln(&port)
	node.Node_address = "127.0.0.1:" + port
}


func currentNodePort(node *Node){
	var port string
	fmt.Print("Enter Your Port: ")
    fmt.Scanln(&port)
	node.Port = port
}

