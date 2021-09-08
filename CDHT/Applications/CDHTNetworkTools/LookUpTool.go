package CDHTNetworkTools


import (
	"cdht/NetworkModule"
	"cdht/RoutingModule"
	"cdht/Util"
	"math/big"
	"fmt"
	"net"
)

type LookUpTool struct {
	AppServerIP string
	AppServerPort string
	ChannelSize int
	NodeId *big.Int
	NodeAddress string
	CommandChannel chan ToolCommand
	ResultChannel chan ToolCommand
}



// register network app & init env
func (netTool *LookUpTool) Init() {
	netTool.CommandChannel = make(chan ToolCommand, netTool.ChannelSize)
    appConn := NetworkModule.NewNetworkManager( netTool.AppServerIP, netTool.AppServerPort)
	go appConn.Connect("TCP", netTool.handleApplicationConnection)
}



func (netTool *LookUpTool)  handleApplicationConnection(connection interface{}){
	if connection, ok := connection.(net.Conn); ok { 
		netChannel := NetworkModule.NetworkChannel{ Connection: connection, ChannelSize: netTool.ChannelSize }
		netChannel.Init()

		// register app name
		netChannel.SendToSocket(Util.RequestObject{  AppName: "LookUpTool", Type: Util.PACKET_TYPE_INIT_APP,})        
		netTool.excuteHopCountToolRequest(netChannel)
	}else {
		fmt.Println("[Network-Tool][LookUpTool]:+ Unable to Connect.")
	}
}



// conduct look up tool tasks, send to the overlay network & prepare response object
func (netTool *LookUpTool)  excuteHopCountToolRequest(netChannel NetworkModule.NetworkChannel){
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
					fmt.Println("[Network-Tool][LookUpTool]:+ Unable to construct request object")
				}

				if status := netChannel.SendToSocket( requestObject ); !status{
					fmt.Println("[Network-Tool][LookUpTool]:+ Unable to send request object")
				}
		}
	}
}



func (netTool *LookUpTool) constructResponse(responseObject Util.RequestObject){
	toolRespObj := ToolCommand{
		Type : COMMAND_TYPE_LOOK_UP,
		OperationID : responseObject.RequestID,
	}
	result := map[string]interface{} { }

	if response, ok := responseObject.RequestBody.(map[string]interface{}); ok  {
		if successor, ok := response["successor"].(RoutingModule.NodeRPC); ok  {
			result["NodeId"] =  successor.Node_id
			result["NodeAddress"] =  successor.Node_address
			result["NodeM"] =  successor.M

		}
		result["rtt"] = response["rtt"]
		toolRespObj.NodeId = netTool.NodeId
		toolRespObj.NodeAddress = netTool.NodeAddress
		toolRespObj.OperationStatus = COMMAND_OPERATION_STATUS_SUCCESS
	}else{
		toolRespObj.OperationStatus = COMMAND_OPERATION_STATUS_FAILED
	}

	toolRespObj.Body = result
	netTool.ResultChannel <- toolRespObj
}


func (netTool *LookUpTool) constructRequestObject(commandObject ToolCommand)  (bool, Util.RequestObject) {
	if params, ok := commandObject.Body.(map[string]string); ok  {
		return true, Util.RequestObject{
			Type: Util.PACKET_TYPE_NETWORK,
			RequestID: commandObject.OperationID,
			AppName: "LookUpTool",
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