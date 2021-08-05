package RoutingModule

import (
	"cdht/Util"
	"cdht/NetworkModule"
    "encoding/gob"
	"math/big"
	"fmt"
	"time"
	"net"
	"log"
	"strconv"
	"sort"
)

type FingerTableEntry struct {
	routeChannel (chan Util.FingerTablePacket)
	routeNodeId *big.Int
}


type FingerTableRoute struct {
	routeConn map[string]FingerTableEntry
	currentNodeInfo Util.NodeInfo
	
	jumpSpacing int
	fingerTableLength int

	routeToClosestNodeChannel (chan Util.FingerTablePacket)
	routeToAnyAvailibleNodeChannel (chan Util.FingerTablePacket) 
	resolvePacketChannel (chan Util.FingerTablePacket) 
	
	closeJoinResponseHandlerServerChannel (chan bool) 
}




// #------------------------------ Init ------------------------# //


func NewFingerTable( currentNodeInfo Util.NodeInfo, jumpSpacing int, fingerTableLength int) FingerTableRoute {

	fingerTableRoute := FingerTableRoute{
		currentNodeInfo : currentNodeInfo,
		jumpSpacing : jumpSpacing,
		fingerTableLength : fingerTableLength,
		routeConn : make(map[string]FingerTableEntry),

		routeToClosestNodeChannel : make(chan Util.FingerTablePacket, 100),
		routeToAnyAvailibleNodeChannel :  make(chan Util.FingerTablePacket, 100),
		resolvePacketChannel : make(chan Util.FingerTablePacket, 100),
		
		closeJoinResponseHandlerServerChannel :  make(chan bool, 5),
	}


	return fingerTableRoute
}


func (fingerTableRoute *FingerTableRoute) StartServices() {
	fmt.Println("[SERVICES]:  starting services..... ")
	go fingerTableRoute.routeToClosestNode()
	go fingerTableRoute.routeToAnyAvailibleNode()
	go fingerTableRoute.joinResponseHandler()
	go fingerTableRoute.availableServerRequestHandler()
	go fingerTableRoute.resolvePacketRequest()
	go fingerTableRoute.registerNodeAsAvaiable() 

}

// #------------------------------ Public Functions ------------------------# //

func (fingerTableRoute *FingerTableRoute) GetFingerTableState() (map[string]FingerTableEntry) {
	return fingerTableRoute.routeConn
}


func (fingerTableRoute *FingerTableRoute) GetJumpSpacing() int {
	return fingerTableRoute.jumpSpacing
}


func (fingerTableRoute *FingerTableRoute) SetJumpSpacing(jumpSpace int) {
	fingerTableRoute.jumpSpacing = jumpSpace
}



func (fingerTableRoute *FingerTableRoute) RunFixFingerAlg() {
	for {
		fmt.Println("[FixFInger]:  starting fixfinger..... ")
		fingerTableRoute.fixFingerAlg()
		time.Sleep(time.Second * 10)
	}
}


// #------------------------------ Internal Functions ------------------------# //


// ************ finger fix functions ************


func (fingerTableRoute *FingerTableRoute) fixFingerAlg() {

	for id := 0; id < fingerTableRoute.fingerTableLength; id++ {
		findNodeId := fingerTableRoute.calculateNextFingerTableEntry(id)

		joinReqPacket :=  Util.FingerTablePacket{
			Type : "JOIN_REQ",
			SenderIp : fingerTableRoute.currentNodeInfo.IP_address,
			SenderNodeId : fingerTableRoute.currentNodeInfo.Node_id,
			ReceiverNodeId : findNodeId,
			FingerTableID : findNodeId,
			Ports : fingerTableRoute.currentNodeInfo.Ports,
		}

		fmt.Println("[FixFInger]:  looking for %s..... ",findNodeId.String())
		fmt.Println(joinReqPacket)


		if len( fingerTableRoute.routeConn) != 0 {

			fmt.Println("[FixFInger]:  looking for closest node")
			fingerTableRoute.routeToClosestNodeChannel <- joinReqPacket
		}else{

			fmt.Println("[FixFInger]:  looking for any node" )
			fingerTableRoute.routeToAnyAvailibleNodeChannel <- joinReqPacket
		}
	}
 
}


// ************ route functions ************

