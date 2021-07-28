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
    "net"  
    "bufio"
)


type NodeData struct {
    ID string
    Node_id string
    IP_address string
    Created_date string
}

func getUserInput() string {
    fmt.Println("Options ")
    fmt.Println("-----------------------------------")
    fmt.Println("[1] for registering your Node")
    fmt.Println("[2] for getting list of Nodes")
    fmt.Println("[3] for pinging a node")
    fmt.Println("-----------------------------------")
    fmt.Print("Enter your option :")
    var userInput string
    fmt.Scanln(&userInput)
    return userInput
}

func main() {

    fmt.Print("Enter Computer IP: ")
    var IpAddress string
    fmt.Scanln(&IpAddress)

    go pingServer(IpAddress)

    for {

        userInput := getUserInput()
        switch userInput {
            case "1":
                go registerNode(IpAddress)
            case "2":
                go gettingRegisteredNodes()
            case "3":
                go pingClient()

        }

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
    fmt.Println("-----------------------------------------")
    
    
    
    var messages []NodeData
    err = json.Unmarshal(body, &messages)
    
    if err != nil {
		fmt.Println("error:", err)
	}

    visited := map[string]bool{}

    for _, node := range messages {
        if _,exitsts := visited[node.IP_address]; !exitsts {
            fmt.Println("[NODE-IP] "+ node.IP_address)
            visited[node.IP_address] = true
        }
    }
}


func registerNode(ipAddress string) {
    nodeID := sha1.New()
    nodeID.Write([]byte( time.Now().String()  ))
    node_id := nodeID.Sum(nil)

    postBody, _ := json.Marshal(map[string]string{
        "Node_id": string( node_id ),
        "IP_address": string( ipAddress ),
    })

    responseBody := bytes.NewBuffer(postBody)
    resp, err := http.Post("https://cdht-monitoring-api.herokuapp.com/nodes", "application/json", responseBody)
    
    if err != nil {
       log.Fatalf("An Error Occured %v", err)
       return
    }
    
    defer resp.Body.Close()
    
    fmt.Println("\n\n[NODE]: Registering Node...")
    fmt.Println("-----------------------------------------")
    fmt.Println("Node registered successfully.")
}



func pingServer(IpAddress string) {
    p := make([]byte, 2048)
    addr := net.UDPAddr{
        Port: 1234,
        IP: net.ParseIP(IpAddress),
    }
    fmt.Printf("[PING] ping server running... \n")
    ser, err := net.ListenUDP("udp", &addr)
    if err != nil {
        fmt.Printf("Some error %v\n", err)
        return
    }
    for {
        _,remoteaddr,err := ser.ReadFromUDP(p)
        fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
        if err !=  nil {
            fmt.Printf("Some error  %v", err)
            continue
        }

        _,err   = ser.WriteToUDP([]byte("From server: Hello I got your message "), remoteaddr)
        if err != nil {
            fmt.Printf("Couldn't send response %v", err)
        }
    }
}



func pingClient(){
    fmt.Print("\n\nEnter Node IP to ping: ")
    var IpAddress string
    fmt.Scanln(&IpAddress)
    fmt.Print("\n")

    p :=  make([]byte, 2048)
    conn, err := net.Dial("udp", IpAddress+ ":1234")
    if err != nil {
        fmt.Printf("Some error %v", err)
        return
    }
    fmt.Fprintf(conn, "Hi UDP Server, How are you doing?")
    _, err = bufio.NewReader(conn).Read(p)
    if err == nil {
        fmt.Printf("%s\n", p)
    } else {
        fmt.Printf("Some error %v\n", err)
    }
    conn.Close()
}