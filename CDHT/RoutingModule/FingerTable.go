package RoutingModule

import (
	"cdht/Util"
	"cdht/NetworkModule"
    "encoding/gob"
	"math/big"
	"fmt"
	"net"
	"log"
	"sort"
	"time"
)


type FingerTableRoute struct {
	currentNodeInfo Util.NodeInfo
	availableNodeInfo Util.NodeInfo
	
	jumpSpacing int
	fingerTableLength int

	routeConn map[string]TableEntry
	closeRequestServer (chan bool) 
	
	routeToClosestNodeChannel (chan Util.FingerTablePacket) 
	resolvePacketChannel (chan Util.FingerTablePacket) 
<<<<<<< HEAD
	availableNodeExist bool
=======
>>>>>>> f50043c628cbbc2aa8330621db2149f78ff4443e
}



// #------------------------------ Init ------------------------# //

<<<<<<< HEAD
func CreateRing(currentNodeInfo Util.NodeInfo, jumpSpacing int, fingerTableLength int) FingerTableRoute {
	fingerTableRoute := FingerTableRoute{
		currentNodeInfo : currentNodeInfo,
		availableNodeExist : false,

		jumpSpacing : jumpSpacing,
		fingerTableLength : fingerTableLength,
		
		routeConn : make(map[string]TableEntry),
		closeRequestServer :  make(chan bool),

		routeToClosestNodeChannel: make(chan Util.FingerTablePacket, 100),
		resolvePacketChannel: make(chan Util.FingerTablePacket, 100), 

	}

	return fingerTableRoute
}



=======
>>>>>>> f50043c628cbbc2aa8330621db2149f78ff4443e
func NewFingerTable( currentNodeInfo Util.NodeInfo, availableNodeInfo Util.NodeInfo, jumpSpacing int, fingerTableLength int) FingerTableRoute {

	fingerTableRoute := FingerTableRoute{
		currentNodeInfo : currentNodeInfo,
		availableNodeInfo : availableNodeInfo,
<<<<<<< HEAD
		availableNodeExist : true,
=======
>>>>>>> f50043c628cbbc2aa8330621db2149f78ff4443e

		jumpSpacing : jumpSpacing,
		fingerTableLength : fingerTableLength,
		
		routeConn : make(map[string]TableEntry),
		closeRequestServer :  make(chan bool),

		routeToClosestNodeChannel: make(chan Util.FingerTablePacket, 100),
		resolvePacketChannel: make(chan Util.FingerTablePacket, 100), 
	}

	return fingerTableRoute
}


func (fingerTableRoute *FingerTableRoute) InitFingerService() {
	fmt.Println("[SERVICES]:  starting services..... ")

	go fingerTableRoute.joinRespListnerService()
	go fingerTableRoute.requestListnerService()
	go fingerTableRoute.resolvePacketService()
	go fingerTableRoute.routeToClosestNodeService()
}

// # ******************************** end *********************************** # //






// #------------------------------ Public Functions ------------------------# //

func (fingerTableRoute *FingerTableRoute) GetFingerTableState() (map[string]TableEntry) {
	return fingerTableRoute.routeConn
}


func (fingerTableRoute *FingerTableRoute) GetJumpSpacing() int {
	return fingerTableRoute.jumpSpacing
}


func (fingerTableRoute *FingerTableRoute) SetJumpSpacing(jumpSpace int) {
	fingerTableRoute.jumpSpacing = jumpSpace
}


