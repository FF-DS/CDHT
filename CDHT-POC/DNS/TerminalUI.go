package DNS

import (
	"poc/app/DNS/Database"
	"poc/app/DNS/Util"
    "strings"
    "bufio"
    "fmt"
    "os"
)

type TerminalUI struct {
	DnsApplicationResult chan Util.AppCommand
	app  *DnsApp
}

func (ui *TerminalUI) Init(){
	ui.DnsApplicationResult = make(chan Util.AppCommand, 100000)
	address := getInput("Enter IP address and Port number separated by space :")
	if len(address) != 2 {
		ui.Init()
	}
	app := DnsApp{ IPAddress: address[0], Port: address[1], ApplicationSize: 100000, DatabaseName: "dns_app2.sql", DnsApplicationResult: ui.DnsApplicationResult}
	ui.app = app.InitApp()
}


func (ui *TerminalUI) UserUI(){
    for {
        params := getInput("> ")

        switch params[0] {
            case "close" : 
				ui.app.CloseApp()
                return
            case "find" :
                ui.findRecord(params)
            case "list" :
                ui.listLocalRecords(params)
            case "update" :
                ui.updateRecord(params)
            case "add":
                ui.addRecords(params)
            case "delete":
                ui.deleteRecords(params)
			case "results":
                ui.resultRecords()
			case "help":
				help()
        }

    }
}

func (ui *TerminalUI) findRecord(params []string) {
	recordType := getInput("   Enter Record Type :")
	recordKey := getInput("   Enter Record Key :")
	ui.app.FindRecord(Database.DnsRecord{
		RecordType: recordType[0],
		RecordKey: recordKey[0],
	})
}

func (ui *TerminalUI) listLocalRecords( params []string) {
	if len(params) != 2 || params[1] != "local" {
		return
	}

    fmt.Println("===================== [Local Records] ======================")
    records := ui.app.ListCurrentData()
    for _, record := range records {
        fmt.Println(record.ToString())
    }
}


func (ui *TerminalUI) addRecords(params []string) {
    recordType := getInput("   Enter Record Type :")
	recordKey := getInput("   Enter Record Key :")
	recordValue := getInput("   Enter Record Value :")

	ui.app.AddRecords(Database.DnsRecord{
		RecordType: recordType[0],
		RecordKey: recordKey[0],
		RecordValue: recordValue[0],
	})
}


func (ui *TerminalUI) updateRecord(params []string) {
    recordType := getInput("   Enter Record Type :")
	recordKey := getInput("   Enter Record Key :")
	recordValue := getInput("   Enter Record Value :")

	ui.app.UpdateRecord(Database.DnsRecord{
		RecordType: recordType[0],
		RecordKey: recordKey[0],
		RecordValue: recordValue[0],
	})
}


func (ui *TerminalUI) deleteRecords(commands []string) {
	recordType := getInput("   Enter Record Type :")
	recordKey := getInput("   Enter Record Key :")

	ui.app.RemoveRecord(Database.DnsRecord{
		RecordType: recordType[0],
		RecordKey: recordKey[0],
	})
} 

func (ui *TerminalUI) resultRecords(){
    for len(ui.DnsApplicationResult) > 0 {
		appCommand := <- ui.DnsApplicationResult
		fmt.Println(appCommand.ToString())
	}
}



// # ------------------  [Helper]  ------------------ #

func getInput(inputStr string) []string {
    fmt.Print(inputStr)
    scanner := bufio.NewScanner(os.Stdin)
    var strInput string
    if scanner.Scan() {
        strInput = scanner.Text()
    }

    return strings.Split(strInput, " ")
}


func help(){
    fmt.Println("=========================================== [Help] ===========================================")
    fmt.Println("    [+]close :  close the application")
    fmt.Println("    [+]find :  find DNS records by specifying d/t parameters")
    fmt.Println("    [+]list local :  list the current node DNS records")
    fmt.Println("    [+]update :  do an update operation on DNS records")
    fmt.Println("    [+]add :  add a DNS records")
    fmt.Println("    [+]delete :  delete a DNS records")
    fmt.Println("    [+]results : this will print all the dns queries of the applications")
    fmt.Println("==============================================================================================")
}