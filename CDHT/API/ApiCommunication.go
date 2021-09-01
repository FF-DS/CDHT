package API

import (
	"cdht/Util"
	"cdht/RoutingModule"
	"cdht/NetworkModule"
	"net"
	"fmt"
)


type ApiCommunication struct {
	Application map[string](chan Util.RequestObject)
	CloseAppServer chan bool

	NodeRoutingTable *RoutingModule.RoutingTable
	ChannelSize int
	PORT string
}


func (Api *ApiCommunication) Init() {
	Api.CloseAppServer = make(chan bool)
	Api.Application = make(map[string](chan Util.RequestObject))
}


func (Api *ApiCommunication) StartAppServer() {
	netManager := NetworkModule.NewNetworkManager("", Api.PORT)
	go netManager.StartServer("TCP", Api.CloseAppServer, Api.appRequestHandler)
	go netManager.StartServer("UDP", Api.CloseAppServer, Api.appRequestHandler)
	fmt.Println("[APP-SERVICE]:+ Starting Application server at port: ", Api.PORT)
}


func (Api *ApiCommunication) appRequestHandler(connection interface{}) {
	if connection, ok := connection.(net.Conn); ok { 

		netChannel := NetworkModule.NetworkChannel{ Connection: connection, ChannelSize: Api.ChannelSize }
		netChannel.Init()
		
		appInitPackt :=  <-netChannel.ReqChannel
		currReqChannel := make(chan Util.RequestObject, Api.ChannelSize)
		Api.Application[appInitPackt.AppName] = currReqChannel

		for {
			select {
				case netReqObj := <- netChannel.ReqChannel:
					if netReqObj.Type == Util.PACKET_TYPE_CLOSE {
						Api.closeApp( appInitPackt.AppName, netChannel )
						return 
					}else {
						Api.NodeRoutingTable.ForwardPacket( netReqObj )
					}
					
				case chanReqObj := <- currReqChannel:
					netChannel.SendToSocket( chanReqObj )
			}
		}

	}else {
		fmt.Println("[API-SERVER]:+ Unable to Connect.")
	}
}


func (Api *ApiCommunication) closeApp(appName string, netChan NetworkModule.NetworkChannel){
	netChan.Connection.Close()
	delete(Api.Application, appName);
}