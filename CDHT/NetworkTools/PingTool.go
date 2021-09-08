package NetworkTools

import (
	"cdht/RoutingModule"
	"math/big"
	"cdht/Util"
	"time"
	"fmt"
)


type PingTool struct {
	routingTable *RoutingModule.RoutingTable
	networkTool *NetworkTool
}


func (ping *PingTool) getPingResult(req Util.RequestObject) (Util.RequestObject, bool) {
	if params, ok := req.RequestBody.(map[string]string); ok  {
		start := time.Now()

		NodeId, _ :=  new(big.Int).SetString(params["NODE_ID"], 10)
		successor := ping.routingTable.LookUp( NodeId )
		
		response := req.GetResponseObject()
		response.Type = Util.PACKET_TYPE_APPLICATION

		status := "Alive"
		if successor.Node_id == nil {
			status = "Down"
		}
		
		response.RequestBody = map[string]string{
			"Node_Status" : status,
			"rtt": time.Since(start).String(),
			"Node_id": successor.Node_id.String(),
			"Node_Address": successor.Node_address,
		}

		return response, true
	}

	return Util.RequestObject{}, false
}



func (ping *PingTool) RunTool(req Util.RequestObject){
	pingResult, success := ping.getPingResult(req)
	
	if !success {
		fmt.Println("[PingTool]:+ Unable to decode packet. ")
		return
	}

	resp := ping.routingTable.ForwardPacket(pingResult)
    
	if resp.ResponseStatus == Util.PACKET_STATUS_FAILED {
		fmt.Println("[PingTool]:+ Unable to create socket. ")
	}
}