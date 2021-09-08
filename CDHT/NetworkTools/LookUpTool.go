package NetworkTools

import (
	"cdht/RoutingModule"
	"cdht/ReportModule"
	"encoding/gob"
	"cdht/Util"
	"math/big"
	"time"
	"fmt"
)


type LookUpTool struct {
	routingTable *RoutingModule.RoutingTable
	networkTool *NetworkTool
}


func init(){
	gob.Register( RoutingModule.NodeRPC{} )
}


func (lookUp *LookUpTool) getLookUpResult(req Util.RequestObject) (Util.RequestObject, bool) {
	if params, ok := req.RequestBody.(map[string]string); ok  {
		start := time.Now()
		
		NodeId, _ :=  new(big.Int).SetString(params["NODE_ID"], 10)
		successor := lookUp.routingTable.LookUp( NodeId )

		response := req.GetResponseObject()
		response.Type = Util.PACKET_TYPE_APPLICATION
		response.RequestBody =  map[string]interface{} {
			"successor": successor,
			"rtt": time.Since(start).String(),	
		}

		// ---- logging ----- //
		lookUp.networkTool.logNetworkReport( 
			ReportModule.LOG_TYPE_NETWORK_TOOL,
			ReportModule.LOG_LOCATION_TYPE_LEAVING,
			ReportModule.LOG_OPERATION_STATUS_SUCCESS,
			map[string]string{
				"Tool_Name" : "LookUpTool",
				"Starting Node" : params["NODE_ID"],
				"Node Address" : successor.Node_address,
			},
		)
		// ---- logging ----- //
		return response, true
	}
	return Util.RequestObject{}, false
}



func (lookUp *LookUpTool) RunTool(req Util.RequestObject){
	lookUpResult, success := lookUp.getLookUpResult(req)
	
	if !success {
		fmt.Println("[LookUpTool]:+ Unable to decode packet. ")
		return
	}

	resp := lookUp.routingTable.ForwardPacket(lookUpResult)
    
	if resp.ResponseStatus == Util.PACKET_STATUS_FAILED {
		fmt.Println("[LookUpTool]:+ Unable to create socket. ")
	}
}