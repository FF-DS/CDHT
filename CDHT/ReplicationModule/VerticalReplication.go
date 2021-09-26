// package ReplicationModule

// import (
// 	"cdht/RoutingModule"
// 	"fmt"
// )


// type VerticalReplication struct {
// 	CurrRoutingTable  *RoutingModule.RoutingModule
// 	MainNode *RoutingModule.NodeRPC
// }


// func (ver *VerticalReplication) InitalizeAsReplica() {
// 	if RoutingModule.CheckNode( ver.CurrRoutingTable ) == nil {
// 		fmt.Println("[REPLICA][ERROR]: Unable to connect to the main node")
// 		return
// 	}

// 	CurrRoutingTable.InitRoutingTable()

// 	CurrRoutingTable.NodeInfo().ReplicaInfos = ver.CurrRoutingTable.
// 	// call RPC method
// 	// extract remote node info
// 	// update info
// }


// func (ver *VerticalReplication) UpdateInfo() {
// 	// with time interval
// 	// *** extract remote info 
// 	// *** update and sleep
// 	// * or 
// 	// *** find the next active node on the replica group
// 	// *** go to step one
// }