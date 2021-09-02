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

    go runTestApp()

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




//  # ----------------------- TEST  APPLICATION ----------------------- # //

func runTestApp(){
    reqObj1 := Util.RequestObject{ Type: "TEST_TYPE_1", RequestID: "REQ_1", AppName: "TEST_APP", AppID: 1}
    appPort1 := testAppAddress("Test App 1", &reqObj1)


    reqObj2 := Util.RequestObject{ Type: "TEST_TYPE_2", RequestID: "REQ_2", AppName: "TEST_APP", AppID: 1}
    appPort2 := testAppAddress("Test App 2", &reqObj2)


    go testApp(reqObj1, appPort1, "Test App 1")
    go testApp(reqObj2, appPort2, "Test App 2")
}


func testApp(reqObj Util.RequestObject, port string, appName string) {
    appConn := NetworkModule.NewNetworkManager("127.0.0.1", port)

    res := appConn.Connect("TCP", func(connection interface{}){
        if connection, ok := connection.(net.Conn); ok { 
            netChannel := NetworkModule.NetworkChannel{ Connection: connection, ChannelSize: 100000 }
            netChannel.Init()
    
            // register app name
            netChannel.SendToSocket(Util.RequestObject{
                AppName: "TEST_APP",
            })        


            fmt.Println("["+appName+"]:+ Connected.")

            for {
                select {
                    case netReqObj := <- netChannel.ReqChannel:
				    	fmt.Printf("[PACKET-RECEIVED][%s][Node-ID][%s]: Packet: %s \n", appName, reqObj.SenderNodeId.String(), netReqObj)
				    default:
                        time.Sleep(time.Millisecond*500)
					    status := netChannel.SendToSocket(reqObj)

                        if !status {
                            fmt.Println("[API] Unable to send")
                        }
			    }
		    }
        }else {
            fmt.Println("["+appName+"]:+ Unable to Connect.")
        }
    })

    if !res {
        fmt.Println("["+appName+"]:+ Unable to Create Soc.")
    }
}




func testAppAddress(appName string, reqObject *Util.RequestObject) string {
    var senderNodeId, recvNodeId, appServerPort string
    var ok bool
    
    fmt.Println("       ",appName)
    fmt.Print("Enter Sender(Connecting) Application Server Port: ")
    fmt.Scanln(&appServerPort)

    fmt.Print("Enter Sender(Connecting) Node Id: ")
    fmt.Scanln(&senderNodeId)

    fmt.Print("Enter Reciever Node Id: ")
    fmt.Scanln(&recvNodeId)


    reqObject.SenderNodeId, ok = new(big.Int).SetString(senderNodeId, 10)
    if !ok { fmt.Println("SetString: error node id") }

    reqObject.ReceiverNodeId, ok = new(big.Int).SetString(recvNodeId, 10)
    if !ok { fmt.Println("SetString: error m") }

    reqObject.RequestBody = "THIS IS REQ BODY FROM NODE: " + senderNodeId + " TO NODE " + recvNodeId

    return appServerPort
}