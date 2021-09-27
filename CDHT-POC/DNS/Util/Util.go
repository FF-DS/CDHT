package Util

import (
	"poc/app/DNS/Database"
)

const (
	INSERT_RECORD_COMMAND 	 string = "Insert"
	FIND_RECORD_COMMAND      string =  "Find"
	UPDATE_RECORD_COMMAND 	 string =  "Update"
	DELETE_RECORD_COMMAND 	 string =  "Delete"
	RESPONSE_RECORD_COMMAND  string =  "Response"
)


type AppCommand struct {
	RecordData Database.DnsRecord
	Command string
	SendToAll bool
	DoValidityCheck bool
	ValidityCheckResult bool
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