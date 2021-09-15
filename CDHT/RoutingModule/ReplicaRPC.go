package RoutingModule

import (
	"fmt"
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
    ReplicaAddress []NodeRPC
}


// [Internal]
func (node *Node) updateReplicaInfo() {
	err, _ := node.MainReplicaNode.NodeReplicaInfo( &node.ReplicaInfos )

    if err != nil {
        fmt.Println("[Replica][Error]:", err)
		return
    }

	node.updateRoutingEntries()
}

// [RPC]
func (node *Node) MakeNodeActive(args *Args, currentReplicaInfo  *ReplicaInfo) error {
	node.NodeState = NODE_STATE_ACTIVE
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
	node.ReplicaInfos.ReplicaAddress = currentActive
}



func copyReplicaInfo(old *ReplicaInfo, new *ReplicaInfo) {
	new.SuccessorsTable = old.SuccessorsTable
	new.Successor = old.Successor
	new.Predcessor = old.Predcessor
	
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