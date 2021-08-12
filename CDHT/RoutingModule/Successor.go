package RoutingModule

import (
	"math/big"
	"cdht/Util"
	"cdht/NetworkModule"
	"fmt"
	"net"
    "encoding/gob"
	"time"
)

type SuccessorTableRoute struct {
	predeccessor TableEntry
	fingerTable *FingerTableRoute
	currentNodeInfo Util.NodeInfo
	closeRequestServer (chan bool) 
}


// #---------------------------- INIT ----------------------------- #

func NewSuccessorTable( currNodeInfo Util.NodeInfo, fingerTable *FingerTableRoute) SuccessorTableRoute{
	succTable := SuccessorTableRoute{
		predeccessor : TableEntry{ EmptyEntry:true },
		fingerTable : fingerTable,
		currentNodeInfo : currNodeInfo,
		closeRequestServer : make(chan bool),
	}

	return succTable
}


func (succTable *SuccessorTableRoute) InitSuccessorService(){
	go succTable.predecessorNotificationListner()
	go succTable.successorReqListner()
}


func (succTable *SuccessorTableRoute) RunStablize(){
	for {
		time.Sleep(time.Second * 10)
		succTable.stablize()
	}
}

// #---------------------------- STABLIZE SERVICES ----------------------------- #

// [SERVICE]
func (succTable *SuccessorTableRoute) stablize(){
	succID := succTable.calculateSuccId().String()
	succ, ok := succTable.fingerTable.GetFingerTableState()[ succID ]

	if !ok { return }

	succReqPackt := Util.FingerTablePacket { Type : "SUCC_JOIN",  SenderNodeId :  succTable.currentNodeInfo.Node_id, ConnNode : succTable.currentNodeInfo  }
	succNodeInfo, ok := succTable.stablizeHelper( succ.CurrNodeInfo,  succReqPackt)

	if ok {
		entry, ok := succTable.createTableEntry( succNodeInfo )
		if ok {
			succTable.fingerTable.GetFingerTableState()[ succID ] = entry
		}
	}
	fmt.Printf("-----------------------------NODE [%s]---------------------------\n", succTable.currentNodeInfo.Node_id)
	fmt.Println("[stablizeHelper]: ROUTING table")
	fmt.Printf("		SUCC Entry : Node Id : %s  IP_ADD : %s  REQ_PORT : %s\n", succTable.fingerTable.GetFingerTableState()[ succID ].CurrNodeInfo.Node_id, succTable.fingerTable.GetFingerTableState()[ succID ].CurrNodeInfo.IP_address, succTable.fingerTable.GetFingerTableState()[ succID ].CurrNodeInfo.Ports["SUCC_REQ"])
	fmt.Printf("		PRED Entry : Node Id : %s  IP_ADD : %s  REQ_PORT : %s\n", succTable.predeccessor.CurrNodeInfo.Node_id, succTable.predeccessor.CurrNodeInfo.IP_address, succTable.predeccessor.CurrNodeInfo.Ports["SUCC_REQ"])

	fmt.Println("------------------------------------------------------------------")
}



// [HELPER]
func (succTable *SuccessorTableRoute) stablizeHelper(contactNode Util.NodeInfo, succPackt Util.FingerTablePacket)  (Util.NodeInfo, bool) {
	networkMnger := NetworkModule.NewNetworkManager( contactNode.IP_address, contactNode.Ports["SUCC_REQ"] )

	if status := networkMnger.CreateTCPConnection(); !status {
		fmt.Println("[stablizeHelper][Error]: Unable to send join request packet.")
		return Util.NodeInfo{}, false
	}

	networkMnger.SendPacket(succPackt)
	recvPkt := networkMnger.RecievePacket()
	
	if recvPkt.Type == "SUCC_FWD" {
		fmt.Printf("[stablizeHelper]: Successor Join request for [%s] forwarded to Node_ID: %s | IP_ADD: %s | PORT: %s\n", succPackt.SenderNodeId, recvPkt.ConnNode.Node_id, recvPkt.ConnNode.IP_address, recvPkt.ConnNode.Ports["SUCC_REQ"] )
		return succTable.stablizeHelper( recvPkt.ConnNode, succPackt)
	}
	
	return recvPkt.ConnNode, true
}




// #---------------------------- PREDECESSOR SERVICES ----------------------------- #

