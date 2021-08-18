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