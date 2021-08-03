package NetworkModule

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "log"
    "bytes"
    "encoding/json"
    "crypto/sha1"
    "net"  
    "github.com/tkanos/gonfig"
    "math/big"
)


type NodeData struct {
    Node_id string 
    IP_address string 
    Ports map[string]string 
}


func NotifyNodeExistance() NodeData {
    nodeData := getNodeConfig() // get ports
    nodeData.IP_address = getOutboundIP().String() // ip address
    nodeData.Node_id = GetNodeId().String() // node id

    registerNode(nodeData)

    return nodeData
}


func GetRegisteredNodes() ([]NodeData) {
    resp, err := http.Get("https://cdht-monitoring-api.herokuapp.com/nodes")

    if err != nil {
        log.Fatal(err)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
       log.Fatalln(err)
    }
    
    var messages []NodeData
    err = json.Unmarshal(body, &messages)
    
    if err != nil {
		fmt.Println("error:", err)
	}


	// filter out duplicates
    visited := map[string]bool{}

	var nodeList []NodeData 
    for _, node := range messages {
        if _,exitsts := visited[node.Node_id]; !exitsts {
            visited[node.Node_id] = true
			nodeList = append(nodeList, node)
        }
    }

	return nodeList
}


// --------------------- Internal function -------------------- //

func registerNode(nodeData NodeData) {
    postBody, err := json.Marshal(nodeData)

    if err != nil {
        fmt.Println(err)
    }

    responseBody := bytes.NewBuffer(postBody)
    resp, err := http.Post("https://cdht-monitoring-api.herokuapp.com/nodes", "application/json", responseBody)

    if err != nil {
       log.Fatalf("An Error Occured %v", err)
       return
    }
    
    defer resp.Body.Close()
}



// --------------- Node info functions ---------- //

func getNodeConfig() NodeData {
    var nodeData NodeData
    err := gonfig.GetConf("gonfig.json", &nodeData)
    if err != nil {
        panic(err)
    }
    return nodeData
}



func getOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}



// ------------------ Get node id ----------------- //

func GetNodeId() *big.Int {
    nodeIdentification := getOutboundIP().String() + ":" + getNodeConfig().Ports["JOIN"]

    hashFunction := sha1.New()
    hashFunction.Write([]byte(nodeIdentification))
    sha := hashFunction.Sum(nil)

    two, m, hashedID := big.NewInt(2), big.NewInt(160),  (&big.Int{}).SetBytes(sha)
    
    hashedID.SetBytes(sha) 
    modulo := two.Exp( two, m, nil)

    return hashedID.Mod(hashedID, modulo) 
}