<<<<<<< HEAD
func (fingerTableRoute *FingerTableRoute) RunFixFingerAlg( loop bool) {
=======
func (fingerTableRoute *FingerTableRoute) RunFixFingerAlg() {
>>>>>>> f50043c628cbbc2aa8330621db2149f78ff4443e
	for {
		time.Sleep(time.Second * 10)

		fmt.Println("[FigerFix]: started....")

<<<<<<< HEAD
		// if len(fingerTableRoute.routeConn) != 0 {
		// 	fmt.Println("[FigerFix]: currently there is an existing finger table... ")
		// }
		fingerTableRoute.fixFingerAlg()
		fmt.Printf("-----------------------------NODE [%s]---------------------------\n", fingerTableRoute.currentNodeInfo.Node_id)
		fmt.Println("Finger Routing Table:")
		for id, entry := range fingerTableRoute.routeConn {
			fmt.Printf("		Finger Entry : %s | Node Id : %s  IP_ADD : %s  REQ_PORT : %s\n",id, entry.CurrNodeInfo.Node_id, entry.CurrNodeInfo.IP_address, entry.CurrNodeInfo.Ports["JOIN_REQ"])
		}
		fmt.Println("------------------------------------------------------------------")
		

		if !loop {
			return
		}
=======
		if len(fingerTableRoute.routeConn) != 0 {
			fmt.Println("[FigerFix]: currently there is an existing finger table... ")
		}

		fingerTableRoute.fixFingerAlg()
	
		fmt.Println("Routing Table:")
		for id, entry := range fingerTableRoute.routeConn {
			fmt.Printf(" Finger Entry : %s | Node Id : %s  IP_ADD : %s  REQ_PORT : %s\n",id, entry.CurrNodeInfo.Node_id, entry.CurrNodeInfo.IP_address, entry.CurrNodeInfo.Ports["JOIN_REQ"])
		}
	
>>>>>>> f50043c628cbbc2aa8330621db2149f78ff4443e
	}
}



// [TEST-FUNC]
func (fingerTableRoute *FingerTableRoute) SendTestPacket(packet Util.FingerTablePacket){
	fingerTableRoute.resolvePacketChannel <- packet;
}

// # ******************************** end ****************************************** # //






// #--------------------------------- Internal Functions ----------------------------# //

// [FIX FINGER]
func (fingerTableRoute *FingerTableRoute) fixFingerAlg() {
	for id := 0; id < fingerTableRoute.fingerTableLength; id++ {
		findNodeId := fingerTableRoute.calculateNextFingerId(id)

<<<<<<< HEAD
		joinReqPacket :=  Util.FingerTablePacket{ Type : "JOIN_REQ", FingerTableID : findNodeId, SenderNodeId: fingerTableRoute.currentNodeInfo.Node_id, SenderConnNode: fingerTableRoute.currentNodeInfo  }
		entry := fingerTableRoute.findNode(joinReqPacket)

		if  !entry.EmptyEntry  {
			fingerTableRoute.routeConn[ findNodeId.String() ] = entry
		}
=======
		joinReqPacket :=  Util.FingerTablePacket{ Type : "JOIN_REQ", FingerTableID : findNodeId, SenderNodeId: fingerTableRoute.currentNodeInfo.Node_id }
		fingerTableRoute.routeConn[ findNodeId.String() ] = fingerTableRoute.findNode(joinReqPacket)
>>>>>>> f50043c628cbbc2aa8330621db2149f78ff4443e
	}
}


// [FIND NODE]
func (fingerTableRoute *FingerTableRoute) findNode(joinReqPacket Util.FingerTablePacket) TableEntry {
<<<<<<< HEAD
	fmt.Println( fingerTableRoute.availableNodeExist, fingerTableRoute.availableNodeInfo)

	if currBestNode := fingerTableRoute.closestPrecedingNode( joinReqPacket ); !currBestNode.EmptyEntry  && currBestNode.CurrNodeInfo.Node_id != fingerTableRoute.currentNodeInfo.Node_id  {
		// fmt.Printf("[Find-Node]: routed to BEST available node ID: %s\n",currBestNode.CurrNodeInfo.Node_id)
		return 	fingerTableRoute.findClosestPredecessorHelper( joinReqPacket, currBestNode.CurrNodeInfo )

	}else if fingerTableRoute.availableNodeExist {
		// fmt.Printf("[Find-Node]: routed to ANY available with IP_ADD: %s | PORT:  %s\n", fingerTableRoute.availableNodeInfo.IP_address, fingerTableRoute.availableNodeInfo.Ports["JOIN_REQ"])
		return fingerTableRoute.findClosestPredecessorHelper( joinReqPacket, fingerTableRoute.availableNodeInfo )
	}

	return TableEntry{ EmptyEntry:true }
=======

	if currBestNode := fingerTableRoute.currentBestNodeHelper( joinReqPacket ); !currBestNode.EmptyEntry {
		fmt.Printf("[Find-Node]: routed to BEST available node ID: %s\n",currBestNode.CurrNodeInfo.Node_id)
		return 	fingerTableRoute.findClosestPredecessorHelper( joinReqPacket, currBestNode.CurrNodeInfo )
	}
	fmt.Printf("[Find-Node]: routed to ANY available with IP_ADD: %s | PORT:  %s\n", fingerTableRoute.availableNodeInfo.IP_address, fingerTableRoute.availableNodeInfo.Ports["JOIN_REQ"])
	return fingerTableRoute.findClosestPredecessorHelper( joinReqPacket, fingerTableRoute.availableNodeInfo )
>>>>>>> f50043c628cbbc2aa8330621db2149f78ff4443e
}



// [FIND NODE HELPER]
func (fingerTableRoute *FingerTableRoute) findClosestPredecessorHelper(joinReqPacket Util.FingerTablePacket, contactNode Util.NodeInfo) TableEntry {
	networkMnger := NetworkModule.NewNetworkManager( contactNode.IP_address, contactNode.Ports["JOIN_REQ"] )

	if status := networkMnger.CreateTCPConnection(); !status {
		fmt.Println("[FindClosestPredecessorHelper][Error]: Unable to send join request packet.")
		return TableEntry{ EmptyEntry:true }
	}

	networkMnger.SendPacket(joinReqPacket)
	recvPkt := networkMnger.RecievePacket()
	
	if recvPkt.Type == "JOIN_FRW" {
<<<<<<< HEAD
		// fmt.Printf("[FindClosestPredecessorHelper]: Join request for [%s] forwarded to Node_ID: %s | IP_ADD: %s | PORT: %s\n", joinReqPacket.FingerTableID, recvPkt.ConnNode.Node_id, recvPkt.ConnNode.IP_address, recvPkt.ConnNode.Ports["JOIN_REQ"] )
=======
		fmt.Printf("[FindClosestPredecessorHelper]: Join request for [%s] forwarded to Node_ID: %s | IP_ADD: %s | PORT: %s\n", joinReqPacket.FingerTableID, recvPkt.ConnNode.Node_id, recvPkt.ConnNode.IP_address, recvPkt.ConnNode.Ports["JOIN_REQ"] )
>>>>>>> f50043c628cbbc2aa8330621db2149f78ff4443e
		return fingerTableRoute.findClosestPredecessorHelper( joinReqPacket, recvPkt.ConnNode)
	}

	networkMnger.SetIPAddress( recvPkt.ConnNode.IP_address, recvPkt.ConnNode.Ports["JOIN_RSP"])
	if status := networkMnger.CreateTCPConnection(); !status {
		fmt.Println("[FindClosestPredecessorHelper][Error]: Unable to receive join response packet.")
		return TableEntry{ EmptyEntry:true }
	}	

<<<<<<< HEAD
	// fmt.Printf("[FindClosestPredecessorHelper]: Best location for [%s] is Found at Node_ID: %s | IP_ADD: %s | PORT: %s\n", joinReqPacket.FingerTableID, recvPkt.ConnNode.Node_id, recvPkt.ConnNode.IP_address, recvPkt.ConnNode.Ports["JOIN_REQ"] )
=======
	fmt.Printf("[FindClosestPredecessorHelper]: Best location for [%s] is Found at Node_ID: %s | IP_ADD: %s | PORT: %s\n", joinReqPacket.FingerTableID, recvPkt.ConnNode.Node_id, recvPkt.ConnNode.IP_address, recvPkt.ConnNode.Ports["JOIN_REQ"] )
>>>>>>> f50043c628cbbc2aa8330621db2149f78ff4443e
	return TableEntry{ CurrNodeInfo:recvPkt.ConnNode,  ConnManager: networkMnger }
}

// # ******************************** end ****************************************** # //







// #------------------------------ services with their handlers ------------------------# //

// [SERVICE]
func (fingerTableRoute *FingerTableRoute) joinRespListnerService() {
	var joinReqServer NetworkModule.NetworkManager
	joinReqServer.SetIPAddress( fingerTableRoute.currentNodeInfo.IP_address, fingerTableRoute.currentNodeInfo.Ports["JOIN_RSP"] )

	joinReqServer.StartServer("TCP", fingerTableRoute.closeRequestServer, fingerTableRoute.fingerTableEntryServiceHandler)
}


// [FUNCTION-HANDLER][JOIN RESPONSE SERVICE]
func (fingerTableRoute *FingerTableRoute) fingerTableEntryServiceHandler(connection interface{}) {
	if connection, ok := connection.(net.Conn); ok { 
<<<<<<< HEAD
		
		for {

			dec := gob.NewDecoder(connection)
			packet := &Util.FingerTablePacket{}
			
			// fmt.Println("[----] Ping packet received ", packet)
			if err:= dec.Decode(packet); err != nil {
				fmt.Println("[fingerTableEntryServiceHandler][Error]: Unable to decode packet.")
				return
			}
	
			if packet.Type != "PING" {
				fingerTableRoute.resolvePacketChannel <- *packet
			}
=======
				 
		dec := gob.NewDecoder(connection)
		packet := &Util.FingerTablePacket{}

		if err:= dec.Decode(packet); err != nil {
			fmt.Println("[fingerTableEntryServiceHandler][Error]: Unable to decode packet.")
		}

		if packet.Type != "PING" {
			fingerTableRoute.resolvePacketChannel <- *packet
>>>>>>> f50043c628cbbc2aa8330621db2149f78ff4443e
		}

	}else{
		fmt.Println("[fingerTableEntryServiceHandler][Error]: Can't decode the connection socket...")
	}
}




// [SERVICE]
func (fingerTableRoute *FingerTableRoute) requestListnerService() {
	var joinReqServer NetworkModule.NetworkManager
	joinReqServer.SetIPAddress( fingerTableRoute.currentNodeInfo.IP_address, fingerTableRoute.currentNodeInfo.Ports["JOIN_REQ"] )

	joinReqServer.StartServer("TCP", fingerTableRoute.closeRequestServer, fingerTableRoute.requestListnerServiceHandler)
}


// [FUNCTION-HANDLER][REQ LISTNER SERVICE]
func (fingerTableRoute *FingerTableRoute) requestListnerServiceHandler(connection interface{}){
	if connection, ok := connection.(net.Conn); ok { 
			
		dec := gob.NewDecoder(connection)
		packet := &Util.FingerTablePacket{}
		if err:= dec.Decode(packet); err != nil {
			fmt.Println("[requestListnerServiceHandler][Error]: Unable to decode packet.")
			return
		}

		fingerTableRoute.availableNodeInfo = packet.SenderConnNode
		fingerTableRoute.availableNodeExist = true

		sendPacketToSocket(connection, fingerTableRoute.findSuccessor(packet) )

		connection.Close()
	}else{
		fmt.Println("[requestListnerServiceHandler][Error]: Can't decode the connection socket...")
	}
}





// [SERVICE]
func (fingerTableRoute *FingerTableRoute) resolvePacketService() {
	// for {
		packet := <- fingerTableRoute.resolvePacketChannel

		currBestNode := fingerTableRoute.closestPrecedingNode( packet )

		if !currBestNode.EmptyEntry && currBestNode.CurrNodeInfo.Node_id == fingerTableRoute.currentNodeInfo.Node_id || between( packet.SenderNodeId, packet.FingerTableID, fingerTableRoute.currentNodeInfo.Node_id){

			fmt.Println("Packet sent to me!", packet)
		}else{
			fingerTableRoute.routeToClosestNodeChannel <- packet
		}
	// }
}





// [SERVICE]
func (fingerTableRoute *FingerTableRoute) routeToClosestNodeService() {
	for {
		packet := <- fingerTableRoute.routeToClosestNodeChannel
		currBestNode := fingerTableRoute.closestPrecedingNode( packet )
		currBestNode.SendPacket(packet)
	}
}

// # ******************************** end ****************************************** # //








// #------------------------------ simple helper functions ------------------------# //


// [HELPER]

func (fingerTableRoute *FingerTableRoute) calculateNextFingerId(i int)  *big.Int {
	offset := new(big.Int).Exp(  big.NewInt( int64(fingerTableRoute.jumpSpacing) ), big.NewInt(int64(i)), nil)

	sum := new(big.Int).Add( fingerTableRoute.currentNodeInfo.Node_id, offset)
	ceil :=  new(big.Int).Exp( big.NewInt( int64(fingerTableRoute.jumpSpacing) ), fingerTableRoute.currentNodeInfo.M, nil)

	return new(big.Int).Mod(sum, ceil)
}



// [HELPER] : reverse sort the map keys

func (fingerTableRoute *FingerTableRoute) revSortFingerTableKeys() []*big.Int {
	keys := []*big.Int{}
	
	for key,_ := range fingerTableRoute.routeConn  {
		nodeId, _ :=  new(big.Int).SetString(key, 10)
		keys = append(keys, nodeId )
	}

	sort.SliceStable(keys,  func(index_1, index_2 int) bool  { return keys[index_1].Cmp( keys[index_2] ) > 0 } )
	return keys
}



// [HELPER]

func between(start, middle, end *big.Int) bool {
	if res := start.Cmp(end); res == -1 {
		return start.Cmp(middle) == -1 && middle.Cmp(end) <= 0
	}
	return start.Cmp(middle) == -1 || middle.Cmp(end) <= 0
}




// [HELPER]

func (fingerTableRoute *FingerTableRoute) closestPrecedingNode(reqPacket Util.FingerTablePacket) TableEntry {
	sortedKeys := fingerTableRoute.revSortFingerTableKeys()
	// fmt.Println(sortedKeys)
	for _, fingerId := range sortedKeys {
		// tableEntry := fingerTableRoute.routeConn[ fingerId.String() ]
		isActive:= true//tableEntry.Ping()

		if between( fingerTableRoute.currentNodeInfo.Node_id, fingerId, reqPacket.FingerTableID) && isActive{
			return fingerTableRoute.routeConn[ fingerId.String() ]

		}
		// else if !isActive {
		// 	delete( fingerTableRoute.routeConn, fingerId.String())
		// }
	}

	return TableEntry{  CurrNodeInfo : fingerTableRoute.currentNodeInfo  }
}


// [HELPER]

func (fingerTableRoute *FingerTableRoute) findSuccessor(packet *Util.FingerTablePacket) Util.FingerTablePacket {
	currBestNode := fingerTableRoute.closestPrecedingNode( *packet )
	fmt.Printf("[requestListnerServiceHandler]: Best location for [%s] is Found at Node_ID: %s \n", packet.FingerTableID, currBestNode.CurrNodeInfo.Node_id )

	successor, successor_exists := fingerTableRoute.routeConn[ fingerTableRoute.calculateNextFingerId( int(0) ).String() ];
	
	if successor_exists && ( between( fingerTableRoute.currentNodeInfo.Node_id, packet.FingerTableID, successor.CurrNodeInfo.Node_id) || packet.FingerTableID.Cmp(successor.CurrNodeInfo.Node_id) == 0 ){
		fmt.Printf("[requestListnerServiceHandler]: Best location is overrided by successor node [%s] \n",  successor.CurrNodeInfo.Node_id )
		return Util.FingerTablePacket{ Type : "JOIN_ACC", ConnNode:  successor.CurrNodeInfo, SenderConnNode: successor.CurrNodeInfo}

	}else if currBestNode.CurrNodeInfo.Node_id == fingerTableRoute.currentNodeInfo.Node_id {
		fmt.Printf("[requestListnerServiceHandler]: Best location is overrided by current node [%s] \n",  fingerTableRoute.currentNodeInfo.Node_id )

		return Util.FingerTablePacket{ Type : "JOIN_ACC", ConnNode:  currBestNode.CurrNodeInfo, SenderConnNode: currBestNode.CurrNodeInfo}

	}else{
		return Util.FingerTablePacket{ Type : "JOIN_FRW", ConnNode: currBestNode.CurrNodeInfo, SenderConnNode: currBestNode.CurrNodeInfo}
	}
}


// [HELPER]

func sendPacketToSocket(connection net.Conn, packet Util.FingerTablePacket) {
	enc := gob.NewEncoder(connection) 

	if err := enc.Encode(&packet); err != nil {
		log.Printf("[requestListnerServiceHandler][Error]: Failed to send packet back to requester: %v ... \n", err)
	}
}

// # ************************ end ************************ # //

