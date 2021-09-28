package Util

import (
	"math/big"
	"fmt"
)
// # --------------------------- eunm --------------------------- # //

const (
	PACKET_STATUS_SUCCESS 	string = "SUCCEEDED"
	PACKET_STATUS_NO_APP    string = "APP_NOT_FOUND"
	PACKET_STATUS_FAILED    string = "FAILED_TO_SEND"
	PACKET_STATUS_CORRUPTED string = "PACKET_CORRUPTED"
	PACKET_STATUS_NO_NODE   string = "NO_NODE_CONNECTED"

	PACKET_TYPE_APPLICATION string = "APPLICATION_PACKET"
	PACKET_TYPE_NETWORK     string = "NETWORK_TOOL"
	PACKET_TYPE_CLOSE       string = "TERMINATE_CONNECTION"
	PACKET_TYPE_INIT_APP    string = "APP_REGISTER"
)


// #--------------------------------- Packet ---------------------------------# //

type Packet struct {
	Type  string
	SenderIp string
	ReceiverIp string
	SenderNodeId *big.Int
	ReceiverNodeId *big.Int
}




// #--------------------------------- Request Object ------------------------# //

type RequestObject struct {
	Type string
	
	RequestID string
	AppName string
	AppID int
	UdpAddress string

	SenderNodeAddress string
	SenderNodeId *big.Int
	ReceiverNodeId *big.Int
	
	ResponseStatus string
	RequestBody interface{}
	
	SendToReplicas bool
	
	TranslateName bool
	TranslateNameID []byte

	ValidityCheck bool
	ValidityCheckResult bool
	ValidationCRCchecksum uint32
}


func (reqObj *RequestObject) GetResponseObject() RequestObject {
	return RequestObject{
		Type: reqObj.Type,
		RequestID: reqObj.RequestID,
		AppName: reqObj.AppName,
		AppID: reqObj.AppID,
		ResponseStatus: PACKET_STATUS_SUCCESS,
		SenderNodeId: reqObj.ReceiverNodeId,
		ReceiverNodeId: reqObj.SenderNodeId,
		ValidityCheck: reqObj.ValidityCheck,
	}
}


func  (reqObj *RequestObject) ToString() string {
	str := "---------------- Request Object Data ----------------\n"  
	str += fmt.Sprintf(" [+] Type : %s\n", reqObj.Type )
	str += fmt.Sprintf(" [+] RequestID : %s\n", reqObj.RequestID )
	str += fmt.Sprintf(" [+] AppName : %s\n", reqObj.AppName )
	str += fmt.Sprintf(" [+] AppID : %d\n", reqObj.AppID )
	str += fmt.Sprintf(" [+] UdpAddress : %s\n", reqObj.UdpAddress ) 
	str += fmt.Sprintf(" [+] SenderNodeId : %s\n", reqObj.SenderNodeId.String() )
	str += fmt.Sprintf(" [+] ReceiverNodeId : %s\n", reqObj.ReceiverNodeId.String() )
	str += fmt.Sprintf(" [+] ResponseStatus : %s\n", reqObj.ResponseStatus )
	str += fmt.Sprintf(" [+] RequestBody : %s\n", reqObj.RequestBody )
	str += fmt.Sprintf(" [+] SendToReplicas : %s\n", reqObj.SendToReplicas )
	str += fmt.Sprintf(" [+] TranslateName : %s\n", reqObj.TranslateName )
	str += fmt.Sprintf(" [+] TranslateNameID : %s\n", reqObj.TranslateNameID )
	str += fmt.Sprintf(" [+] ValidityCheck : %s\n", reqObj.ValidityCheck )
	str += fmt.Sprintf(" [+] ValidityCheckResult : %s\n", reqObj.ValidityCheckResult )
	str += fmt.Sprintf(" [+] ValidationCRCchecksum : %d\n", reqObj.ValidationCRCchecksum )
	str += "------------------------------------------\n"  
	return str;
}