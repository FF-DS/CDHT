package main;

import (
    "io/ioutil"
    "fmt"
    "net/http"
    "log"
    "time"
    "bytes"
    "encoding/json"
    "crypto/sha1"
)


type NodeData struct {
    _id string
    node_id string
    iP_address string
    created_date string
}

type Nodes []NodeData 


func main() {

    for {

        go gettingRegisteredNodes()
        go registerNode()

        time.Sleep(time.Minute)
    }
    
}


func gettingRegisteredNodes() {
    resp, err := http.Get("https://cdht-monitoring-api.herokuapp.com/nodes")

    if err != nil {
        log.Fatal(err)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
       log.Fatalln(err)
    }

    fmt.Println("\n\n[NODE]: Receiving Nodes...")
    fmt.Println("-----------------------------------------\n")
    
    
    stringRep := string(body)
    
    messages := []NodeData{} 
    json.Unmarshal(body, &messages)
    fmt.Println(messages)
    fmt.Println(stringRep)
}


func registerNode() {
    nodeID := sha1.New()
    nodeID.Write([]byte( time.Now().String()  ))
    node_id := nodeID.Sum(nil)

    postBody, _ := json.Marshal(map[string]string{
        "Node_id": string( node_id ),
    })

    responseBody := bytes.NewBuffer(postBody)
    resp, err := http.Post("https://cdht-monitoring-api.herokuapp.com/nodes", "application/json", responseBody)
    
    if err != nil {
       log.Fatalf("An Error Occured %v", err)
    }
    
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
       log.Fatalln(err)
    }
    
    fmt.Println("\n\n[NODE]: Registering Node...")
    fmt.Println("-----------------------------------------\n")
    fmt.Println(string(body))
}

