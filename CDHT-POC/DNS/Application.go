package DNS


import (
	"poc/app/DNS/Database"
	"poc/app/DNS/Util"
	"fmt"
)


type DnsApp struct {
	DnsApplicationResponseChannel chan Util.RequestObject
	DnsApplicationRequestChannel chan Util.AppCommand
	DnsApplicationRequestResponseChannel chan Util.RequestObject

	DnsApplicationResult chan Util.AppCommand

	DatabaseName string
	ApplicationSize int

	IPAddress string
	Port string

	conn  *Connection
	db  *DBHandler
}


func (dns *DnsApp) InitApp() *DnsApp{
	dns.DnsApplicationResponseChannel = make(chan Util.RequestObject, dns.ApplicationSize)
	dns.DnsApplicationRequestChannel = make(chan Util.AppCommand, dns.ApplicationSize)
	dns.DnsApplicationRequestResponseChannel = make(chan Util.RequestObject, dns.ApplicationSize)

	conn := Connection{
		DnsApplicationResponseChannel: dns.DnsApplicationResponseChannel,  DnsApplicationRequestChannel: dns.DnsApplicationRequestChannel,
		IPAddress: dns.IPAddress, Port: dns.Port, NetChannelSize: dns.ApplicationSize, DnsApplicationRequestResponseChannel: dns.DnsApplicationRequestResponseChannel,
	}
	dns.conn = conn.Init()

	db := DBHandler{ DatabaseName: dns.DatabaseName}
	dns.db = db.InitApp()

	go dns.executeRequests()
	return dns
}


func (dns *DnsApp) ListCurrentData() []Database.DnsRecord {
	return dns.db.ListCurrentData()
}


func (dns *DnsApp) AddRecords(record Database.DnsRecord) {
	dns.DnsApplicationRequestChannel <- Util.AppCommand {
		RecordData: record,
		Command: Util.INSERT_RECORD_COMMAND,
		SendToAll: true,
	}
}


func (dns *DnsApp) AddRecordsToCurrentNode(command Util.AppCommand, reqObj Util.RequestObject) Util.RequestObject {
	status := dns.db.AddRecords(command.RecordData)
	respObj := reqObj.GetResponseObject()

	if !status {
		respObj.ResponseStatus = Util.PACKET_STATUS_FAILED
	}

	command.Command = Util.RESPONSE_ADD_RECORD_COMMAND
	respObj.RequestBody = command.ToMap()
	return respObj
}


func (dns *DnsApp) FindRecord(record Database.DnsRecord) {
	dns.DnsApplicationRequestChannel <- Util.AppCommand {
		RecordData: record,
		Command: Util.FIND_RECORD_COMMAND,
		SendToAll: false,
	}
}

func (dns *DnsApp) FindRecordFromCurrentNode(command Util.AppCommand, reqObj Util.RequestObject) Util.RequestObject {
	record := dns.db.FindRecord(command.RecordData.RecordKey, command.RecordData.RecordType)
	respObj := reqObj.GetResponseObject()

	if record.RecordKey == "" {
		respObj.ResponseStatus = Util.PACKET_STATUS_FAILED
	}

	command.Command = Util.RESPONSE_FIND_RECORD_COMMAND
	command.RecordData = record
	respObj.RequestBody = command.ToMap()
	return respObj
}

func (dns *DnsApp) UpdateRecord(record Database.DnsRecord) {
	dns.DnsApplicationRequestChannel <- Util.AppCommand {
		RecordData: record,
		Command: Util.UPDATE_RECORD_COMMAND,
		SendToAll: true,
	}
}

func (dns *DnsApp) UpdateRecordOnCurrentNode(command Util.AppCommand, reqObj Util.RequestObject)  Util.RequestObject {
	status := dns.db.UpdateRecord(command.RecordData.RecordValue, command.RecordData.RecordKey, command.RecordData.RecordType)
	respObj := reqObj.GetResponseObject()

	if !status {
		respObj.ResponseStatus = Util.PACKET_STATUS_FAILED
	}

	command.Command = Util.RESPONSE_UPDATE_RECORD_COMMAND
	respObj.RequestBody = command.ToMap()
	return respObj
}

func (dns *DnsApp) RemoveRecord(record Database.DnsRecord) {
	dns.DnsApplicationRequestChannel <- Util.AppCommand {
		RecordData: record,
		Command: Util.DELETE_RECORD_COMMAND,
		SendToAll: true,
	}
}

func (dns *DnsApp) RemoveRecordFromCurrentNode(command Util.AppCommand, reqObj Util.RequestObject) Util.RequestObject {
	status := dns.db.RemoveRecord(command.RecordData.RecordKey, command.RecordData.RecordType)
	respObj := reqObj.GetResponseObject()

	if !status {
		respObj.ResponseStatus = Util.PACKET_STATUS_FAILED
	}

	command.Command = Util.RESPONSE_DELETE_RECORD_COMMAND
	respObj.RequestBody = command.ToMap()
	return respObj
}



func (dns *DnsApp) executeRequests() {
	for {
		reqObj := <- dns.DnsApplicationResponseChannel
		if appCmdMap, ok := reqObj.RequestBody.(map[string]interface{}); ok { 
			appCmd := Util.ToAppCommand(appCmdMap, reqObj)
			switch (appCmd.Command) {
				case Util.INSERT_RECORD_COMMAND:
					dns.DnsApplicationRequestResponseChannel <- dns.AddRecordsToCurrentNode(appCmd, reqObj)
				case Util.FIND_RECORD_COMMAND:
					dns.DnsApplicationRequestResponseChannel <- dns.FindRecordFromCurrentNode(appCmd, reqObj)
				case Util.UPDATE_RECORD_COMMAND:
					dns.DnsApplicationRequestResponseChannel <- dns.UpdateRecordOnCurrentNode(appCmd, reqObj)
				case Util.DELETE_RECORD_COMMAND:
					dns.DnsApplicationRequestResponseChannel <- dns.RemoveRecordFromCurrentNode(appCmd, reqObj)  // 
				default:
					appCmd.NodeID = reqObj.SenderNodeId.String()
					appCmd.NodeAddress = reqObj.SenderNodeAddress
					dns.DnsApplicationResult <- appCmd
			}
		}else{
			fmt.Printf("[DNS_APP]:+ Unable to Connect.\n")
		}
	}
}


func (dns *DnsApp) CloseApp() {
	dns.db.Close()
}