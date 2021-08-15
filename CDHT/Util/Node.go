package Util

import (
    "crypto/sha1"
    "net"  
    "math/big"
    "net/rpc"
    "fmt"
    "os"
)


// # --------------------------- NODE INFO --------------------------- # //

type Node struct {
	Node_id *big.Int 
	M *big.Int 

    Port string 
    IP_address string 

    predecessor *NodeRPC
    successor *NodeRPC

    fingerTableEntry map[int]*NodeRPC
    JumpSpacing int
    FingerTableLength int
    
    defaultArgs *Args
}






// # ---------------------------- Init --------------------------------- # 
func (node *Node) initNode() {
    node.generateNodeInfo()

    rpc.Register(node)

    tcpAddr, err := net.ResolveTCPAddr("tcp", ":" + node.Port)
    checkError(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)

    go rpc.Accept(listener)
}


func (node *Node) CreateRing() {
    node.initNode()

    var curr NodeRPC
    node.GetNodeInfo(node.defaultArgs, &curr)
    checkNode(&curr)

    node.successor = &curr
    node.predecessor = &curr
}


func (node *Node) Join(available *NodeRPC) {
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




// # --------------------- node level info ----------------------------- # 

func (node *Node) generateNodeInfo() Node {
    node.initializeNode()
	// node.getOutboundIP();
    node.IP_address = "127.0.0.1"
	// node.generateNodeId();
	return *node
}

func (node *Node) initializeNode() {
    node.defaultArgs = &Args{}
    node.predecessor = &NodeRPC{}
    node.successor = &NodeRPC{}
    node.fingerTableEntry = make(map[int]*NodeRPC)
}


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


func (node *Node) getOutboundIP() {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    checkError(err)

    defer conn.Close()
    localAddr := conn.LocalAddr().(*net.UDPAddr)
    node.IP_address = localAddr.IP.String()
}

// # ---------------------  NODE  ----------------------------- # 


// [RPC]
func (node *Node) GetNodeInfo(args *Args, nodeRPC *NodeRPC) error {
    nodeRPC.M = node.M
    nodeRPC.Node_address = node.IP_address + ":" + node.Port
    nodeRPC.Node_id = node.Node_id

    return nil
}


// [RPC]
func (node *Node) GetSuccessor(args *Args, nodeRPC *NodeRPC) error {
    if checkNode(node.successor) != nil {
        copyNodeData(node.successor, nodeRPC)
    }

    return nil
}

// [RPC]
func (node *Node) GetPredecessor(args *Args, nodeRPC *NodeRPC) error {
    if checkNode(node.predecessor) != nil {
        copyNodeData(node.predecessor, nodeRPC)
    }

    return nil
}





// # --------------------- finger table  ----------------------------- # 

// [RPC]
func (node *Node) FindSuccessor(nodeId *big.Int, remoteNode *NodeRPC) error {
    succ := checkNode( node.successor )
    if succ == nil {
        node.GetNodeInfo(node.defaultArgs, remoteNode)
        return nil
    }

    if between(node.Node_id, nodeId, succ.Node_id) || nodeId.Cmp(succ.Node_id) == 0 {
        copyNodeData(succ, remoteNode)
        return nil
    }else {
        pred := node.closestPrecedingNode(nodeId)

        if pred.Node_id.Cmp(node.Node_id) == 0 {
            if succ := checkNode(node.successor); succ != nil {
                copyNodeData(succ, remoteNode)

                return nil
            }
            node.GetNodeInfo(node.defaultArgs, remoteNode)
            return nil
        }


        err, pred2 := pred.FindSuccessor(nodeId)

        if err != nil || checkNode( pred2 ) == nil {
            node.GetNodeInfo(node.defaultArgs, remoteNode)
            return nil
        }

        copyNodeData(pred2, remoteNode)
    }
    return nil
}



func  (node *Node)  closestPrecedingNode(nodeId *big.Int) *NodeRPC {
    var curr NodeRPC
    node.GetNodeInfo(node.defaultArgs, &curr)

    for i := len(node.fingerTableEntry) - 1; i >= 0; i-- {
        entry := node.fingerTableEntry[i];

        if entry == nil {
            fmt.Println("empty entry ----------------", i)
            continue
        }
        
        if betweenClosest(node.Node_id, entry.Node_id, nodeId){
            if checkNode( entry ) == nil {
                continue
            }

            return entry
        }
    }

    checkNode( &curr)
    return &curr
}


func (node *Node) FixFinger(){
    for i := 0; i < node.FingerTableLength; i++ { 
        var entry NodeRPC
        node.FindSuccessor( node.calculateFingerId(i), &entry)

        if checkNode(&entry) != nil {
            node.fingerTableEntry[i].Close()
            node.fingerTableEntry[i] = &entry
        } 
    }
}


func (node *Node) calculateFingerId(i int)  *big.Int {
	offset := new(big.Int).Exp(  big.NewInt( int64(node.JumpSpacing) ), big.NewInt(int64(i)), nil)

	sum := new(big.Int).Add( node.Node_id, offset)
	ceil :=  new(big.Int).Exp( big.NewInt( int64(node.JumpSpacing) ), node.M, nil)

	return new(big.Int).Mod(sum, ceil)
}






// # --------------------- successor  ----------------------------- # 

func (node *Node) Stablize() {
    succ := checkNode(node.successor)

    if succ == nil {
        return 
    }


    err, pred := succ.GetPredecessor()
    

    if err != nil {
        fmt.Println("[STABLIZE][Error]:", err)
        return 
    }


    pred = checkNode( pred )
    
    
    if pred != nil && between(node.Node_id, pred.Node_id, succ.Node_id) {
        node.successor.Close()
        node.successor = pred
    }

    var curr NodeRPC
    node.GetNodeInfo(node.defaultArgs, &curr)
    node.successor.Notify(&curr)
}


// [RPC]
func (node *Node) Notify(pred *NodeRPC, curr *NodeRPC) error {
    if checkNode( node.predecessor ) == nil || between( node.predecessor.Node_id, pred.Node_id, node.Node_id) {
        if node.predecessor != nil {
            node.predecessor.Close()
        }

        copyNodeData(pred, node.predecessor)
    }

    node.GetNodeInfo(node.defaultArgs, curr)
    return nil
}   


func (node *Node) CheckPredecessor() {
    node.predecessor = checkNode(node.predecessor)
}







// # ------------------------ print info ----------------------- #

func (node *Node) CurrentNodeInfo() {
    fmt.Printf("-----------------Current node Info[%s]--------------------\n",node.Node_id.String())
    fmt.Printf("Node ID : %s \n", node.Node_id.String())
    fmt.Printf("M       : %s \n", node.M.String())
    fmt.Printf("Address : %s:%s \n", node.IP_address, node.Port)
    fmt.Println("----------------------------------------------------------")
}


func (node *Node) CurrentFingerTableInfo() {
    fmt.Printf("-----------------Finger Table Info[%s]--------------------\n",node.Node_id.String())
    for i := 0; i < len(node.fingerTableEntry); i++ {
        entry := checkNode( node.fingerTableEntry[i] )
        if entry != nil {
            fmt.Printf(" [%d]. Entry ID: |%s| Node ID : %s  Address : %s \n", i, node.calculateFingerId(i).String(), entry.Node_id.String(), entry.Node_address)
        }else{
            fmt.Printf(" [%d]. Entry ID: |%s| NOT AVAILABLE \n", i, node.calculateFingerId(i).String())
        }
    }
    fmt.Println("---------------------------------------------------------")
}


func (node *Node) CurrentSuccessorTableInfo() {
    fmt.Printf("-----------------Successor Table Info[%s]-----------------\n",node.Node_id.String())
    
    if succ := checkNode( node.successor);  succ != nil {
        fmt.Printf(" [SUCC] | Node ID : %s  Address : %s \n", succ.Node_id.String(), succ.Node_address)
    }else{
        fmt.Printf(" [SUCC] | NOT AVAILABLE \n")
    }

    if pred := checkNode( node.predecessor); pred != nil {
        fmt.Printf(" [PRED] | Node ID : %s  Address : %s \n", pred.Node_id.String(), pred.Node_address)
    }else{
        fmt.Printf(" [PRED] | NOT AVAILABLE \n")
    }

    fmt.Println("-----------------------------------------------------------")
}








// # ------------------------ helpers ----------------------- #

func between(start, middle, end *big.Int) bool {
	if res := start.Cmp(end); res == -1 {
		return start.Cmp(middle) == -1 && middle.Cmp(end) <= 0
	}
	return start.Cmp(middle) == -1 || middle.Cmp(end) <= 0
}


func betweenClosest(start, middle, end *big.Int) bool {
	if res := start.Cmp(end); res == -1 {
		return start.Cmp(middle) == -1 && middle.Cmp(end) < 0
	}
	return start.Cmp(middle) == -1 || middle.Cmp(end) < 0
}


func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}


func copyNodeData(old *NodeRPC, new *NodeRPC) {
    new.M = old.M
    new.Node_address = old.Node_address
    new.Node_id = old.Node_id
    new.DefaultArgs = nil
}


func checkNode(node *NodeRPC) *NodeRPC {
    if node == nil || node.Node_id == nil {
        return nil
    }

    var nodeRPC *NodeRPC
    if node.DefaultArgs == nil {
        // fmt.Println("CHECK connection")
        _, nodeRPC = node.Connect()
    }else{
        _, nodeRPC = node.GetNodeInfo()
    }
    
    return nodeRPC
}