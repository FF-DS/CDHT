package NetworkModule

import (
	"net"
	"fmt"
    "encoding/gob"
    "cdht/Util"
)

// Network class

// *** socketConnection - to store tcp connection socket object
// *** ipAddr - to store ip address
// *** conn_type - to store the connection type [TCP,UDP]
// *** port - to store port
// *** status - to store the status of the network object

type NetworkManager struct {
	socketConnection net.Conn
	ipAddr string
	port string
	conn_type string
	status string
}


// UDP Packet Data class

// *** UDPsocketConnection - this will be used to send data from the client place 
// *** Address - to store ip address & port which will be used to identify who connected with the server
// *** Data - since there is no delegated socket type(like for tcp) for UDP in golang we will be storing the sent packet data here. 

type UDPPacketData struct {
	UDPsocketConnection *net.UDPConn
	Address *net.UDPAddr
	Data []byte
}




// ## ---- preparing envt ----- //

func init() {
	fmt.Println("[TEST]: Network Manager")
}


// ## ------------------------------- init ----------------------------- ##

func NewNetworkManager(ipAddr string, port string) NetworkManager {
	var networkMnger NetworkManager
	networkMnger.SetIPAddress(ipAddr, port)

	return networkMnger
}




// ## ------------------------------- public methods ----------------------------- ##

func (network *NetworkManager) SetIPAddress(ipAddr string, port string){
	network.ipAddr = ipAddr
	network.port = port
}

func (network *NetworkManager) GetIPAddress() (ipAddr string, port string){
	ipAddr = network.ipAddr
	port = network.port
	return 
}

func (network *NetworkManager) GetStatus() (string) {
	return network.status
}


func (network *NetworkManager) StartServer(conn_type string, close_server chan bool, handle_request func(interface{})) bool {
	network.conn_type = conn_type
	switch conn_type {
		case "UDP":
			return network.startUdpServer(close_server, handle_request)
		case "TCP":
			return network.startTcpServer(close_server, handle_request)
	}
	return false
}

func (network *NetworkManager) Connect(conn_type string,handle_request func(interface{})) bool {
	network.conn_type = conn_type
	switch conn_type {
		case "UDP":
			return network.udpConnection(handle_request)
		case "TCP":
			return network.tcpConnection(handle_request)
	}
	return false
}


func (network *NetworkManager) CloseConn() bool {
	network.socketConnection.Close()
	network.status = "CLOSED"
	return true;
}



// ---------- methods only for tcp connection object ------------------ //

func (network *NetworkManager) CreateTCPConnection() bool{
	socketConnection, err := net.Dial("tcp4", fmt.Sprintf("%s:%s", network.ipAddr, network.port) ) 
	if err != nil {
		panic(err)
		return false;
	}
	network.socketConnection = socketConnection
	return true;
}


func (network *NetworkManager) SendPacket(packet Util.FingerTablePacket) bool {
	enc := gob.NewEncoder(network.socketConnection) 

	if err := enc.Encode(&packet); err != nil {
		return false;
	}
	return true
}


func (network *NetworkManager) RecievePacket() Util.FingerTablePacket {
	dec := gob.NewDecoder(network.socketConnection)
	packet := &Util.FingerTablePacket{}

	if err := dec.Decode(packet); err != nil {
		return *packet
	}
	return *packet
}

// ---------- [end] methods only for tcp connection object ------------------ //




// ## -------------------------- internal functions ---------------------------- ##

// *** This function will accept a channel which will control whether the server should close or connect &
// *** a handler function which will be called with the packet data (the packet data struct) 
// *** 
func (network *NetworkManager) startUdpServer(close_server chan bool, handle_request func(interface{}) )  bool {
	localAddress, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%s", network.ipAddr, network.port))
	// localAddress, err := net.ResolveUDPAddr("udp",  ":"+strconv.Itoa(network.port))
	network.status = "LISTENING"

	if err != nil {
		panic(err)
		return false
	}
	
	connListner, err := net.ListenUDP("udp4", localAddress)
	if err != nil {
		panic(err)
		return false
	}
	
	defer connListner.Close()

	if err != nil {
		panic(err)
	}

    for {
        select {
			case <-close_server:
				return true
			default:
				inputBytes := make([]byte, 4096)
				length, add, err := connListner.ReadFromUDP(inputBytes)

				if err != nil {
					panic(err)
				}

				go handle_request( UDPPacketData{ Data:inputBytes[:length],  Address: add } )
		}
    }
}


// *** This function will accept a channel which will control whether the server should close or connect &
// *** a handler function which will be called with the tcp connection socket, then the function will be respossible
// *** for reading data from the socket, for maintaing the connection or sending data via the socket 
// *** 
func (network *NetworkManager) startTcpServer(close_server chan bool, handle_request func(interface{}) )  bool {
	connListner, err := net.Listen("tcp4", fmt.Sprintf("%s:%s", network.ipAddr, network.port) )
	// connListner, err := net.Listen("tcp", ":"+strconv.Itoa(network.port) )
	network.status = "LISTENING"

	if err != nil {
		panic(err)
	}

	defer connListner.Close()

	for {

		select {
			case <-close_server:
				return true
			default:
				conn, err := connListner.Accept()
				if err != nil {
					panic(err)
					return false
				}
				go handle_request(conn)
		}
	}
}


// *** This function will accept handler which will be used to forward the tcp connection object 
// *** the handler function will be respossible for reading data from the socket, for maintaing the connection or sending data via the socket 
// *** 

func (network *NetworkManager) tcpConnection(handle_request func(interface{}))  bool {
	socketConnection, err := net.Dial("tcp4", fmt.Sprintf("%s:%s", network.ipAddr, network.port) ) 
	if err != nil {
		panic(err)
		return false;
	}
	network.status = "CONNECTED"
	network.socketConnection = socketConnection
	go handle_request(network.socketConnection)
	return true
}


// *** This function will accept handler which will be used to forward the udp pakcet struct data
// *** the handler function will be respossible for reading data from the socket, for maintaing the connection or sending data via the socket 
// *** 

func (network *NetworkManager) udpConnection(handle_request func(interface{}) )  bool {
	connServerAddress, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%s", network.ipAddr, network.port) )
	socketConnection, err := net.DialUDP("udp4", nil, connServerAddress)
	network.status = "CONNECTED"

    if err != nil {
		panic(err)
        return false
    }
	go handle_request(  UDPPacketData{ UDPsocketConnection: socketConnection,  Address: connServerAddress }  )
	return true
}