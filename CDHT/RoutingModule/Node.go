package RoutingModule

import (
    "crypto/sha1"
    "cdht/Util"
    "net"  
    "math/big"
    "net/rpc"
    "fmt"
)


type Node struct {
	Node_id *big.Int 
	M *big.Int 

    Port string 
    IP_address string 

    Applications  map[string](chan Util.RequestObject)

    predecessor *NodeRPC
    successor *NodeRPC
    currentSuccessors Successors

    fingerTableEntry map[int]*NodeRPC
    FingerTableLength int

    JumpSpacing int
    defaultArgs *Args
}




// # ---------------------------- Init --------------------------------- # 

// [INTERNAL]
func (node *Node) initNode() {
    node.generateNodeInfo()

    rpc.Register(node)

    tcpAddr, err := net.ResolveTCPAddr("tcp", ":" + node.Port)
    checkError(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)

    go rpc.Accept(listener)
}


// [ROUTING-MODULE]
func (node *Node) createRing() {
    node.initNode()

    var curr NodeRPC
    node.GetNodeInfo(node.defaultArgs, &curr)
    checkNode(&curr)

    node.successor = &curr
    node.predecessor = &curr
}


// [ROUTING-MODULE]
func (node *Node) join(available *NodeRPC) {
    node.initNode()
    
    err, succ := available.FindSuccessor(node.Node_id)
    
    if err == nil && checkNode(succ) != nil {
        node.successor = succ 
        fmt.Println("[JOIN]: Joined with ", succ.Node_id)
    }else{
        fmt.Println("[JOIN][Error]:", err)
    }


    var curr NodeRPC
    node.GetNodeInfo(node.defaultArgs, &curr)
    checkNode(&curr)

    node.predecessor = &curr
}

// # --------------------------- [END] Init  --------------------------- # //




// # --------------------- node level info ----------------------------- # 

// [INTERNAL]
func (node *Node) generateNodeInfo() Node {
    node.initializeNode()
	// node.getOutboundIP();
    node.IP_address = "127.0.0.1"
	// node.generateNodeId();
	return *node
}

// [INTERNAL]
func (node *Node) initializeNode() {
    node.defaultArgs = &Args{}
    node.predecessor = &NodeRPC{}
    node.successor = &NodeRPC{}
    node.currentSuccessors = Successors{}
    node.fingerTableEntry = make(map[int]*NodeRPC)
}

// [INTERNAL]
func (node *Node) generateNodeId() {	
    nodeIdentification := node.IP_address + ":" + node.Port

    hashFunction := sha1.New()
    hashFunction.Write([]byte(nodeIdentification))
    sha := hashFunction.Sum(nil)

    two, m, hashedID := big.NewInt(2), big.NewInt(160),  (&big.Int{}).SetBytes(sha)

    modulo := two.Exp( two, m, nil)

    node.Node_id = hashedID.Mod(hashedID, modulo)
    node.M = m
}


// [INTERNAL]
func (node *Node) getOutboundIP() {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    checkError(err)

    defer conn.Close()
    localAddr := conn.LocalAddr().(*net.UDPAddr)
    node.IP_address = localAddr.IP.String()
}

// [RPC]
func (node *Node) GetNodeInfo(args *Args, nodeRPC *NodeRPC) error {
    nodeRPC.M = node.M
    nodeRPC.Node_address = node.IP_address + ":" + node.Port
    nodeRPC.Node_id = node.Node_id

    return nil
}

// # ----------------------- [END] node level info ----------------------------- # 





// # ----------------------  NODE FUNCTIONALITIES  ----------------------------- # 

// [RPC]
func (node *Node) ResolvePacket(requestObject *Util.RequestObject, responseObject *Util.ResponseObject) error {
    if appChannel, exists := node.Applications[ requestObject.AppName ]; exists {

        appChannel <- *requestObject
        response := requestObject.GetResponseObject()
        copyResponseObject( &response, responseObject)
        return nil
    }

    resp := requestObject.GetResponseObject()
    resp.ResponseStatus = Util.PACKET_STATUS_NO_APP
    copyResponseObject( &resp, responseObject)
    return nil
}


// # ------------------ [END] NODE FUNCTIONALITIES ----------------------------- # 





// # ------------------------ print info ----------------------- #

// [ROUTING-MODULE]
func (node *Node) currentNodeInfo() {
    fmt.Printf("-----------------Current node Info[%s]--------------------\n",node.Node_id.String())
    fmt.Printf("Node ID : %s \n", node.Node_id.String())
    fmt.Printf("M       : %s \n", node.M.String())
    fmt.Printf("Address : %s:%s \n", node.IP_address, node.Port)
    fmt.Println("----------------------------------------------------------")
}





