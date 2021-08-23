package RoutingModule

import (
    "net/rpc"
    "log"
	"math/big"
	"time"
)

// # -------------------- NodeRPC -------------------- # //

type NodeRPC struct {
	handle *rpc.Client
	Node_address string
	
	Node_id *big.Int 
	M *big.Int 

	DefaultArgs *Args
	ticker *time.Ticker
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
	
	node.ticker = time.NewTicker(time.Second * 15)
	go node.autoClose()
	return nil, node
}


func (node *NodeRPC) GetNodeInfo() (error, *NodeRPC) {
	node.resetConnTimeOut();

	err := node.handle.Call("Node.GetNodeInfo", node.DefaultArgs, node)
    
	if err != nil {
        log.Println("Node.GetNodeInfo:", err)
		return err, nil
    }
	return nil, node
}


func (node *NodeRPC) FindSuccessor(nodeId *big.Int) (error, *NodeRPC) {
	node.resetConnTimeOut();

	var successor NodeRPC

	err := node.handle.Call("Node.FindSuccessor", nodeId, &successor)
    if err != nil {
        log.Println("Node.FindSuccessor:", err)
		return err, nil
    }
	return nil, &successor
}


func (node *NodeRPC) GetSuccessor() (error, *NodeRPC) {
	node.resetConnTimeOut();

	var successor NodeRPC

	err := node.handle.Call("Node.GetSuccessor", node.DefaultArgs, &successor)
    if err != nil {
        log.Println("Node.GetSuccessor:", err)
		return err, nil
    }
	return nil, &successor
}


func (node *NodeRPC) GetPredecessor() (error, *NodeRPC) {
	node.resetConnTimeOut();

	var predecessor NodeRPC

	err := node.handle.Call("Node.GetPredecessor", node.DefaultArgs, &predecessor)
    if err != nil {
        log.Println("Node.GetPredecessor:", err)
		return err,  nil
    }

	return nil,  &predecessor
}


func (node *NodeRPC) Notify(predecessor *NodeRPC) (error, *Successors) {
	node.resetConnTimeOut();
	
	successors := Successors{}

	err := node.handle.Call("Node.Notify", predecessor, &successors)
    if err != nil {
        log.Println("Node.Notify:", err)
		return err, nil
    }

	return nil,  &successors
}


func (node *NodeRPC) Close() {
	if node == nil || node.handle == nil {
		return 
	}

	node.DefaultArgs = nil
	node.handle.Close()
}

func (node *NodeRPC) autoClose(){
	defer node.ticker.Stop()

	<-node.ticker.C
	node.handle.Close()
	node.DefaultArgs = nil
	// log.Println("connection closed...", node.Node_address)
}

func (node *NodeRPC) resetConnTimeOut(){
	node.ticker.Stop()
	node.ticker = time.NewTicker(time.Second * 10)
}


// # -------------------- Successors -------------------- # //

type Successors []NodeRPC

func (successors *Successors) UpdateSuccessors(newSuccessors Successors) {
	*successors = newSuccessors
}

func (successors *Successors) PopFirst() {
	*successors = (*successors)[1:]
}

// # -------------------- Args -------------------- # //

type Args struct {
	DefaultParams string
}



