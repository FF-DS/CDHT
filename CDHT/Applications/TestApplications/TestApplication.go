package TestApplications


import (
    "cdht/Util"
	"math/big"
    "time"
    "fmt"
)


type TestApplication struct {
	IPAddress string
	Port string
	
	UDPListenerPort string
	NetChannelSize int
	
	AppName string
	PacketDelay time.Duration
	close_server chan bool

    requestObject Util.RequestObject
}



// # -------------------------- Init environment -------------------------- # //

func (testApp  *TestApplication) Init() TestApplication {
	testApp.close_server = make(chan bool)
	if testApp.IPAddress == "" {
		testApp.IPAddress = "127.0.0.1"
	}
	if testApp.PacketDelay == 0 {
		testApp.PacketDelay = 500
	}
	if testApp.NetChannelSize == 0{
		testApp.NetChannelSize = 10000
	}
	return *testApp
}




// # --------------------- TEST Packet Craft -------------------- # //

func TestPacket() (Util.RequestObject, string, string) {
    var senderNodeId, recvNodeId, appServerPort, UDPListenerPort string
    var ok bool
    reqObject := Util.RequestObject{}
    
	fmt.Println(" ------------------- [TEST-PACKET] ------------------- ")
	fmt.Print("      [+]: Enter App Name: ")
    fmt.Scanln(&reqObject.AppName)

    fmt.Print("      [+]: Enter Sender(Connecting) Application Server Port: ")
    fmt.Scanln(&appServerPort)

	fmt.Print("      [+]: Enter Current (UDP) Server Port: ")
    fmt.Scanln(&UDPListenerPort)
    
    fmt.Print("      [+]: Enter Sender(Connecting) Node Id: ")
    fmt.Scanln(&senderNodeId)
    
    fmt.Print("      [+]: Enter Reciever Node Id: ")
    fmt.Scanln(&recvNodeId)
    fmt.Println(" ------------------------------------------------------ ")

    reqObject.SenderNodeId, ok = new(big.Int).SetString(senderNodeId, 10)
    if !ok { fmt.Println("SetString: error node id") }

    reqObject.ReceiverNodeId, ok = new(big.Int).SetString(recvNodeId, 10)
    if !ok { fmt.Println("SetString: error m") }

    reqObject.RequestBody = "THIS IS REQ BODY FROM NODE: " + senderNodeId + " TO NODE " + recvNodeId

	return reqObject, appServerPort, UDPListenerPort
}




// # --------------- TEST APP RUNNER ---------------------------- #

func RunTestTCPApp(){
    reqObj1, appPort1, _ := TestPacket()
    reqObj2, appPort2, _ := TestPacket()


	app1 := TestApplication{ Port: appPort1, AppName: reqObj1.AppName }
	app1.Init()
	app2 := TestApplication{ Port: appPort2, AppName: reqObj2.AppName }
	app2.Init()

    go app1.TestAppTCP(reqObj1)
    go app2.TestAppTCP(reqObj2)
}


func RunTestUDPApp(){
	reqObj1, appPort1, localPort1 := TestPacket()
    reqObj2, appPort2, localPort2 := TestPacket()


	app1 := &TestApplication{ Port: appPort1, UDPListenerPort: localPort1, AppName: reqObj1.AppName }
	app1.Init()
	app2 := &TestApplication{ Port: appPort2, UDPListenerPort: localPort2, AppName: reqObj2.AppName }
	app2.Init()

    go app1.TestAppUDP(reqObj1)
    go app2.TestAppUDP(reqObj2)
}