func (succTable *SuccessorTableRoute) predecessorNotificationListner(){
	var networkMnger NetworkModule.NetworkManager

	networkMnger.SetIPAddress( succTable.currentNodeInfo.IP_address, succTable.currentNodeInfo.Ports["PRED_RSP"])
	networkMnger.StartServer( "TCP",  succTable.closeRequestServer, predecessorNotificationHandler)
}


func predecessorNotificationHandler(connection interface{}){
	if connection, ok := connection.(net.Conn); ok { 
				 
		dec := gob.NewDecoder(connection)
		packet := &Util.FingerTablePacket{}

		if err:= dec.Decode(packet); err != nil {
			fmt.Println("[predecessorNotification][Error]: Unable to decode packet.")
		}else{
			fmt.Println("[predecessorNotification][PING]: ping received from")
		}
	}else{
		fmt.Println("[predecessorNotification][Error]: Can't decode the connection socket...")
	}
} 





// #---------------------------- SUCCESSOR SERVICES ----------------------------- #
// [SERVICE]
func (succTable *SuccessorTableRoute) successorReqListner(){
	var joinReqServer NetworkModule.NetworkManager
	joinReqServer.SetIPAddress( succTable.currentNodeInfo.IP_address, succTable.currentNodeInfo.Ports["SUCC_REQ"] )

	joinReqServer.StartServer("TCP", succTable.closeRequestServer, succTable.successorReqHandler)
}



func (succTable *SuccessorTableRoute) successorReqHandler(connection interface{}){
	if connection, ok := connection.(net.Conn); ok { 
		dec := gob.NewDecoder(connection)
		packet := &Util.FingerTablePacket{}

		if err:= dec.Decode(packet); err != nil {
			fmt.Println("[successorReqHandler][Error]: Unable to decode packet.")

		}else if succTable.predeccessor.EmptyEntry || !succTable.predeccessor.Ping() || packet.SenderNodeId.Cmp( succTable.predeccessor.CurrNodeInfo.Node_id ) >= 0 {
			// update predecessor table
			networkMnger := NetworkModule.NewNetworkManager( packet.ConnNode.IP_address, packet.ConnNode.Ports["PRED_RSP"] )
			if status := networkMnger.CreateTCPConnection(); !status {
				fmt.Println("[successorReqHandler][Error]: Unable to receive join response packet.")
				return
			}
			succTable.predeccessor = TableEntry{ CurrNodeInfo: packet.ConnNode,  ConnManager: networkMnger}

			sendPacketToSocket( connection, Util.FingerTablePacket{ Type : "SUCC_ACC", ConnNode: succTable.currentNodeInfo })
		}else{
			sendPacketToSocket( connection, Util.FingerTablePacket{ Type : "SUCC_FWD", ConnNode: succTable.predeccessor.CurrNodeInfo })
		}
	}else{
		fmt.Println("[successorReqHandler][Error]: Can't decode the connection socket...")
	}
}





// #---------------------------- HELPER FUNCTIONS ----------------------------- #

func (succTable *SuccessorTableRoute) calculateSuccId()  *big.Int {
	offset := new(big.Int).Exp(  big.NewInt( int64(succTable.fingerTable.GetJumpSpacing()) ), big.NewInt(int64(0)), nil)

	sum := new(big.Int).Add( succTable.currentNodeInfo.Node_id, offset)
	ceil :=  new(big.Int).Exp( big.NewInt( int64(succTable.fingerTable.GetJumpSpacing()) ), succTable.currentNodeInfo.M, nil)

	return new(big.Int).Mod(sum, ceil)
}

func (succTable *SuccessorTableRoute) createTableEntry(nodeInfo Util.NodeInfo) (TableEntry, bool) {
	fmt.Println("[Successor]: createTableEntry.")

	networkMnger := NetworkModule.NewNetworkManager(nodeInfo.IP_address, nodeInfo.Ports["JOIN_RSP"])

	if status := networkMnger.CreateTCPConnection(); !status {
		fmt.Println("[Successor][Error]: Unable to receive join response packet.")
		return TableEntry{ EmptyEntry:true }, false
	}	
		
	fmt.Printf("[Successor]: successor is updated with node id %s \n", nodeInfo.Node_id.String() )
	return TableEntry{ CurrNodeInfo:nodeInfo,  ConnManager: networkMnger }, true
}