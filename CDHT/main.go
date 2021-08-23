package main

import (
    "cdht/RoutingModule"
    "time"
)


func main() {

    go runFirstNode();

    // go runSecondNode();

    time.Sleep(time.Minute * 35)
}


func runFirstNode() {
    firstNode := RoutingModule.RoutingTable{ 
        RingPort: "9898",
        JumpSpacing: 2,
        FingerTableLength: 4,
    }

    firstNode.CreateRing()
}


func runSecondNode() {
    secondNode := RoutingModule.RoutingTable{ 
        RemoteNodeAddr: "127.0.0.1:9898",
        NodePort: "3456",
        JumpSpacing: 2,
        FingerTableLength: 4,
    }

    secondNode.RunNode()
}

