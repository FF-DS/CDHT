package Util

import (
	"math/big"
)
// # --------------------------- PACKET INFO --------------------------- # //

type Packet struct {
	Type  string
	SenderIp string
	ReceiverIp string
	SenderNodeId *big.Int
	ReceiverNodeId *big.Int
}

type FingerTablePacket struct {
	Type  string
	SenderIp string
	ReceiverIp string
	SenderNodeId *big.Int
	ReceiverNodeId *big.Int
	FingerTableID  *big.Int
	Ports map[string]string
}



