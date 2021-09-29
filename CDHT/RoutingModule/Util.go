package RoutingModule

import  (
    "math/big"
    "os"
    "fmt"
	"cdht/Util"
	"cdht/ReportModule"
)


// # ------------------------ helpers ----------------------- #

func between(start, middle, end *big.Int) bool {
	if res := start.Cmp(end); res == -1 {
		return start.Cmp(middle) == -1 && middle.Cmp(end) <= 0
	}
	return start.Cmp(middle) == -1 || middle.Cmp(end) <= 0
}


func betweenClosest(start, middle, end *big.Int) bool {
	if res := start.Cmp(end); res == -1 {
		return start.Cmp(middle) == -1 && middle.Cmp(end) < 0
	}
	return start.Cmp(middle) == -1 || middle.Cmp(end) < 0
}


func checkError(err error) {
    if err != nil {
        fmt.Println("[Fatal Error]: + ", err.Error())
        os.Exit(1)
    }
}


func copyNodeData(old *NodeRPC, new *NodeRPC) {
    new.M = old.M
    new.Node_address = old.Node_address
    new.Node_id = old.Node_id
    new.DefaultArgs = nil
    new.NodeState = old.NodeState
    new.ReplicationCount = old.ReplicationCount

}


func copyResponseObject(curr *Util.RequestObject, resp *Util.RequestObject){
    resp.Type = curr.Type
	resp.RequestID = curr.RequestID
	resp.AppName = curr.AppName
	resp.AppID = curr.AppID
	resp.ResponseStatus = curr.ResponseStatus
	resp.SenderNodeId = curr.SenderNodeId
	resp.ReceiverNodeId = curr.ReceiverNodeId
}


func checkNode(node *NodeRPC) *NodeRPC {
    if node == nil || node.Node_id == nil {
        return nil
    }

    var nodeRPC *NodeRPC
    if node.DefaultArgs == nil || node.handle == nil{
        // fmt.Println("CHECK connection")
        _, nodeRPC = node.Connect()
    }else{
        _, nodeRPC = node.GetNodeInfo()
    }
    
    return nodeRPC
}

func CheckNode(node *NodeRPC) *NodeRPC {
    if node == nil || node.Node_id == nil {
        return nil
    }

    var nodeRPC *NodeRPC
    if node.DefaultArgs == nil {
        // fmt.Println("CHECK connection")
        _, nodeRPC = node.Connect()
    }else{
        _, nodeRPC = node.GetNodeInfo()
    }
    
    return nodeRPC
}



// # -------------------- report logging -------------------- # 
func (node *Node) logRoutingTableReport() {
    routeReport := ReportModule.Log {
        Type: ReportModule.LOG_TYPE_ROUTING_TABLE,
        OperationStatus: ReportModule.LOG_OPERATION_STATUS_SUCCESS,
        LogLocation: ReportModule.LOG_LOCATION_TYPE_SELF,
        NodeId: node.Node_id,
        NodeAddress: node.IP_address + ":" + node.Port,
        LogBody: map[string]([][]string){
            "FingerTable" : node.getFingerTableRouteEnty() ,
            "SuccessorsTable": node.getSuccessorsRouteEnty() ,
            "SuccPredTable": node.getSuccPredRouteEnty() ,
            "Replicas": node.getReplicas() ,
        },
    }

    node.Logger.RouteTableLog(routeReport)
}


func (node *Node) logNodeReport(logType string, location string, status string, logData map[string]string ) {
    fwdReport := ReportModule.Log {
        Type: logType,
        OperationStatus: status,
        LogLocation: location,
        NodeId: node.Node_id,
        NodeAddress: node.IP_address + ":" + node.Port,
        LogBody: logData,
    }

    node.Logger.NodeLog(fwdReport)
}


// # ------------------------------------ [Print node Info] ------------------------------------ #

func (node *Node) updatedNodeInfo() {
    fmt.Printf("------------------Current node Info[%s]------------------\n",node.Node_id.String())
    fmt.Printf("      [+]: Node ID : %s [%s] \n", node.Node_id.String(), node.NodeState)
    fmt.Printf("      [+]: M       : %s \n", node.M.String())
    fmt.Printf("      [+]: Address : %s:%s \n", node.IP_address, node.Port)
    fmt.Printf("      [+]: Jump Spacing : %d \n", node.JumpSpacing)
    fmt.Printf("      [+]: Replication Count : %d \n", node.ReplicationCount)
    fmt.Printf("      [+]: Successors Table Length : %d \n", node.SuccessorsTableLength)
    fmt.Printf("---------------------------------------------------------\n")
}