package TestApplications

import (
    "cdht/NetworkModule"
    "cdht/Util"
    "time"
    "fmt"
    "net"
)


func (testApp  *TestApplication) TestAppTCP(reqObj Util.RequestObject) {
    appConn := NetworkModule.NewNetworkManager( testApp.IPAddress, testApp.Port)
	testApp.requestObject = reqObj
	appConn.Connect("TCP", testApp.handleTCPConnection)
}


func (testApp  *TestApplication) handleTCPConnection(connection interface{}){
	if connection, ok := connection.(net.Conn); ok { 
		netChannel := NetworkModule.NetworkChannel{ Connection: connection, ChannelSize: testApp.NetChannelSize }
		netChannel.Init()

		// register app name
		netChannel.SendToSocket(Util.RequestObject{  AppName: testApp.AppName, Type: Util.PACKET_TYPE_INIT_APP,})        

		fmt.Printf("[%s]:+ Connected.\n", testApp.AppName)
		testApp.handlePacketTransaction(netChannel)
	}else {
		fmt.Printf("[%s]:+ Unable to Connect.\n", testApp.AppName)
	}
}


func (testApp  *TestApplication) handlePacketTransaction(netChannel NetworkModule.NetworkChannel) {
	for {
		select {
			case netReqObj := <- netChannel.ReqChannel:
				fmt.Printf("[PACKET-RECEIVED][%s][Node-ID][%s]: Packet: %s \n", testApp.AppName, testApp.requestObject.SenderNodeId.String(), netReqObj)
			default:
				time.Sleep(time.Millisecond*testApp.PacketDelay)
				if status := netChannel.SendToSocket( testApp.requestObject ); !status {
					fmt.Println("[API] Unable to send")
				}
		}
	}
}