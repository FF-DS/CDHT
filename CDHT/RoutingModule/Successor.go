package RoutingModule

import (
    "fmt"
    "strconv"
)


// # --------------------- successor  ----------------------------- # 


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


// [ROUTING-MODULE]
func (node *Node) stablize() {
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
    _, nextSuccessors := node.successor.Notify(&curr)

    if nextSuccessors != nil {
        node.updateSuccessors( *nextSuccessors )
    }
}


// [INTERNAL]
func (node *Node) updateSuccessors(nextSuccessors Successors){
    node.currentSuccessors = Successors{}
    succ := NodeRPC{}
    copyNodeData(node.successor, &succ)
    node.currentSuccessors = append(node.currentSuccessors, succ)
    for i, succ := range nextSuccessors {
        if i == node.SuccessorsTableLength - 1 { break }
        node.currentSuccessors = append(node.currentSuccessors, succ)
    }
}


// [RPC]
func (node *Node) Notify(pred *NodeRPC, currSuccessors *Successors) error {
    if checkNode( node.predecessor ) == nil || between( node.predecessor.Node_id, pred.Node_id, node.Node_id) {
        if node.predecessor != nil {
            node.predecessor.Close()
        }
        
        if node.predecessor == nil {
            node.predecessor = &NodeRPC{} 
        }
        copyNodeData(pred, node.predecessor)
    }

    currSuccessors.UpdateSuccessors( node.currentSuccessors)
    return nil
}   


// [ROUTING-MODULE]
func (node *Node) checkPredecessor() {
    node.predecessor = checkNode(node.predecessor)
}


// [ROUTING-MODULE]
func (node *Node) checkSeccessors() {
    if checkNode(node.successor) != nil {
        return
    }

    for  checkNode(node.successor) == nil && len(node.currentSuccessors) > 1 {
        node.currentSuccessors.PopFirst()
        node.successor = &node.currentSuccessors[0]
    }
}



// # ------------------------ str info ----------------------- #
func (node *Node) getSuccessorsRouteEnty() []([]string) {
    routes := []( []string ){}

    for i := 0; i < len(node.currentSuccessors); i++ {
        entry := node.currentSuccessors[i] 
        routes = append( routes,  []string{ strconv.Itoa(i), entry.Node_id.String(), entry.Node_address  } )
    }
    return routes
}


func (node *Node) getSuccPredRouteEnty() []([]string)  {
    routes := []( []string ){ }

    if succ := checkNode( node.successor); succ != nil {
        routes = append( routes,   []string{ "Successor",  succ.Node_id.String(), succ.Node_address  } )
    }else{
        routes = append( routes,   []string{ "Successor",  "NOT AVAILABLE", "NOT AVAILABLE" } )
    }

    if pred := checkNode( node.predecessor); pred != nil {
        routes = append( routes,   []string{ "Predecessor", pred.Node_id.String(), pred.Node_address } )
    }else{
        routes = append( routes,   []string{ "Predecessor",  "NOT AVAILABLE", "NOT AVAILABLE" } )
    }
    return routes
}




// # ------------------------ print info ----------------------- #

// [ROUTING-MODULE]
func (node *Node) currentSuccessorsInfo() {
    fmt.Printf("-----------------Successors Table Info[%s]--------------------\n",node.Node_id.String())
    for i := 0; i < len(node.currentSuccessors); i++ {
        entry := node.currentSuccessors[i] 
        fmt.Printf(" [%d]. Node ID : %s  Address : %s \n", i + 1, entry.Node_id.String(), entry.Node_address)
    }
    fmt.Println("-------------------------------------------------------------")
}


// [ROUTING-MODULE]
func (node *Node) currentSuccessorTableInfo() {
    fmt.Printf("-----------------[Successor|Predecessor] Table Info[%s]-----------------\n",node.Node_id.String())
    
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

    fmt.Println("---------------------------------------------------------------------")
}
