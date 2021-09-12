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

    requestObject *Util.RequestObject
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

func (testApp  *TestApplication) TestPacket() (Util.RequestObject, bool) {
    var recvNodeId, repeat, payload string
	var ok = false

    reqObject := Util.RequestObject{ Type: Util.PACKET_TYPE_APPLICATION, RequestID: "12221", AppName: testApp.AppName , AppID: 1221}
    
	fmt.Println(" ------------------- [TEST-PACKET] ------------------- ")
    fmt.Print("      [+]: Enter Reciever Node Id: ")
    fmt.Scanln(&recvNodeId)
	fmt.Print("      [+]: Enter Packet PAYLOAD : ")
    fmt.Scanln(&payload)
	fmt.Print("      [+]: Use this packet repeatedly (y/N) : ")
    fmt.Scanln(&repeat)
    fmt.Println(" ------------------------------------------------------ ")

    reqObject.ReceiverNodeId, ok = new(big.Int).SetString(recvNodeId, 10)
    for !ok { 
		fmt.Print("      [+]: Enter (Valid) Reciever Node Id: ")
		fmt.Scanln(&recvNodeId)
		reqObject.ReceiverNodeId, ok = new(big.Int).SetString(recvNodeId, 10)

	}

	reqObject.RequestBody = payload
	if reqObject.RequestBody == "" {
		reqObject.RequestBody = "THIS IS REQ BODY IS FOR NODE " + recvNodeId
	}else
	
	if repeat == "y" || repeat == "yes" {
		return reqObject, true
	}
	return reqObject, false
}


func (testApp  *TestApplication) getPacket() Util.RequestObject {
	if testApp.requestObject == nil {
		requestObject, repeat := testApp.TestPacket()
		if repeat {
			testApp.requestObject = &requestObject 
		}

		return requestObject
	}
	return	*testApp.requestObject
}