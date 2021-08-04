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

	FingerTableID  *big.Int
	SenderNodeId *big.Int
}



