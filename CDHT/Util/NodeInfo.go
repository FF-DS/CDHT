package Util

import (
    "log"
    "crypto/sha1"
    "net"  
    "github.com/tkanos/gonfig"
    "math/big"
)


// # --------------------------- NODE INFO --------------------------- # //

type NodeInfo struct {
	Node_id *big.Int 
    IP_address string 
    Ports map[string]string 
	M *big.Int 
}


func (nodeInfo *NodeInfo) GetNodeInfo() NodeInfo {
	nodeInfo.getNodeConfig();
	// nodeInfo.setOutboundIP();
    nodeInfo.IP_address = "127.0.0.1"
	nodeInfo.generateNodeId();
	return *nodeInfo
}


func (nodeInfo *NodeInfo) generateNodeId() {	
    nodeIdentification := nodeInfo.IP_address + ":" + nodeInfo.Ports["JOIN_REQ"]

    hashFunction := sha1.New()
    hashFunction.Write([]byte(nodeIdentification))
    sha := hashFunction.Sum(nil)

    two, m, hashedID := big.NewInt(2), big.NewInt(160),  (&big.Int{}).SetBytes(sha)
    
    // hashedID.SetBytes(sha) 
    modulo := two.Exp( two, m, nil)

    nodeInfo.Node_id = hashedID.Mod(hashedID, modulo)
    nodeInfo.M = m
}

func (nodeInfo *NodeInfo)  setOutboundIP() {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }

    defer conn.Close()
    localAddr := conn.LocalAddr().(*net.UDPAddr)

    nodeInfo.IP_address = localAddr.IP.String()
}


func (nodeInfo *NodeInfo)  getNodeConfig() {
    err := gonfig.GetConf("gonfig.json", nodeInfo)
    if err != nil {
        panic(err)
    }
}


