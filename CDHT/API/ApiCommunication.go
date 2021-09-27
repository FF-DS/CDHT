package API

import (
	"cdht/Util"
	"cdht/RoutingModule"
	"cdht/NetworkModule"
	"net"
	"fmt"
	"strings"
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
	netManagerTCP := NetworkModule.NewNetworkManager("", Api.PORT)
	go netManagerTCP.StartServer("TCP", Api.CloseAppServer, Api.appRequestHandlerTCP)
	netManagerUDP := NetworkModule.NewNetworkManager("", Api.PORT)
	go netManagerUDP.StartServer("UDP", Api.CloseAppServer, Api.appRequestHandlerUDP)
	// fmt.Println("[APP-SERVICE]:+ Starting Application server at port: ", Api.PORT)
}


func (Api *ApiCommunication) appRequestHandlerTCP(connection interface{}) {
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
						connection.Close()
						delete(Api.Application, netReqObj.AppName);
						return 
					}
						
					if netReqObj.TranslateName {
						netReqObj.ReceiverNodeId = Api.DataToNodeIDTranslator(netReqObj.TranslateNameID)
					}
					if netReqObj.ValidityCheck{
						netReqObj.ValidationCRCchecksum = Api.CalculateCRCchecksum(netReqObj.RequestBody)
					}

					netReqObj.SenderNodeId = Api.NodeRoutingTable.NodeInfo().Node_id
					
					if netReqObj.ReceiverNodeId != nil {
						Api.NodeRoutingTable.ForwardPacket( netReqObj )
					}
										
				case chanReqObj := <- currReqChannel:
					if chanReqObj.ValidityCheck{
						checksumValue := Api.CalculateCRCchecksum(chanReqObj.RequestBody)
						chanReqObj.ValidityCheckResult = checksumValue == chanReqObj.ValidationCRCchecksum
					}
					netChannel.SendToSocket( chanReqObj )
			}
		}

	}else {
		fmt.Println("[API-SERVER]:+ Unable to Connect.")
	}
}



func (Api *ApiCommunication) appRequestHandlerUDP(packetData interface{}) {
	if requestObject, ok := packetData.(Util.RequestObject); ok { 
		if requestObject.Type == Util.PACKET_TYPE_INIT_APP{
			currReqChannel := make(chan Util.RequestObject, Api.ChannelSize)
			Api.Application[requestObject.AppName] = currReqChannel
			
			if port, ok := requestObject.RequestBody.(string); ok {
				netChannel := NetworkModule.NetworkChannel{ ReqChannel: currReqChannel }

				IPAddress := strings.Split(requestObject.UdpAddress, ":")[0]
				go netChannel.SendToUDPSocket( IPAddress, port)
			}
		}else if requestObject.Type == Util.PACKET_TYPE_CLOSE{
			Api.Application[requestObject.AppName] <- Util.RequestObject{ Type : Util.PACKET_TYPE_CLOSE }
			delete(Api.Application, requestObject.AppName);
			return
		}else {
			Api.NodeRoutingTable.ForwardPacket( requestObject )
		}
	}else {
		fmt.Println("[API-SERVER]:+ Unable to decode udp packet.")
	}
}
