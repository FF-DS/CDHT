package TestApplications

import (
    "cdht/NetworkModule"
    "cdht/Util"
    "time"
    "fmt"
    "net"
)


func (testApp  *TestApplication) TestAppTCP() {
    appConn := NetworkModule.NewNetworkManager( testApp.IPAddress, testApp.Port)
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
				fmt.Printf("[PACKET-RECEIVED][%s][NODE-ID][%s]: Packet: %s \n", testApp.AppName, netReqObj.ReceiverNodeId.String(), netReqObj)
			default:
				time.Sleep(time.Millisecond*testApp.PacketDelay)

				requestObject := testApp.getPacket()
				if status := netChannel.SendToSocket( requestObject ); !status {
					fmt.Println("[API] Unable to send")
				}
		}
	}
}
