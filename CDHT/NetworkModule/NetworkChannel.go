package NetworkModule

import (
	"cdht/Util"
	"net"
	"fmt"
	"encoding/gob"
	"log"
)


type NetworkChannel struct {
	ReqChannel chan Util.RequestObject
	Connection net.Conn
	ChannelSize int
}


func init(){
	gob.Register( map[string]string{} )
	gob.Register( map[string]interface{}{} )
}


func (netChan *NetworkChannel) Init(){
	netChan.ReqChannel = make(chan Util.RequestObject, netChan.ChannelSize)
	go netChan.readFromSocket()
}



func (netChan *NetworkChannel) readFromSocket(){
	for {
		dec := gob.NewDecoder(netChan.Connection)
		reqObj := &Util.RequestObject{}
	
		if err:= dec.Decode(reqObj); err != nil {
			fmt.Println("[ApiCommunication][Error]: Unable to decode packet.")
			return
		}
		netChan.ReqChannel <- *reqObj
	}
}



func (netChan *NetworkChannel) SendToSocket(reqObj Util.RequestObject) bool {
	enc := gob.NewEncoder(netChan.Connection) 

	if err := enc.Encode(&reqObj); err != nil {
		log.Printf("[NetChannel][Error]: Failed to send packet back to requester: %v ... \n", err)
		return false
	}
	return true
}



func (netChan *NetworkChannel) SendToUDPSocket(address string, port string) {
	var UDPnetworkManager NetworkManager
    UDPnetworkManager.SetIPAddress( address, port)
    UDPnetworkManager.Connect("UDP", func( udpConnection interface{}){
		if soc, ok := udpConnection.(*net.UDPConn); ok { 
			for {
				packet := <- netChan.ReqChannel
		
				if packet.Type == Util.PACKET_TYPE_CLOSE {
					return
				}
				SendUDPPacket(soc, &packet)
			}
		}
	})
}