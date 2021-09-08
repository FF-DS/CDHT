package CDHTNetworkTools

import (
	"math/big"
	"fmt"
)
// # --------------------------- eunm --------------------------- # //

const (
	COMMAND_OPERATION_STATUS_SUCCESS 	string = "SUCCESS"
	COMMAND_OPERATION_STATUS_FAILED     string = "FIALED"

	COMMAND_TYPE_HOP_COUNT		      	string = "COMMAND_TYPE_HOP_COUNT"
	COMMAND_TYPE_LOOK_UP            	string = "COMMAND_TYPE_LOOK_UP"
	COMMAND_TYPE_PING               	string = "COMMAND_TYPE_PING"
	COMMAND_TYPE_CLOSE_TOOL         	string = "COMMAND_TYPE_CLOSE_TOOL"

	URL_GET_COMMAND_FROM_SERVER         string = "https://cdht-monitoring-api.herokuapp.com/network-tools"
	URL_SEND_COMMAND_RESULT             string = "https://cdht-monitoring-api.herokuapp.com/network-tools"
)



// #--------------------------------- Log Object ---------------------------------# //

type ToolCommand struct {
	Type  string
	OperationID string
	OperationStatus  string
	NodeId *big.Int
	NodeAddress string
	Body interface{}
}



// #--------------------------------- string log ---------------------------------# //
func  (command *ToolCommand) ToString() string {
	str := "---------------- Tool Result ----------------\n"  
	str += fmt.Sprintf(" [+] Operation ID : %s\n", command.OperationID )
	str += fmt.Sprintf(" [+] Operation Type : %s\n", command.Type )
	str += fmt.Sprintf(" [+] Operation Status : %s\n", command.OperationStatus )
	str += fmt.Sprintf(" [+] Node Id : %s\n", command.NodeId.String() )
	str += fmt.Sprintf(" [+] Node Address : %s\n", command.NodeAddress ) 
	str += fmt.Sprintf(" [+] Result : %s\n", command.Body )
	str += "------------------------------------------\n"  
	return str;
}