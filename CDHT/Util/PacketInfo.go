package Util

import (
	"math/big"
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
	SenderNodeId *big.Int
	ReceiverNodeId *big.Int
	RequestBody interface{}
}


func (reqObj *RequestObject) GetResponseObject() ResponseObject {
	return ResponseObject{
		Type: reqObj.Type,
		ResponseID: reqObj.RequestID,
		AppName: reqObj.AppName,
		AppID: reqObj.AppID,
		ResponseStatus: PACKET_STATUS_SUCCESS,
		SenderNodeId: reqObj.ReceiverNodeId,
		ReceiverNodeId: reqObj.SenderNodeId,
	}
}





// #--------------------------------- Response Object ------------------------# //

type ResponseObject struct {
	Type string
	ResponseID string
	AppName string
	AppID int
	ResponseStatus string
	SenderNodeId *big.Int
	ReceiverNodeId *big.Int
}


func (respObj *ResponseObject) GetRequestObject() RequestObject {
	return RequestObject{
		Type: respObj.Type,
		RequestID: respObj.ResponseID,
		AppName: respObj.AppName,
		AppID: respObj.AppID,
		SenderNodeId: respObj.ReceiverNodeId,
		ReceiverNodeId: respObj.SenderNodeId,
	}
}


