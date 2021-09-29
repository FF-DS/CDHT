package RoutingModule

import (
    "crypto/sha1"
    "cdht/Util"
    "cdht/ReportModule"
    "net"  
    "math/big"
    "net/rpc"
    "strconv"
    "fmt"
    "os"
)


type Node struct {
	Node_id *big.Int 
	M *big.Int 

    Port string 
    IP_address string 

    Applications  map[string](chan Util.RequestObject)
    NetworkTools chan Util.RequestObject

    predecessor *NodeRPC
    successor *NodeRPC
    currentSuccessors Successors

    fingerTableEntry map[int]*NodeRPC
    FingerTableLength int
    SuccessorsTableLength int

    JumpSpacing int
    defaultArgs *Args
    Logger *ReportModule.Logger


    // Replication
    NodeState string
    ReplicaInfos ReplicaInfo
    MainReplicaNode *NodeRPC
    ReplicationCount int
}




// # ---------------------------- Init --------------------------------- # 

// [INTERNAL]
func (node *Node) initNode() {
    node.NodeState = NODE_STATE_ACTIVE

    node.generateNodeInfo()

    rpc.Register(node)

    tcpAddr, err := net.ResolveTCPAddr("tcp", ":" + node.Port)
    checkError(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)

    go rpc.Accept(listener)
}


func (node *Node) initReplica(){
    node.NodeState = NODE_STATE_REPLICA

    node.Node_id = node.MainReplicaNode.Node_id
    node.M = node.MainReplicaNode.M

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
    
    if checkNode(available) == nil {
        fmt.Println("[JOIN][Error]: available node is not Alive!")
        os.Exit(0)
        return
    }

    err, succ := available.FindSuccessor(node.Node_id)
    
    // check replica 
    if canBeReplica := node.checkIfNodeCanBeReplica(succ); canBeReplica {
        node.internalReplicaInit(succ)
        return
    }


    succ.NodeTraversalLogs =  []NodeRPC{}
    if err == nil && checkNode(succ) != nil {
        node.successor = succ 
        fmt.Println("[JOIN]: Joined with ", succ.Node_id)
    }else{
        fmt.Println("[JOIN][Error]:", err)
        os.Exit(0)
    }


    var curr NodeRPC
    node.GetNodeInfo(node.defaultArgs, &curr)
    checkNode(&curr)

    node.predecessor = &curr
}


// [ROUTING-MODULE]
func (node *Node) makeReplicaOf(mainNode *NodeRPC) {
    node.MainReplicaNode = mainNode
    node.initReplica()
    
    remoteRep := NodeRPC{}
    node.GetNodeInfo(node.defaultArgs, &remoteRep)

    if checkNode(mainNode) == nil {
        fmt.Println("[JOIN][Replica][Error]: Main node is not Alive!")
        os.Exit(0)
        return
    }

    err, replicaInfo := mainNode.AddReplica( &remoteRep )

    if err == nil {
        node.ReplicaInfos = replicaInfo

        fmt.Println("[JOIN][Replica]: Joined as a replica of ", mainNode.Node_id)
    }else{
        fmt.Println("[JOIN][Replica][Error]:", err)
        os.Exit(0)
    }
}

// # --------------------------- [END] Init  --------------------------- # //




// # --------------------- node level info ----------------------------- # 

// [INTERNAL]
func (node *Node) generateNodeInfo() Node {
    node.initializeNode()
    if node.IP_address == "" {
        node.getOutboundIP();
    }

    if node.Node_id == nil && node.NodeState != NODE_STATE_REPLICA {
	    node.generateNodeId();
    }
	return *node
}


// [INTERNAL]
func (node *Node) initializeNode() {
    node.defaultArgs = &Args{}
    node.predecessor = &NodeRPC{}
    node.successor = &NodeRPC{}
    node.currentSuccessors = Successors{}
    node.fingerTableEntry = make(map[int]*NodeRPC)
    node.ReplicaInfos = ReplicaInfo{ SuccessorsTable : Successors{}, FingerTable : map[int]NodeRPC{}, Successor : NodeRPC{}, 
                            Predcessor : NodeRPC{},ReplicaAddress : []NodeRPC{}, MasterNode : NodeRPC{}}
}

