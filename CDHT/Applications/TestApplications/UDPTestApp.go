package TestApplications

import (
    "cdht/NetworkModule"
    "cdht/Util"
    "time"
    "fmt"
    "net"
)

func (testApp  *TestApplication) TestAppUDP() {

    appConn := NetworkModule.NewNetworkManager( testApp.IPAddress, testApp.Port)
    go appConn.Connect("UDP", testApp.sendUDPPacket)

	appServerConn := NetworkModule.NewNetworkManager( "", testApp.UDPListenerPort)
    go appServerConn.StartServer("UDP", testApp.close_server, testApp.receiveUDPPacket)
}


func (testApp  *TestApplication) sendUDPPacket(udpConnection interface{}){
	if soc, ok := udpConnection.(*net.UDPConn); ok { 
		fmt.Printf("[%s]:+ Connected.\n", testApp.AppName)

		appInitPacket := &Util.RequestObject{ AppName: testApp.AppName, Type: Util.PACKET_TYPE_INIT_APP, RequestBody : testApp.UDPListenerPort, }
		NetworkModule.SendUDPPacket(soc, appInitPacket)

		for {
			time.Sleep(time.Millisecond*testApp.PacketDelay)
			requestObject := testApp.getPacket()
			NetworkModule.SendUDPPacket(soc, &requestObject)
		}
	}
}


func (testApp  *TestApplication) receiveUDPPacket(requestObject interface{}){
	if requestObject, ok := requestObject.(Util.RequestObject); ok { 
		fmt.Printf("[PACKET-RECEIVED][%s][NODE-ID][%s]: Packet: %s \n", testApp.AppName, requestObject.ReceiverNodeId.String(), requestObject)
	}
}