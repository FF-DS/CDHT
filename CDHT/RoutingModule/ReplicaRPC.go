package RoutingModule

import (
	"fmt"
	"sort"
	"time"
)

const (
	NODE_STATE_ACTIVE   string  = "STATE_ACTIVE"
	NODE_STATE_REPLICA  string  = "STATE_REPLICA"
)


type ReplicaInfo struct {
	SuccessorsTable Successors
	FingerTable map[int]NodeRPC
	Successor NodeRPC
	Predcessor NodeRPC

    JumpSpacing int
    SuccessorsTableLength int

	NodeState string
	MasterNode NodeRPC
    ReplicaAddress []NodeRPC
}

// ## ------------------------ [init] ------------------------ ## 
// [INTERNAL-MODULE]
func (node *Node) checkIfNodeCanBeReplica(mainNode *NodeRPC) bool {
    replicaInfos := ReplicaInfo{ SuccessorsTable : Successors{}, FingerTable : map[int]NodeRPC{}, Successor : NodeRPC{}, Predcessor : NodeRPC{}, ReplicaAddress : []NodeRPC{},  MasterNode : NodeRPC{}}	
    
    if checkNode(mainNode) == nil {
        return false
    }

    err, _ := mainNode.NodeReplicaInfo(&replicaInfos)

    if err != nil {
        return false
    }
    return len(replicaInfos.ReplicaAddress) < mainNode.ReplicationCount
}


// [INTERNAL-MODULE]
func (node *Node) internalReplicaInit(mainNode *NodeRPC){
    node.MainReplicaNode = mainNode

    node.NodeState = NODE_STATE_REPLICA
    node.Node_id = node.MainReplicaNode.Node_id
    node.M = node.MainReplicaNode.M

    remoteRep := NodeRPC{}
    node.GetNodeInfo(node.defaultArgs, &remoteRep)

    if checkNode(mainNode) == nil {
        fmt.Println("[JOIN][Replica][Error]: Main node is not Alive!")
        return
    }

    err, replicaInfo := mainNode.AddReplica( &remoteRep )

    if err == nil {
        node.ReplicaInfos = replicaInfo

        fmt.Println("[JOIN][Replica]: Joined as a replica of ", mainNode.Node_id)
    }else{
        fmt.Println("[JOIN][Replica][Error]:", err)
    }
}

// ## ------------------------ [End - init] ------------------------ ## 


// [Internal]
func (node *Node) updateReplicaInfo() {
	if checkNode(node.MainReplicaNode) == nil {
        fmt.Println("[Replica][Error]: Master node is down!")
		node.updateRemoteReplica()
		return
	}
	err, _ := node.MainReplicaNode.NodeReplicaInfo( &node.ReplicaInfos )

    if err != nil {
        fmt.Println("[Replica][Error]:", err)
		return
    }

	if node.MainReplicaNode.Node_address != node.ReplicaInfos.MasterNode.Node_address {
		node.MainReplicaNode = &node.ReplicaInfos.MasterNode
	}
	node.updateRoutingEntries()
}


func (node *Node) updateRemoteReplica() {
	sort.SliceStable(node.ReplicaInfos.ReplicaAddress, func(i, j int) bool {
		return node.ReplicaInfos.ReplicaAddress[i].Node_address < node.ReplicaInfos.ReplicaAddress[j].Node_address
	})

	for _, replica := range node.ReplicaInfos.ReplicaAddress {
		replica.DefaultArgs = nil
		if remoteNode := checkNode(&replica); remoteNode != nil && replica.NodeState == NODE_STATE_ACTIVE {
			node.MainReplicaNode = remoteNode
			fmt.Println("[Replica][UPDATE]: Master node changed to ",remoteNode.Node_address)
			return
		}
		
		if remoteNode := checkNode(&replica); remoteNode != nil {
			remoteNode.MakeNodeActive(&node.ReplicaInfos)
			node.MainReplicaNode = remoteNode
			fmt.Println("[Replica][UPDATE]: Master node changed to ",remoteNode.Node_address)
			return
		}
	}
}


