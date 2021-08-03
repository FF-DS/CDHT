package main;

import (
    "net"
    "fmt"
    "cdht/NetworkModule"
    "bufio"
    "strings"
    "log"
    "time"
)



func main() {
    // fmt.Println(NetworkModule.NotifyNodeExistance())
    NetworkModule.NotifyNodeExistance()

    time.Sleep(time.Second * 3)
    fmt.Println(NetworkModule.GetRegisteredNodes())
}


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
    UDPnetworkManager.SetIPAddress("", 6000)
    UDPnetworkManager.StartServer("UDP", close_server, TEST_udp_server)

}

func tcpServer(){
    close_server := make(chan bool)

    var TCPnetworkManager NetworkModule.NetworkManager
    TCPnetworkManager.SetIPAddress("", 5400)
    TCPnetworkManager.StartServer("TCP", close_server, TEST_tcp_server)

}

//10.6.152.221

func client(){
    fmt.Print("Enter The IP: ")
    var ipAddrr string
    fmt.Scanln(&ipAddrr)

    var UDPnetworkManager NetworkModule.NetworkManager
    UDPnetworkManager.SetIPAddress(ipAddrr, 6000)
    UDPnetworkManager.Connect("UDP", TEST_udp_client)

    var TCPnetworkManager NetworkModule.NetworkManager
    TCPnetworkManager.SetIPAddress(ipAddrr, 5400)
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
