package RoutingModule

import  (
    "math/big"
    "os"
    "fmt"
)


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