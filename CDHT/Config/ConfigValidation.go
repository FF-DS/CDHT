package Config

import (
	"strconv"
	"math/big"
	"fmt"
)


func (config *Configuration) ValidateConfig(){
	config.validateMode()
	config.validateRemoteAddr()
	config.validateNodePort()
	config.validateNodeM()
	config.validateNodeId()
	config.validateApplicationPort()
	config.validateCDHTPort()
}





func (config *Configuration) validateMode(){
	if config.Application_Mode == "FIRST_NODE" ||  config.Application_Mode == "SECOND_NODE" || config.Application_Mode == "REPLICA_NODE"{
		return 
	}
	inputVal :=  getInput("Enter  1 = FIRST_NODE / 2 = SECOND_NODE / 3 = REPLICA_NODE  Node Mode :")

	if inputVal == "1" {
		config.Application_Mode = "FIRST_NODE"
	}else if inputVal == "2" {
		config.Application_Mode = "SECOND_NODE"
	}else if inputVal == "3" {
		config.Application_Mode = "REPLICA_NODE"
	}

	config.validateMode()
}


func (config *Configuration) validateRemoteAddr(){
	if config.Application_Mode == "FIRST_NODE" || config.Remote_Node_Address != "" {
		return
	}
	config.Remote_Node_Address =  getInput("Enter Remote Node Address(ip:port) :")
	config.validateRemoteAddr()
}	


func (config *Configuration) validateNodePort(){
	port, err :=  strconv.Atoi( config.Node_Port )
	if err == nil && (1000 < port  &&  port < 65530) {
		return 
	}
	config.Node_Port =  getInput("Enter Node Port :")
	config.validateNodePort()
}	


func (config *Configuration) validateNodeM() {
	if config.Application_Mode == "REPLICA_NODE" {
		return
	}

	M, err :=  strconv.Atoi( config.Node_M )
	if err == nil && (3 < M  &&  M <= 160) {
		return
	}

	config.Node_M =  getInput("Enter M :")
	config.validateNodeM()
}


func (config *Configuration) validateNodeId() {
	if config.Application_Mode == "REPLICA_NODE" {
		return
	}

	_, ok := new(big.Int).SetString(config.Node_ID, 10)
	if config.Node_M == "160" || ok {
		return
	}

	_, config.Node_ID = config.getNodeId("Enter Node id :")
}


func (config *Configuration) validateApplicationPort(){
	port, err :=  strconv.Atoi( config.Application_Connecting_Port )
	if err == nil && (1000 < port  &&  port < 65530)  {
		return 
	}
	config.Application_Connecting_Port =  getInput("Enter Application Port :")
	config.validateApplicationPort()
}


func (config *Configuration) validateCDHTPort(){
	port, err :=  strconv.Atoi( config.CDHT_Ping_Tool_Listening_Port )
	if (err == nil && (1000 < port  &&  port < 65530)) || config.RUN_NETWORK_TEST_APPLICATION == false {
		return 
	}
	config.CDHT_Ping_Tool_Listening_Port =  getInput("Enter CDHT Ping Port :")
	config.validateCDHTPort()
}






// # -------------- helper -------------- #

func getInput(message string) string {
    var input string;
    fmt.Print(message)
    fmt.Scanln(&input)
	return input
}


func (config *Configuration) getNodeId(message string) (*big.Int, string) {
    var nodeID string;
	var NodeID *big.Int;
	ok := false

	base :=  big.NewInt( int64(config.Jump_Spacing) )
    m, _ := new(big.Int).SetString(config.Node_M, 10)
	minVal, maxVal := big.NewInt( -1 ),  base.Exp( base, m, nil)

	
    for !ok { 
        fmt.Print(message)
        fmt.Scanln(&nodeID)
        NodeID, ok = new(big.Int).SetString(nodeID, 10)

		if ok {
			ok = (minVal.Cmp(NodeID) < 0)  && (maxVal.Cmp(NodeID) > 0)
		}
	}
    return NodeID, nodeID
}