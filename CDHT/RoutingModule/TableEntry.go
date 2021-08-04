package RoutingModule

import (
	"cdht/Util"
	"cdht/NetworkModule"
)

type TableEntry struct {
	CurrNodeInfo  Util.NodeInfo
	ConnManager NetworkModule.NetworkManager
	EmptyEntry bool
}



func (tableEntry *TableEntry) SendPacket(packet Util.FingerTablePacket) bool {
	return tableEntry.ConnManager.SendPacket(packet)
}


func (tableEntry *TableEntry) Ping() bool {
	pingPacket := Util.FingerTablePacket{
		Type: "PING",
	}

	return tableEntry.ConnManager.SendPacket(pingPacket)
}