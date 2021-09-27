package DNS


import (
    "poc/app/DNS/NetworkModule"
    "poc/app/DNS/Util"
    "fmt"
    "net"
)

type Connection struct {
	DnsApplicationResponseChannel chan Util.RequestObject
	DnsApplicationRequestChannel chan Util.AppCommand

	IPAddress string
	Port string
	
	UDPListenerPort string
	NetChannelSize int
	
	AppName string
	close_server chan bool
}


func (conn  *Connection) Init() *Connection {
	conn.close_server = make(chan bool)
	conn.AppName = "DNS_APP"
	conn.StartTCPconnection()
	return conn
}


func (conn  *Connection) StartTCPconnection() {
    appConn := NetworkModule.NewNetworkManager( conn.IPAddress, conn.Port)
	appConn.Connect("TCP", conn.handleTCPConnection)
}


func (conn  *Connection) handleTCPConnection(connection interface{}){
	if connection, ok := connection.(net.Conn); ok { 
		netChannel := NetworkModule.NetworkChannel{ Connection: connection, ChannelSize: conn.NetChannelSize }
		netChannel.Init()

		// register app name
		netChannel.SendToSocket(Util.RequestObject{  AppName: conn.AppName, Type: Util.PACKET_TYPE_INIT_APP,})        

		fmt.Printf("[%s]:+ Connected.\n", conn.AppName)
		conn.handlePacketTransaction(netChannel)
	}else {
		fmt.Printf("[%s]:+ Unable to Connect.\n", conn.AppName)
	}
}


func (conn  *Connection) handlePacketTransaction(netChannel NetworkModule.NetworkChannel) {
	for {
		select {
			case netReqObj := <- netChannel.ReqChannel:
				conn.DnsApplicationResponseChannel <- netReqObj

			case apiReqObj := <-conn.DnsApplicationRequestChannel:
				if status := netChannel.SendToSocket( conn.constructRequestObject( apiReqObj ) ); !status {
					fmt.Println("[DNS-APP] Unable to send")
				}
		}
	}
}


func (conn  *Connection) constructRequestObject(command Util.AppCommand) Util.RequestObject {
    return Util.RequestObject{ 
		Type: Util.PACKET_TYPE_APPLICATION, RequestID: "1221", AppName: conn.AppName , AppID: 1221,
		TranslateName: true , TranslateNameID: []byte(command.RecordData.RecordType +":"+ command.RecordData.RecordKey), 
		SendToReplicas: command.SendToAll, RequestBody: command.ToMap(), ValidityCheck: command.DoValidityCheck,
	}
}