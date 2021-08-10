package main

import (
    "net"
    "fmt"
    "math/big"
    "cdht/NetworkModule"
    "cdht/RoutingModule"
    "cdht/Util"
    "bufio"
    "strings"
    "log"
    "time"
    "strconv"
)



func main() {
    var currNodeInfo Util.NodeInfo
    currNodeInfo.GetNodeInfo()

    M_VAL := testChangeFingerTableID(&currNodeInfo)

    fmt.Println("-----------------Current node Info--------------------")
    fmt.Printf("Node ID    : %s \n", currNodeInfo.Node_id.String())
    fmt.Printf("M          : %s \n", currNodeInfo.M.String())
    fmt.Printf("IP Address : %s \n", currNodeInfo.IP_address)
    fmt.Println("Ports      : ", currNodeInfo.Ports)
    fmt.Println("------------------------------------------------------")
    
    fmt.Println("[TEST]: starting figer fix")

    availableNodeInfo := Util.NodeInfo {
        IP_address : currNodeInfo.IP_address,
        // Ports : map[string]string{ "JOIN_RSP" : "8989", "JOIN_REQ" : "9898",},
        Ports : map[string]string{ "JOIN_RSP" : "6010", "JOIN_REQ" : "2705",},
    }


    fingerTableRoute := RoutingModule.NewFingerTable( currNodeInfo, availableNodeInfo, 2, M_VAL);
    fingerTableRoute.StartServices()

    time.Sleep(time.Second * 10)
    fingerTableRoute.RunFixFingerAlg()


    time.Sleep(time.Minute * 35)
}


func testChangeFingerTableID(nodeInfo *Util.NodeInfo) int{
    var nodeId, m string
    fmt.Print("Enter Node Id: ")
    fmt.Scanln(&nodeId)

    fmt.Print("Enter M: ")
    fmt.Scanln(&m)

    NodeId, ok := new(big.Int).SetString(nodeId, 10)
    if !ok { fmt.Println("SetString: error node id") }

    M, ok := new(big.Int).SetString(m, 10)
    if !ok { fmt.Println("SetString: error m") }

    nodeInfo.Node_id = NodeId
    nodeInfo.M = M

    i, _ := strconv.Atoi(m)
    return i
}

// ## ------------------ TESTS ----------------------- ## 

func main_network_manager_test() {
    fmt.Println("[Network Manager]: Testing")
    
    go udpServer()
    go tcpServer()
    time.Sleep(time.Second * 2)

    go client()
    time.Sleep(time.Minute * 5)
}


func udpServer(){
    close_server := make(chan bool)

    var UDPnetworkManager NetworkModule.NetworkManager
    UDPnetworkManager.SetIPAddress("", "6000")
    UDPnetworkManager.StartServer("UDP", close_server, TEST_udp_server)

}

func tcpServer(){
    close_server := make(chan bool)

    var TCPnetworkManager NetworkModule.NetworkManager
    TCPnetworkManager.SetIPAddress("", "5400")
    TCPnetworkManager.StartServer("TCP", close_server, TEST_tcp_server)

}


func client(){
    fmt.Print("Enter The IP: ")
    var ipAddrr string
    fmt.Scanln(&ipAddrr)

    var UDPnetworkManager NetworkModule.NetworkManager
    UDPnetworkManager.SetIPAddress(ipAddrr, "6000")
    UDPnetworkManager.Connect("UDP", TEST_udp_client)

    var TCPnetworkManager NetworkModule.NetworkManager
    TCPnetworkManager.SetIPAddress(ipAddrr, "5400")
    TCPnetworkManager.Connect("TCP", TEST_tcp_client)
}



// ## ------------------ NETWORK MANAGER CONNECTION TESTS ----------------------- ## 
func TEST_tcp_server(connection interface{}){

    if connection, ok := connection.(net.Conn); ok {
        fmt.Println("[SERVER][TCP]: Connected with client... ")

        clientReader := bufio.NewReader(connection)
        clientRequest, _ := clientReader.ReadString('\n')
        fmt.Println("[SERVER][TCP]:",strings.TrimSpace(clientRequest))

        if _, err := connection.Write([]byte("FROM SERVER.... recieved. \n")); err != nil {
            log.Printf("failed to send the client request: %v\n", err)
        }

    }else{
        fmt.Println("[SERVER][TCP][err]: can't decode")
    }

}


func TEST_tcp_client(connection interface{}){

    if connection, ok := connection.(net.Conn); ok {
        fmt.Println("[CLIENT][TCP]: Connected with server... ")

        if _, err := connection.Write([]byte("HELLO SERVER... \n")); err != nil {
            log.Printf("failed to send the client request: %v\n", err)
        }

        clientReader := bufio.NewReader(connection)
        clientRequest, _ := clientReader.ReadString('\n')

        fmt.Println("[CLIENT][TCP]:",strings.TrimSpace(clientRequest))
    }else{
        fmt.Println("[CLIENT][TCP][err]: can't decode")
    }

}



func TEST_udp_server(packetData interface{}){
    fmt.Println("[SERVER][UDP]: Connected with client... ")

    if packetData, ok := packetData.(NetworkModule.UDPPacketData); ok {
        fmt.Println("[SERVER][UDP]: " + string(packetData.Data), packetData )
    }else{
        fmt.Println("[SERVER][UDP][err]: can't decode")
    }

}



func TEST_udp_client(packetData interface{}){
    fmt.Println("[CLIENT][UDP]: Connected with server... ")

    if packetData, ok := packetData.(NetworkModule.UDPPacketData); ok {
        fmt.Println("[CLIENT][UDP]: connection data")
        fmt.Println("[CLIENT][UDP]:", packetData)
        packetData.UDPsocketConnection.Write( []byte("hello"))
    }else{
        fmt.Println("[CLIENT][UDP][err]: can't decode")
    }
}
