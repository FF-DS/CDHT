package main

import (
    "fmt"
    // "math/big"
    "cdht/Util"
    "time"
    // "strconv"
)



func main() {
    go runFirstNode();

    // go runSecondNode();

    time.Sleep(time.Minute * 35)
}


func runFirstNode() {
    var node Util.Node
    node.CreateRing("9898")

    for {
        time.Sleep(time.Second * 5)
        node.CurrentSuccessorTableInfo()
        node.CheckPredecessor()
        node.Stablize()
    }

}


func runSecondNode() {
    remoteNode := Util.NodeRPC{ Node_address : "127.0.0.1:9898" }
    remoteNode.Connect()
    printRemoteNodeInfo(&remoteNode)


    var port string
    fmt.Print("Enter Port: ")
    fmt.Scanln(&port)

    var node Util.Node
    node.Join( &remoteNode, port)
    


    for {
        time.Sleep(time.Second * 5)
        node.CurrentSuccessorTableInfo()
        node.CheckPredecessor()
        node.Stablize()
    }
}


func printRemoteNodeInfo(remoteNode *Util.NodeRPC) {
    fmt.Println("-----------------Remote node Info------------------------------")
    fmt.Printf("Node ID : %s \n", remoteNode.Node_id.String())
    fmt.Printf("M       : %s \n", remoteNode.M.String())
    fmt.Printf("Address : %s \n", remoteNode.Node_address)
    fmt.Println("---------------------------------------------------------------")
}

// func testChangeFingerTableID(nodeInfo *Util.NodeInfo) int{
//     var nodeId, m string
//     fmt.Print("Enter Node Id: ")
//     fmt.Scanln(&nodeId)

//     fmt.Print("Enter M: ")
//     fmt.Scanln(&m)

//     NodeId, ok := new(big.Int).SetString(nodeId, 10)
//     if !ok { fmt.Println("SetString: error node id") }

//     M, ok := new(big.Int).SetString(m, 10)
//     if !ok { fmt.Println("SetString: error m") }

//     nodeInfo.Node_id = NodeId
//     nodeInfo.M = M

//     i, _ := strconv.Atoi(m)
//     return i
// }