// ### [Service]
func (fingerTableRoute *FingerTableRoute) routeToClosestNode() {
	fmt.Println("[SERVICES]:  starting [routeToClosestNode] services.....\n[SERVICES]:  listning for packets..... ")
	for {

		packet :=  <-fingerTableRoute.routeToClosestNodeChannel

		minDistance := &big.Int{}
		minDistance.Exp(big.NewInt( int64(fingerTableRoute.jumpSpacing) ), fingerTableRoute.currentNodeInfo.M , nil)

		currChoosenNodeID := ""
		
		sortedMapKeys := fingerTableRoute.sortFingerTableKeys()

		for _, currNodeId := range sortedMapKeys {
			currNode_id, _ := new(big.Int).SetString(currNodeId, 10)
			currDistance := calculateManhanttanDistance(packet.ReceiverNodeId, currNode_id)

			if currDistance.Cmp(minDistance) == -1  &&  currDistance.Cmp( big.NewInt(int64(0)) ) > 0 {
				minDistance = currDistance
				currChoosenNodeID = currNodeId
			}
		}


		if fingerTableRoute.currentNodeInfo.Node_id != packet.SenderNodeId  &&  calculateManhanttanDistance(packet.ReceiverNodeId, fingerTableRoute.currentNodeInfo.Node_id).Cmp(minDistance) <= 0 {
			fingerTableRoute.resolvePacketChannel <- packet
		}

		
		fingerTableRoute.routeConn[currChoosenNodeID].routeChannel <- packet
	}
}


// ### [Service]
func (fingerTableRoute *FingerTableRoute) routeToAnyAvailibleNode(){
	fmt.Println("[SERVICES]:  starting [routeToAnyAvailibleNode] services.....\n[SERVICES]:  listning for packets..... ")
	for {

		availibleNodes := NetworkModule.GetRegisteredNodes()
		if len(availibleNodes) == 0 {
			fmt.Println("[FINGER TABLE ROUTE]: No availible node, sleeping for 10sec....")
			time.Sleep(time.Second * 10)
			continue
		}


		joinReqPacket :=  <-fingerTableRoute.routeToAnyAvailibleNodeChannel
		availibleNode := availibleNodes[0]

		var connAvailble NetworkModule.NetworkManager

		Port,_ := strconv.Atoi( availibleNode.Ports["JOIN_REQ"]) 

		connAvailble.SetIPAddress(availibleNode.IP_address, Port )
		connAvailble.Connect("TCP", func(connection interface{}){

			if connection, ok := connection.(net.Conn); ok { 

				enc := gob.NewEncoder(connection) 

				if err := enc.Encode(&joinReqPacket); err != nil {
					log.Printf("[FINGER TABLE ROUTE]: Failed to send the join request: %v ... \n", err)
				}
				connection.Close()

			}else{
				fmt.Println("[FINGER TABLE ROUTE]: Can't decode the connection socket...")
			}
			
		})
	}
}

// ************************ end ************************




// ************************ join handler functions ************************


// ### [Service]
func (fingerTableRoute *FingerTableRoute) joinResponseHandler() {
	fmt.Println("[SERVICES]:  starting [joinResponseHandler] services.....\n[SERVICES]:  listning for packets..... ")

	var joinResponseServer NetworkModule.NetworkManager
	Port, _  := strconv.Atoi( fingerTableRoute.currentNodeInfo.Ports["JOIN_RESP"] )
	joinResponseServer.SetIPAddress( fingerTableRoute.currentNodeInfo.IP_address, Port )

	joinResponseServer.StartServer("TCP", fingerTableRoute.closeJoinResponseHandlerServerChannel, fingerTableRoute.fingerTableConnectionHandler)

}


// ### [Service]
func (fingerTableRoute *FingerTableRoute) availableServerRequestHandler() {
	fmt.Println("[SERVICES]:  starting [availableServerRequestHandler] services.....\n[SERVICES]:  listning for packets..... ")

	var joinResponseServer NetworkModule.NetworkManager
	Port, _  := strconv.Atoi( fingerTableRoute.currentNodeInfo.Ports["JOIN_REQ"] )
	joinResponseServer.SetIPAddress( fingerTableRoute.currentNodeInfo.IP_address, Port )

	joinResponseServer.StartServer("TCP", fingerTableRoute.closeJoinResponseHandlerServerChannel, func(connection interface{}){
			if connection, ok := connection.(net.Conn); ok { 
				 
				dec := gob.NewDecoder(connection)
				packet := &Util.FingerTablePacket{}
				dec.Decode(packet)

				fingerTableRoute.routeToClosestNodeChannel <- *packet

			}else{
				fmt.Println("[FINGER TABLE ROUTE]: Can't decode the connection socket...")
			}
	})

}


// ### [Service] notify c&c it's existance
func (fingerTableRoute *FingerTableRoute) registerNodeAsAvaiable() {
	for {
		if len(fingerTableRoute.routeConn) > 0 {
			NetworkModule.NotifyNodeExistance( fingerTableRoute.currentNodeInfo )
			fmt.Println("[C&C]: Registering node to c&c server...")
		}

		time.Sleep(time.Minute * 3)
	}
}


// ************************ finger Table conn function functions ************************

