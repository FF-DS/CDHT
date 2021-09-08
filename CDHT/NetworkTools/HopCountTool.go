package NetworkTools

import (
	"cdht/RoutingModule"
	"cdht/ReportModule"
	"encoding/gob"
	"cdht/Util"
	"strconv"
	"math/big"
	"fmt"
	"time"
)


func init(){
	gob.Register( RoutingModule.NodeRPC{} )
}



type HopCountTool struct {
	routingTable *RoutingModule.RoutingTable
	networkTool *NetworkTool
}


func (hop *HopCountTool) getHopCount(req Util.RequestObject) (Util.RequestObject, bool) {
	if params, ok := req.RequestBody.(map[string]string); ok  {
		start := time.Now()

		endNodeId, _ :=  new(big.Int).SetString(params["END_NODE_ID"], 10)
		successor := hop.routingTable.LookUp( endNodeId )

		response := req.GetResponseObject()
		response.Type = Util.PACKET_TYPE_APPLICATION
		response.RequestBody =  map[string]interface{} {
			"successor": successor,
			"rtt": time.Since(start).String(),	
		}

		// ---- logging ----- //
		hop.networkTool.logNetworkReport( 
			ReportModule.LOG_TYPE_NETWORK_TOOL,
			ReportModule.LOG_LOCATION_TYPE_INCOMMING,
			ReportModule.LOG_OPERATION_STATUS_SUCCESS,
			map[string]string{
				"Tool_Name" : "HopCountTool",
				"Starting Node" : params["START_NODE_ID"],
				"End Node" : params["END_NODE_ID"],
				"Count" : strconv.Itoa( len(successor.NodeTraversalLogs) ) ,
			},
		)
		// ---- logging ----- //
		return response, true
	}
	return Util.RequestObject{}, false
}



func (hop *HopCountTool) RunTool(req Util.RequestObject){
	hopCountResult, success := hop.getHopCount(req)
	
	if !success {
		fmt.Println("[HopCountTool]:+ Unable to decode packet. ")
		return
	}

	resp := hop.routingTable.ForwardPacket(hopCountResult)
    
	if resp.ResponseStatus == Util.PACKET_STATUS_FAILED {
		fmt.Println("[HopCountTool]:+ Unable to create socket. ")
	}
}

