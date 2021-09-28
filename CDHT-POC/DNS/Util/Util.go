package Util

import (
	"poc/app/DNS/Database"
	"fmt"
)

const (
	INSERT_RECORD_COMMAND 	 string = "Insert"
	FIND_RECORD_COMMAND      string =  "Find"
	UPDATE_RECORD_COMMAND 	 string =  "Update"
	DELETE_RECORD_COMMAND 	 string =  "Delete"
	RESPONSE_ADD_RECORD_COMMAND  string =  "Add Command Response"
	RESPONSE_FIND_RECORD_COMMAND  string =  "Find Command Response"
	RESPONSE_UPDATE_RECORD_COMMAND  string =  "Update Command Response"
	RESPONSE_DELETE_RECORD_COMMAND  string =  "Delete Command Response"
)


type AppCommand struct {
	RecordData Database.DnsRecord
	Command string
	SendToAll bool
	DoValidityCheck bool
	ValidityCheckResult bool
	NodeID string
	NodeAddress string
}

func (app *AppCommand) ToMap() map[string]interface{} {
	return map[string]interface{} {
		"RecordData" : app.RecordData.ToMap(),
		"Command" : app.Command,
	}
}

func ToAppCommand(app map[string]interface{}, req RequestObject) AppCommand {
	appCmd := AppCommand{}
	if record, ok := app["RecordData"].(map[string]string); ok { 
		appCmd.RecordData = Database.ToRecord(record)
	}

	if command, ok := app["Command"].(string); ok { 
		appCmd.Command = command
	}

	appCmd.DoValidityCheck = req.ValidityCheck
	appCmd.ValidityCheckResult = req.ValidityCheckResult

	return appCmd
}

func (app *AppCommand) ToString() string {
	str := "   ---------------- Command Response Data ----------------\n"  
	str += fmt.Sprintf("    [+] Command : %s\n", app.Command )
	str += fmt.Sprintf("    [+] Validity Check Required : %t\n", app.DoValidityCheck )
	str += fmt.Sprintf("    [+] Validity Check Result : %t\n", app.ValidityCheckResult )
	str += fmt.Sprintf("    [+] Sender Node ID : %s\n", app.NodeID )
	str += fmt.Sprintf("    [+] Sender Node Address : %s\n", app.NodeAddress )
	str += app.RecordData.ToString()
	str += "   ---------------------------------------------\n"  
	return str
}