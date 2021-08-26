package main

import (
    "cdht/RoutingModule"
    "cdht/API"
    "cdht/NetworkModule"
    "cdht/Util"
    "time"
    "fmt"
    "math/big"
    "net"
)


func main() {

    // go runFirstNode();

    // go runSecondNode();

    go testApp1();
    go testApp2();

    time.Sleep(time.Minute * 350)
}


func runFirstNode() {
    apiComm := API.ApiCommunication{
        ChannelSize: 100000,
        PORT : "6789",
    }
    apiComm.Init()

    firstNode := RoutingModule.RoutingTable{ 
        RingPort: "9898",
        JumpSpacing: 2,
        FingerTableLength: 4,
        Applications: apiComm.Application,
    }

    // TEST
    currentAppServerPort(&apiComm)
    // END TEST

    firstNode.CreateRing()
    apiComm.NodeRoutingTable = &firstNode

    apiComm.StartAppServer()
}


func runSecondNode() {
    apiComm := API.ApiCommunication{
        ChannelSize: 100000,
        PORT : "3456",
    }
    apiComm.Init()

    secondNode := RoutingModule.RoutingTable{ 
        RemoteNodeAddr: "127.0.0.1:9898",
        NodePort: "3456",
        JumpSpacing: 2,
        FingerTableLength: 4,
        Applications: apiComm.Application,
    }

    // TEST
    currentAppServerPort(&apiComm)
    // END TEST

    secondNode.RunNode()
    apiComm.NodeRoutingTable = &secondNode

    apiComm.StartAppServer()
}





//  # ----------------------- TEST ----------------------- # //

func currentAppServerPort(apiComm *API.ApiCommunication){
	var port string
	fmt.Print("Enter App Server Port: ")
    fmt.Scanln(&port)
	apiComm.PORT = port
}


func testApp1() {
    appConn := NetworkModule.NewNetworkManager("127.0.0.1", "7777")
    fmt.Println("[APP-1]:+.")

    res := appConn.Connect("TCP", func(connection interface{}){
        if connection, ok := connection.(net.Conn); ok { 
            netChannel := NetworkModule.NetworkChannel{ Connection: connection, ChannelSize: 100000 }
            netChannel.Init()
    
            // register app name
            netChannel.SendToSocket(Util.RequestObject{
                AppName: "TEST_APP",
            })        


            fmt.Println("[APP-1]:+ Connected.")

            for {
                select {
                    case netReqObj := <- netChannel.ReqChannel:
				    	fmt.Println("[PACKET-RECEIVED]: ", 4)
                        fmt.Println(netReqObj)
                        
				    default:
                        time.Sleep(time.Millisecond*500)
					    status := netChannel.SendToSocket(Util.RequestObject{
                            Type: "TEST_TYPE_1",
                            RequestID: "REQ_1",
                            AppName: "TEST_APP",
                            AppID: 1,
                            SenderNodeId: big.NewInt(4),
                            ReceiverNodeId: big.NewInt(9),
                            RequestBody: "THIS IS REQ BODY FROM NODE 4",
                        })

                        if !status {
                            fmt.Println("[API] Unable to send")
                        }
			    }
		    }
        }else {
            fmt.Println("[APP-1]:+ Unable to Connect.")
        }
    })

    if !res {
        fmt.Println("[APP-1]:+ Unable to Create Soc.")
    }
}



func testApp2() {
    appConn := NetworkModule.NewNetworkManager("127.0.0.1", "9999")
    fmt.Println("[APP-2]:+.")

    res := appConn.Connect("TCP", func(connection interface{}){
        if connection, ok := connection.(net.Conn); ok { 
            netChannel := NetworkModule.NetworkChannel{ Connection: connection, ChannelSize: 100000 }
            netChannel.Init()

            
            // register app name
            netChannel.SendToSocket(Util.RequestObject{
                AppName: "TEST_APP",
            })


            fmt.Println("[APP-2]:+ Connected.")
    
            for {
                select {
                    case netReqObj := <- netChannel.ReqChannel:
				    	fmt.Println("[PACKET-RECEIVED]", 9)
                        fmt.Println(netReqObj)
                        
				    default:
                        time.Sleep(time.Millisecond*500)
					    status := netChannel.SendToSocket(Util.RequestObject{
                            Type: "TEST_TYPE_2",
                            RequestID: "REQ_2",
                            AppName: "TEST_APP",
                            AppID: 1,
                            SenderNodeId: big.NewInt(9),
                            ReceiverNodeId: big.NewInt(4),
                            RequestBody: "THIS IS REQ BODY FROM NODE 9",
                        })

                        if !status {
                            fmt.Println("[API] Unable to send")
                        }
			    }
		    }
        }else {
            fmt.Println("[APP-2]:+ Unable to Connect.")
        }
    })

    if !res {
        fmt.Println("[APP-1]:+ Unable to Create Soc.")
    }
}