// [INTERNAL]
func (node *Node) generateNodeId() {	
    nodeIdentification := node.IP_address + ":" + node.Port

    hashFunction := sha1.New()
    hashFunction.Write([]byte(nodeIdentification))
    sha := hashFunction.Sum(nil)

    base, m, hashedID := big.NewInt( int64(node.JumpSpacing) ), node.M,  (&big.Int{}).SetBytes(sha)

    modulo := base.Exp( base, m, nil)

    node.Node_id = hashedID.Mod(hashedID, modulo)
    node.M = m

    fingerTableLen, _ := strconv.Atoi(m.String())
    node.FingerTableLength = fingerTableLen
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
    if nodeRPC ==  nil || node == nil {
        fmt.Println("[Network][Error]: it seems you got disconnected!")
        os.Exit(0)
    }

    nodeRPC.M = node.M
    nodeRPC.Node_address = node.IP_address + ":" + node.Port
    nodeRPC.NodeState = node.NodeState
    nodeRPC.Node_id = node.Node_id
    nodeRPC.ReplicationCount = node.ReplicationCount

    return nil
}

func (node *Node) getLocalNodeInfo() NodeRPC {
    return NodeRPC{
        M : node.M,
        Node_address : node.IP_address + ":" + node.Port,
        Node_id : node.Node_id,
        NodeState : node.NodeState,
    }
}
// # ----------------------- [END] node level info ----------------------------- # 





// # ----------------------  NODE FUNCTIONALITIES  ----------------------------- # 

// [RPC]
func (node *Node) ResolvePacket(requestObject *Util.RequestObject, responseObject *Util.RequestObject) error {
    if node.NodeState == NODE_STATE_ACTIVE {
        node.sendacketToReplicas(requestObject)
    }

    if requestObject.Type == Util.PACKET_TYPE_NETWORK {
        node.NetworkTools <- *requestObject


    } else if appChannel, exists := node.Applications[ requestObject.AppName ]; exists {
        appChannel <- *requestObject
        response := requestObject.GetResponseObject()
        copyResponseObject( &response, responseObject)

        // ---- logging ----- //
        node.logNodeReport( 
            ReportModule.LOG_TYPE_NODE_INFORMATION,
            ReportModule.LOG_LOCATION_TYPE_INCOMMING,
            ReportModule.LOG_OPERATION_STATUS_SUCCESS,
            map[string]string{
                "rtt" : "Not Available",
                "srt" : "Not Available",
                "latency" : "Not Available",
                "app_name" : requestObject.AppName,
            },
        )
        // ---- logging ----- //

        return nil
    }

    resp := requestObject.GetResponseObject()
    resp.ResponseStatus = Util.PACKET_STATUS_NO_APP
    copyResponseObject( &resp, responseObject)

    return nil
}



func (node *Node) sendacketToReplicas(requestObject *Util.RequestObject) bool {
    if !requestObject.SendToReplicas {
        return true
    }

    success := true
    for _, replicas := range node.ReplicaInfos.ReplicaAddress {
        if checkNode(&replicas) == nil {
            continue
        }
        err, _ := replicas.ResolvePacket( *requestObject )

        if err != nil {
            success = false
        }
    }

    return success
}
// # ------------------ [END] NODE FUNCTIONALITIES ----------------------------- # 





// # ------------------------ print info ----------------------- #

// [ROUTING-MODULE]
func (node *Node) currentNodeInfo() {
    fmt.Printf("                                ------------------Current node Info[%s]------------------\n",node.Node_id.String())
    fmt.Printf("                                      [+]: Node ID : %s [%s] \n", node.Node_id.String(), node.NodeState)
    fmt.Printf("                                      [+]: M       : %s \n", node.M.String())
    fmt.Printf("                                      [+]: Address : %s:%s \n", node.IP_address, node.Port)
    fmt.Printf("                                ---------------------------------------------------------\n")
}





