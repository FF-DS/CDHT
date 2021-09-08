package CDHTNetworkTools

import (
	"cdht/NetworkModule"
	"cdht/RoutingModule"
	"cdht/Util"
	"math/big"
	"fmt"
	"net"
	"strconv"
)

type HopCountTool struct {
	AppServerIP string
	AppServerPort string
	ChannelSize int
	NodeId *big.Int
	NodeAddress string
	CommandChannel chan ToolCommand
	ResultChannel chan ToolCommand
}



// register network app & init env
func (netTool *HopCountTool) Init() {
	netTool.CommandChannel = make(chan ToolCommand, netTool.ChannelSize)
    appConn := NetworkModule.NewNetworkManager( netTool.AppServerIP, netTool.AppServerPort)
	go appConn.Connect("TCP", netTool.handleApplicationConnection)
}



func (netTool *HopCountTool)  handleApplicationConnection(connection interface{}){
	if connection, ok := connection.(net.Conn); ok { 
		netChannel := NetworkModule.NetworkChannel{ Connection: connection, ChannelSize: netTool.ChannelSize }
		netChannel.Init()

		// register app name
		netChannel.SendToSocket(Util.RequestObject{  AppName: "HopCountTool", Type: Util.PACKET_TYPE_INIT_APP,})        
		netTool.excuteHopCountToolRequest(netChannel)
	}else {
		fmt.Println("[Network-Tool][HopCountTool]:+ Unable to Connect.")
	}
}



// conduct hop count tool tasks, send to the overlay network & prepare response object
func (netTool *HopCountTool)  excuteHopCountToolRequest(netChannel NetworkModule.NetworkChannel){
	for {
		select {
			case netReqObj := <- netChannel.ReqChannel:
				if netReqObj.Type == Util.PACKET_TYPE_CLOSE {
					return
				}
				
				netTool.constructResponse(netReqObj)
			case commandObject := <- netTool.CommandChannel:
				succ, requestObject := netTool.constructRequestObject(commandObject); 
				if !succ{
					fmt.Println("[Network-Tool][HopCountTool]:+ Unable to construct request object")
				}

				if status := netChannel.SendToSocket( requestObject ); !status{
					fmt.Println("[Network-Tool][HopCountTool]:+ Unable to send request object")
				}
		}
	}
}



func (netTool *HopCountTool) constructResponse(responseObject Util.RequestObject){
	toolRespObj := ToolCommand{
		Type : COMMAND_TYPE_HOP_COUNT,
		OperationID : responseObject.RequestID,
	}
	result := map[string]interface{} { }

	if response, ok := responseObject.RequestBody.(map[string]interface{}); ok  {
		hops := []map[string]string{}
		logged := make(map[string]string)

		if successor, ok := response["successor"].(RoutingModule.NodeRPC); ok  {
			for _, node := range successor.NodeTraversalLogs {
				if _, logged := logged[node.Node_id.String()]; logged { continue }

				hops = append( hops, map[string]string{ "Node_id": node.Node_id.String(), "Node_address": node.Node_address, } )
				logged[node.Node_id.String()] = node.Node_id.String()
			}
		}

		result["hops"] = hops
		result["rtt"] = response["rtt"]
		result["length"] = strconv.Itoa( len(logged) )
		toolRespObj.NodeId = netTool.NodeId
		toolRespObj.NodeAddress = netTool.NodeAddress

		toolRespObj.OperationStatus = COMMAND_OPERATION_STATUS_SUCCESS
	}else{
		toolRespObj.OperationStatus = COMMAND_OPERATION_STATUS_FAILED
	}

	toolRespObj.Body = result
	netTool.ResultChannel <- toolRespObj
}


func (netTool *HopCountTool) constructRequestObject(commandObject ToolCommand)  (bool, Util.RequestObject) {
	if params, ok := commandObject.Body.(map[string]string); ok  {
		startNodeId, _ :=  new(big.Int).SetString(params["START_NODE_ID"], 10)
		return true, Util.RequestObject{
			Type: Util.PACKET_TYPE_NETWORK,
			RequestID: commandObject.OperationID,
			AppName: "HopCountTool",
			AppID: 0,
			SenderNodeId: netTool.NodeId,
			ReceiverNodeId: startNodeId,
			RequestBody: map[string]string {
				"START_NODE_ID" : params["START_NODE_ID"],
				"END_NODE_ID" : params["END_NODE_ID"],
			},
		}

	}else{
		return false, Util.RequestObject{}
	}
}