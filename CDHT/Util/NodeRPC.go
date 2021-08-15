package Util

import (
    "net/rpc"
    "log"
	"math/big"
	"time"
)

type Args struct {
	DefaultParams string
}

type NodeRPC struct {
	handle *rpc.Client
	Node_address string
	
	Node_id *big.Int 
	M *big.Int 

	DefaultArgs *Args
}



func (node *NodeRPC) Connect() (error, *NodeRPC) {
	client, err := rpc.Dial("tcp", node.Node_address)
    if err != nil {
		log.Println("[NODE-CONNECT] dialing:", err)
		return err, nil
    } 
	
	node.handle = client
	node.DefaultArgs = &Args{}

	err = node.handle.Call("Node.GetNodeInfo", node.DefaultArgs, &node)
    
	if err != nil {
        log.Println("Node.GetNodeInfo:", err)
		return err, nil
    }

	go node.autoClose()
	return nil, node
}


func (node *NodeRPC) GetNodeInfo() (error, *NodeRPC) {
	err := node.handle.Call("Node.GetNodeInfo", node.DefaultArgs, node)
    
	if err != nil {
        log.Println("Node.GetNodeInfo:", err)
		return err, nil
    }
	return nil, node
}


func (node *NodeRPC) FindSuccessor(nodeId *big.Int) (error, *NodeRPC) {
	var successor NodeRPC

	err := node.handle.Call("Node.FindSuccessor", nodeId, &successor)
    if err != nil {
        log.Println("Node.FindSuccessor:", err)
		return err, nil
    }
	return nil, &successor
}


func (node *NodeRPC) GetSuccessor() (error, *NodeRPC) {
	var successor NodeRPC

	err := node.handle.Call("Node.GetSuccessor", node.DefaultArgs, &successor)
    if err != nil {
        log.Println("Node.GetSuccessor:", err)
		return err, nil
    }
	return nil, &successor
}


func (node *NodeRPC) GetPredecessor() (error, *NodeRPC) {
	var predecessor NodeRPC

	err := node.handle.Call("Node.GetPredecessor", node.DefaultArgs, &predecessor)
    if err != nil {
        log.Println("Node.GetPredecessor:", err)
		return err,  nil
    }

	return nil,  &predecessor
}


func (node *NodeRPC) Notify(predecessor *NodeRPC) (error, *NodeRPC) {
	var successor NodeRPC

	err := node.handle.Call("Node.Notify", predecessor, &successor)
    if err != nil {
        log.Println("Node.Notify:", err)
		return err, nil
    }

	return nil,  &successor
}


func (node *NodeRPC) Close() {
	if node.handle == nil {
		return 
	}

	node.DefaultArgs = nil
	node.handle.Close()
}

func (node *NodeRPC) autoClose(){
	ticker := time.NewTicker(time.Second * 15)
	defer ticker.Stop()

	<-ticker.C
	node.handle.Close()
	node.DefaultArgs = nil
	log.Println("connection closed...", node.Node_address)
}