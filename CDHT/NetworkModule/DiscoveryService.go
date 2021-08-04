package NetworkModule

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "log"
    "bytes"
    "encoding/json"
    "cdht/Util"
)


type NodeData struct {
    Node_id string 
    IP_address string 
    Ports map[string]string 
}


func NotifyNodeExistance(nodeInfo Util.NodeInfo) NodeData {
    nodeData := NodeData{
        Node_id : nodeInfo.Node_id.String(), 
        IP_address : nodeInfo.IP_address,
        Ports : nodeInfo.Ports,
    }
    registerNode(nodeData)

    return nodeData
}


func GetRegisteredNodes(currentNodeId string) ([]NodeData) {
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
        if _,exitsts := visited[node.Node_id]; !exitsts && node.Node_id != currentNodeId {
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

    requestBody := bytes.NewBuffer(postBody)
    resp, err := http.Post("https://cdht-monitoring-api.herokuapp.com/nodes", "application/json", requestBody)

    if err != nil {
       log.Fatalf("An Error Occured %v", err)
       return
    }
    
    defer resp.Body.Close()
}