// [RPC]
func (node *Node) MakeNodeActive(args *Args, currentReplicaInfo  *ReplicaInfo) error {
	node.NodeState = NODE_STATE_ACTIVE
	node.GetNodeInfo(node.defaultArgs, node.MainReplicaNode)
	node.GetNodeInfo(node.defaultArgs, &node.ReplicaInfos.MasterNode)

	currentActive := []NodeRPC{}
	for _,replica := range node.ReplicaInfos.ReplicaAddress {
		if replica.Node_address !=  node.MainReplicaNode.Node_address{ 
			currentActive = append(currentActive, replica)
		}
	}
	node.ReplicaInfos.ReplicaAddress = currentActive

	copyReplicaInfo(&node.ReplicaInfos, currentReplicaInfo)
	return nil
}

// [RPC]
func (node *Node) NodeReplicaInfo(args *Args, currentReplicaInfo  *ReplicaInfo) error {
	copyReplicaInfo(&node.ReplicaInfos, currentReplicaInfo)
	return nil
}

// [RPC]
func (node *Node) AddReplica(remoteNode  *NodeRPC, currentReplicaInfo  *ReplicaInfo) error {
	if node.NodeState == NODE_STATE_REPLICA {
		node.MainReplicaNode.AddReplica(remoteNode)
		return nil
	}

	if node.ReplicaInfos.ReplicaAddress == nil {
		node.ReplicaInfos.ReplicaAddress = []NodeRPC{}
	}
	
	node.ReplicaInfos.ReplicaAddress = append(node.ReplicaInfos.ReplicaAddress, *remoteNode)
	copyReplicaInfo(&node.ReplicaInfos, currentReplicaInfo)
	return nil
}


// [ROUTING-MODULE]
func (node *Node) currentNodeReplicaInfo() {
	node.ReplicaInfos.SuccessorsTable = node.currentSuccessors
	if (node.successor != nil) {
		node.ReplicaInfos.Successor = *node.successor
	}
	if (node.predecessor != nil) {
		node.ReplicaInfos.Predcessor = *node.predecessor
	}
	
	node.ReplicaInfos.FingerTable = map[int]NodeRPC{}
	node.ReplicaInfos.SuccessorsTableLength = node.SuccessorsTableLength
	
	node.ReplicaInfos.JumpSpacing = node.JumpSpacing
	node.ReplicaInfos.NodeState = node.NodeState

	for key, nodeConn := range node.fingerTableEntry {
		node.ReplicaInfos.FingerTable[ key ] = *nodeConn
	}

	currentActive := []NodeRPC{}
	for _,replica := range node.ReplicaInfos.ReplicaAddress {
		if checkNode(&replica) != nil{ 
			currentActive = append(currentActive, replica)
		}
	}
	if node.ReplicaInfos.MasterNode.Node_address == "" {
		node.GetNodeInfo(node.defaultArgs, &node.ReplicaInfos.MasterNode)
	}
	node.ReplicaInfos.ReplicaAddress = currentActive
}



func copyReplicaInfo(old *ReplicaInfo, new *ReplicaInfo) {
	new.SuccessorsTable = old.SuccessorsTable
	new.Successor = old.Successor
	new.Predcessor = old.Predcessor
	new.MasterNode = old.MasterNode
	
	new.SuccessorsTableLength = old.SuccessorsTableLength
	
	new.JumpSpacing = old.JumpSpacing
	new.NodeState = old.NodeState

	if new.FingerTable == nil {
		new.FingerTable = map[int]NodeRPC{}
	}
	for key, nodeConn := range old.FingerTable {
		nodeConn.DefaultArgs = nil
		new.FingerTable[ key ] = nodeConn
	}

	new.ReplicaAddress = old.ReplicaAddress 
}



