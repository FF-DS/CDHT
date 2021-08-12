package NetworkModule

import (
    "net/rpc"
    "fmt"
    "log"
    "os"
)


type NodeInfoRPC struct {
	handle *rpc.Client
}


func (nodeInfo *NodeInfoRPC) Connect(address string)  (error, *NodeInfoRPC) {
	client, err := rpc.Dial("tcp", address)
    if err != nil {
		log.Fatal("dialing:", err)
		return err, nodeInfo
    } 
	
	nodeInfo.handle = client

	return nil, nodeInfo
}


