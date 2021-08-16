package main

import (
    "fmt"
    "math/big"
    "cdht/Util"
    "time"
    "strconv"
)


func main() {
    go runFirstNode();

    // go runSecondNode();

    time.Sleep(time.Minute * 35)
}


func runFirstNode() {
    node := Util.Node{
        Port : "9898",
        JumpSpacing : 2,
        FingerTableLength : 4,
    }
    testChangeNodeInfo(&node)
    node.CurrentNodeInfo()

    node.CreateRing()
    
    
    go func() {
        for {
            time.Sleep(time.Second)
            node.CheckPredecessor()
            node.CheckSeccessors()

            node.Stablize()
            node.CurrentSuccessorTableInfo()

            node.FixFinger()
            node.CurrentFingerTableInfo()

            node.CurrentSuccessorsInfo()
        }
    }()

}


func runSecondNode() {
    var port string
    fmt.Print("Enter Node Port: ")
    fmt.Scanln(&port)


    remoteNode := Util.NodeRPC{ Node_address : "127.0.0.1:" + port }
    remoteNode.Connect()
    printRemoteNodeInfo(&remoteNode)


    fmt.Print("Enter Your Port: ")
    fmt.Scanln(&port)

    node := Util.Node{
        Port : port,
        JumpSpacing : 2,
        FingerTableLength : 4,
    }
    testChangeNodeInfo(&node)
    node.CurrentNodeInfo()
    
    node.Join( &remoteNode)
    
    go func() {
        for {
            time.Sleep(time.Second)

            node.CheckPredecessor()
            node.CheckSeccessors()

            node.Stablize()
            node.CurrentSuccessorTableInfo()

            node.FixFinger()
            node.CurrentFingerTableInfo()

            node.CurrentSuccessorsInfo()
        }
    }()
}


func check(){
    remoteNode := Util.NodeRPC{ Node_address : "127.0.0.1:9898" }
    err, remote := remoteNode.Connect()
    for {
        err, remote = remoteNode.GetNodeInfo()
        
        if err != nil {
            fmt.Println("error")
        }
        printRemoteNodeInfo(remote)
    }
}


func printRemoteNodeInfo(remoteNode *Util.NodeRPC) {
    fmt.Println("-----------------Remote node Info------------------------------")
    fmt.Printf("Node ID : %s \n", remoteNode.Node_id.String())
    fmt.Printf("M       : %s \n", remoteNode.M.String())
    fmt.Printf("Address : %s \n", remoteNode.Node_address)
    fmt.Println("---------------------------------------------------------------")
}

func testChangeNodeInfo(node *Util.Node) {
    var nodeId, m string
    fmt.Print("Enter Node Id: ")
    fmt.Scanln(&nodeId)

    fmt.Print("Enter M: ")
    fmt.Scanln(&m)

    NodeId, ok := new(big.Int).SetString(nodeId, 10)
    if !ok { fmt.Println("SetString: error node id") }

    M, ok := new(big.Int).SetString(m, 10)
    if !ok { fmt.Println("SetString: error m") }

    node.Node_id = NodeId
    node.M = M

    i, _ := strconv.Atoi(m)
    node.FingerTableLength = i
}