// [Internal]
func (node *Node) updateRoutingEntries(){
	for _, succ := range node.ReplicaInfos.SuccessorsTable {
		succ.DefaultArgs = nil
	}
	node.currentSuccessors = node.ReplicaInfos.SuccessorsTable
	node.ReplicaInfos.Successor.DefaultArgs = nil
	node.successor = &node.ReplicaInfos.Successor 
	node.ReplicaInfos.Predcessor.DefaultArgs = nil
	node.predecessor = &node.ReplicaInfos.Predcessor 
	node.FingerTableLength = len(node.ReplicaInfos.FingerTable)
	node.SuccessorsTableLength = node.ReplicaInfos.SuccessorsTableLength
	node.JumpSpacing = node.ReplicaInfos.JumpSpacing
    node.ReplicationCount = node.MainReplicaNode.ReplicationCount

	currentFingerTable := map[int]*NodeRPC{}
	for key, nodeConn := range node.ReplicaInfos.FingerTable {
		var entry NodeRPC
		copyNodeData(&nodeConn, &entry)
		currentFingerTable[key] = &entry
	}
	node.fingerTableEntry = currentFingerTable
}

// ## ---------------------------- [printer] ---------------------------- ## 

func (node *Node) remoteReplicaInfo() {
	fmt.Printf("------------------[Remote Replica Info]------------------\n")
	fmt.Printf("  [+]: Remote Node Info \n", )
	printNodeRPC(node.MainReplicaNode, true)
	fmt.Printf("  [+]: Successors Table \n", )
	for _, succ := range node.ReplicaInfos.SuccessorsTable {
		printNodeRPC(&succ, false)
	}
	fmt.Printf("  [+]: Finger Table  \n")
	for i, finger := range node.ReplicaInfos.FingerTable{
		fmt.Printf("      [+]: Entry ID: %s\n", node.calculateFingerId(i).String())
		printNodeRPC(&finger, false)
	}
	fmt.Printf("  [+]: Successor \n")
	printNodeRPC(&node.ReplicaInfos.Successor, false)
	fmt.Printf("  [+]: Predcessor \n")
	printNodeRPC(&node.ReplicaInfos.Predcessor, false)
	fmt.Printf("  [+]: JumpSpacing : %d \n", node.ReplicaInfos.JumpSpacing)
	fmt.Printf("  [+]: SuccessorsTableLength : %d \n", node.ReplicaInfos.SuccessorsTableLength)
	fmt.Printf("  [+]: NodeState : %s \n", node.ReplicaInfos.NodeState)
	fmt.Printf("  [+]: REPLICAS \n")
	for _, replicaNode := range node.ReplicaInfos.ReplicaAddress {
		if node.NodeState == NODE_STATE_REPLICA {
			replicaNode.DefaultArgs = nil
		}
		printNodeRPC(&replicaNode, true)
	}
	fmt.Printf("----------------------------------------------------------\n")
}


func (node *Node) currentReplicaInfo() {
	fmt.Printf("------------------[Current Replica Info]------------------\n")
	fmt.Printf("  [+]: REPLICAS \n")
	for _, replicaNode := range node.ReplicaInfos.ReplicaAddress {
		if node.NodeState == NODE_STATE_REPLICA {
			replicaNode.DefaultArgs = nil
		}
		printNodeRPC(&replicaNode, true)
	}
	fmt.Printf("----------------------------------------------------------\n")
}


func (node *Node) getReplicas() [][]string {
	replicas := [][]string{}
	for _, replicaNode := range node.ReplicaInfos.ReplicaAddress {
		start := time.Now()
		if node.NodeState == NODE_STATE_REPLICA {
			replicaNode.DefaultArgs = nil
		}
		status := "Alive"
		if checkNode(&replicaNode) == nil {
			status = "Down"
		}
		replicas = append(replicas, []string{ 
			replicaNode.Node_address,
			status,
			time.Since(start).String(),
		})
	}

	return replicas
}


func printNodeRPC(remoteNode *NodeRPC, check bool) {
	if remoteNode == nil {
		return
	}
	status := "Alive"
	if check && checkNode(remoteNode) == nil {
		status = "Down"
	}
	fmt.Printf("      [+]: Node ID : %s [%s]\n", remoteNode.Node_id.String(), remoteNode.NodeState)
	fmt.Printf("      [+]: M       : %s \n", remoteNode.M.String())
	fmt.Printf("      [+]: Address : %s \n", remoteNode.Node_address)
	fmt.Printf("      [+]: Status  : %s \n", status)
	fmt.Println()
}