func (fingerTableRoute *FingerTableRoute) fingerTableConnectionHandler(connection interface{}) {
	fmt.Println("[NODE]:  node connection created [fingerTableConnectionHandler] starting.....\n[NODE]:  listning for packets..... ")

	if connection, ok := connection.(net.Conn); ok { 
		
		readPacketFromSocketChannel := make(chan Util.FingerTablePacket,100)
		readPacketFromCurrNodeChannel := make(chan Util.FingerTablePacket,100)

		go readPacketFromSocket(connection, readPacketFromSocketChannel)

		firstJoinReqPacket :=  <-readPacketFromSocketChannel
		fingerTableRoute.routeConn[ firstJoinReqPacket.FingerTableID.String() ].routeChannel <- Util.FingerTablePacket{ Type: "CLOSE_CONN"  } 

		fingerTableRoute.routeConn[ firstJoinReqPacket.FingerTableID.String() ] = FingerTableEntry{ readPacketFromCurrNodeChannel, firstJoinReqPacket.SenderNodeId }



		for {
			select {
				case packet := <-readPacketFromCurrNodeChannel:
					if packet.Type == "CLOSE_CONN" {
						connection.Close()
						return
					}else{
						sendPacketToSocket(connection, packet)
					}	

				case packet := <-readPacketFromSocketChannel:
					// if packet.Type == "PING_REQ" {
					// 	resolvePingReq(packet)
					// }else 
					if packet.Type == "JOIN_REQ" {
						fingerTableRoute.routeToClosestNodeChannel <- packet			
					}
			}	
		}

	}else{
		fmt.Println("[FINGER TABLE ROUTE]: Can't decode the connection socket...")
	}
}



 
// ************************ resolve packet handler functions ************************

// ### [Service]
func  (fingerTableRoute *FingerTableRoute) resolvePacketRequest(){
	fmt.Println("[SERVICES]:  starting [resolvePacketRequest] services.....\n[SERVICES]:  listning for packets..... ")
	for {
		packetReq:= <-fingerTableRoute.resolvePacketChannel

		var connToServer NetworkModule.NetworkManager


		joinRespPacket :=  Util.FingerTablePacket{
			Type : "JOIN_RESP",
			SenderIp : fingerTableRoute.currentNodeInfo.IP_address,
			ReceiverIp : packetReq.SenderIp,
			SenderNodeId : fingerTableRoute.currentNodeInfo.Node_id,
			ReceiverNodeId : packetReq.SenderNodeId,
			FingerTableID : packetReq.FingerTableID,
			Ports : fingerTableRoute.currentNodeInfo.Ports,
		}


		Port, _ :=  strconv.Atoi( packetReq.Ports["JOIN_RESP"] )

		connToServer.SetIPAddress(packetReq.SenderIp, Port)
		connToServer.Connect("TCP", func(connection interface{}){

			if connection, ok := connection.(net.Conn); ok { 

				enc := gob.NewEncoder(connection) 

				if err := enc.Encode(&joinRespPacket); err != nil {
					log.Printf("[FINGER TABLE ROUTE]: Failed to send the join request: %v ... \n", err)
				}
				connection.Close()

			}else{
				fmt.Println("[FINGER TABLE ROUTE]: Can't decode the connection socket...")
			}
			
		})

	}
}





// #------------------------------ helper functions ------------------------# //


func (fingerTableRoute *FingerTableRoute) calculateNextFingerTableEntry(i int)  *big.Int {
	offset := (&big.Int{}).Exp( big.NewInt( int64(fingerTableRoute.jumpSpacing) ), big.NewInt(int64(i)), nil)
	// Sum
	sum := (&big.Int{}).Add( fingerTableRoute.currentNodeInfo.Node_id, offset)
	// Get the ceiling
	ceil := (&big.Int{}).Exp(  big.NewInt( int64(fingerTableRoute.jumpSpacing) ), fingerTableRoute.currentNodeInfo.M, nil)

	return (&big.Int{}).Mod(sum, ceil)
}


func (fingerTableRoute *FingerTableRoute) sortFingerTableKeys() []string {
	strKeys := []string{}
	
	for key,_ := range fingerTableRoute.routeConn  {
		strKeys = append(strKeys, key )
	}

	sort.SliceStable(strKeys, func(index_1, index_2 int) bool {
		nodeIdOne, _ :=  new(big.Int).SetString(strKeys[index_1], 10)
		nodeIdTwo, _ := new(big.Int).SetString(strKeys[index_2], 10)
		return nodeIdOne.Cmp(nodeIdTwo) < 0 
	})

	return strKeys
}



func calculateManhanttanDistance(nodeIdOne *big.Int, nodeIdTwo *big.Int) *big.Int {
	return nodeIdTwo.Sub(nodeIdOne, nodeIdTwo)
}


func readPacketFromSocket(connection net.Conn, sentPackets chan Util.FingerTablePacket) {
	for {

		dec := gob.NewDecoder(connection)
		packet := &Util.FingerTablePacket{}
		dec.Decode(packet)

		sentPackets <- *packet
	}
}


func sendPacketToSocket(connection net.Conn, packet Util.FingerTablePacket) {
	enc := gob.NewEncoder(connection) 

	if err := enc.Encode(&packet); err != nil {
		log.Printf("[FINGER TABLE ROUTE]: Failed to send the join request: %v ... \n", err)
	}
}

// # ************************ end ************************ # //

