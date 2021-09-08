package CDHTNetworkTools


import (
	"cdht/NetworkModule"
	"cdht/Util"
	"math/big"
	"fmt"
	"net"
)

type PingTool struct {
	AppServerIP string
	AppServerPort string
	ToolListeningPort string
	ChannelSize int
	NodeId *big.Int
	NodeAddress string
	CommandChannel chan ToolCommand
	ResultChannel chan ToolCommand
	close_server chan bool
}



// register network app & init env
func (netTool *PingTool) Init() {
	netTool.CommandChannel = make(chan ToolCommand, netTool.ChannelSize)
	netTool.close_server = make(chan bool)

	appConn := NetworkModule.NewNetworkManager( netTool.AppServerIP, netTool.AppServerPort)
    go appConn.Connect("UDP", netTool.sendCommandToNode)

	appServerConn := NetworkModule.NewNetworkManager( "", netTool.ToolListeningPort)
    go appServerConn.StartServer("UDP", netTool.close_server, netTool.receiveCommandResult)
}



func (netTool *PingTool)  sendCommandToNode(udpConnection interface{}){
	if soc, ok := udpConnection.(*net.UDPConn); ok { 

		appInitPacket := &Util.RequestObject{ AppName: "PingTool", Type: Util.PACKET_TYPE_INIT_APP, RequestBody : netTool.ToolListeningPort, }
		NetworkModule.SendUDPPacket(soc, appInitPacket)

		for {
			commandObject := <- netTool.CommandChannel
			if succ, requestObject := netTool.constructRequestObject(commandObject); !succ{
				fmt.Println("[Network-Tool][PingTool]:+ Unable to construct request object")
			}else{
				NetworkModule.SendUDPPacket(soc, &requestObject)
			}
		}
	}else {
		fmt.Println("[Network-Tool][PingTool]:+ Unable to Connect.")
	}
}



// conduct ping tool tasks, send to the overlay network & prepare response object
func (netTool *PingTool)  receiveCommandResult(requestObject interface{}){
	if requestObject, ok := requestObject.(Util.RequestObject); ok { 
		netTool.constructResponse(requestObject)
	}
}


func (netTool *PingTool) constructRequestObject(commandObject ToolCommand)  (bool, Util.RequestObject) {
	if params, ok := commandObject.Body.(map[string]string); ok  {
		return true, Util.RequestObject{
			Type: Util.PACKET_TYPE_NETWORK,
			RequestID: commandObject.OperationID,
			AppName: "PingTool",
			AppID: 0,
			SenderNodeId: netTool.NodeId,
			ReceiverNodeId: netTool.NodeId,
			RequestBody: map[string]string {
				"NODE_ID" : params["NODE_ID"],
			},
		}

	}else{
		return false, Util.RequestObject{}
	}
}



func (netTool *PingTool) constructResponse(responseObject Util.RequestObject){
	toolRespObj := ToolCommand{
		Type : COMMAND_TYPE_PING,
		OperationID : responseObject.RequestID,
	}
	result := map[string]interface{} { }

	if response, ok := responseObject.RequestBody.(map[string]string); ok  {
		result["rtt"] = response["rtt"]
		result["Node_Status"] = response["Node_Status"]
		result["Node_id"] = response["Node_id"]
		result["Node_Address"] = response["Node_Address"]
		toolRespObj.NodeId = netTool.NodeId
		toolRespObj.NodeAddress = netTool.NodeAddress
		toolRespObj.OperationStatus = COMMAND_OPERATION_STATUS_SUCCESS
	}else{
		toolRespObj.OperationStatus = COMMAND_OPERATION_STATUS_FAILED
	}

	toolRespObj.Body = result
	netTool.ResultChannel <- toolRespObj
}
