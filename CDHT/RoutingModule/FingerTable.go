package RoutingModule

import (
    "math/big"
    "fmt"
)

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


// [INTERNAL]
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


// [ROUTING-MODULE]
func (node *Node) fixFinger(){
    for i := 0; i < node.FingerTableLength; i++ { 
        var entry NodeRPC
        node.FindSuccessor( node.calculateFingerId(i), &entry)

        if checkNode(&entry) != nil {
            node.fingerTableEntry[i].Close()
            node.fingerTableEntry[i] = &entry
        } 
    }
}


// [INTERNAL]
func (node *Node) calculateFingerId(i int)  *big.Int {
	offset := new(big.Int).Exp(  big.NewInt( int64(node.JumpSpacing) ), big.NewInt(int64(i)), nil)

	sum := new(big.Int).Add( node.Node_id, offset)
	ceil :=  new(big.Int).Exp( big.NewInt( int64(node.JumpSpacing) ), node.M, nil)

	return new(big.Int).Mod(sum, ceil)
}




// # ------------------------ print info ----------------------- #

// [ROUTING-MODULE]
func (node *Node) currentFingerTableInfo() {
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