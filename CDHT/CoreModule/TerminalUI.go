package CoreModule

import (
    "cdht/Applications/CDHTNetworkTools"
    "cdht/ReportModule"

    "strconv"
    "strings"
    "math/big"
    "bufio"
    "fmt"
    
    "time"
    "os"
)

type TerminalUI struct {
	CoreLink *Core
}

func (ui *TerminalUI) UserUI(){
    for {
        params := getInput("> ")

        switch params[0] {
            case "close" : 
                return
            case "route" :
                ui.printRoutes(params)
            case "lookup" :
                ui.lookUpUI(params)
            case "log" :
                ui.logDump(params)
            case "tool":
                ui.testCDHTtool(params[1])
        }

    }
}

func (ui *TerminalUI) printRoutes(params []string) {
    if len(params) == 1 {
        ui.CoreLink.RoutingTableInfo.PrintRoutingInfo()
    }

    if len(params) == 2 {
        count, err := strconv.Atoi(params[1]) 
        if err != nil {
            count = 3
        }
        for i := 0; i < count; i++ {
            time.Sleep(time.Second * 2)
            ui.CoreLink.RoutingTableInfo.PrintRoutingInfo()
        }
    }
}

func (ui *TerminalUI) lookUpUI( params []string) {
    var nodeId string
    fmt.Print("Enter Node Id: ")
    fmt.Scanln(&nodeId)

    NodeId, ok := new(big.Int).SetString(nodeId, 10)
    if !ok { fmt.Println("SetString: error node id") }

    succ := ui.CoreLink.RoutingTableInfo.LookUp(NodeId)
    succ.PrintNodeInfo()
    fmt.Println("===================== [Trace] ======================")

    logged := make(map[string]string)
    for _, node := range succ.NodeTraversalLogs {
        if _, logged := logged[node.Node_id.String()]; logged {
            continue
        }

        node.PrintNodeInfo()
        logged[node.Node_id.String()] = node.Node_id.String()
    }
}

func (ui *TerminalUI) logDump(params []string) {
    if len(params) < 2 {
        return 
    }

    switch params[1] {
        case "route":
            printLog( ui.CoreLink.LogManager.RouteLogs() )
        case "node":
            printLog( ui.CoreLink.LogManager.NodeLogs() )
        case "network":
            printLog( ui.CoreLink.LogManager.NetworkToolLogs() )
        case "config":
            printLog( ui.CoreLink.LogManager.ConfigToolLogs() )
    }

}


func (ui *TerminalUI) testCDHTtool(command string) {
    if  ui.CoreLink.CdhtNetworkTools ==  nil {
        return
    }
    switch command {
        case "hop":
            ui.hopCountTool()
        case "lookup":
            ui.lookUpTool()
        case "ping":
            ui.pingTool()
        case "log":
            ui.printCDHTlog()
    }
} 

func (ui *TerminalUI) hopCountTool(){
    var startNodeId, endNodeId string
    fmt.Print("Start Node ID: ")
    fmt.Scanln(&startNodeId)
    fmt.Print("End Node ID: ")
    fmt.Scanln(&endNodeId)

    command := CDHTNetworkTools.ToolCommand{
        Type: CDHTNetworkTools.COMMAND_TYPE_HOP_COUNT,
	    OperationID: "1232",
	    Body: map[string]string{
            "START_NODE_ID" : startNodeId,
            "END_NODE_ID" : endNodeId,
        },
    }
    ui.CoreLink.CdhtNetworkTools.DispatchCommands(command)
}


func (ui *TerminalUI) lookUpTool(){
    var NodeId string
    fmt.Print("Node ID: ")
    fmt.Scanln(&NodeId)

    command := CDHTNetworkTools.ToolCommand{
        Type: CDHTNetworkTools.COMMAND_TYPE_LOOK_UP,
	    OperationID: "1232",
	    Body: map[string]string{
            "NODE_ID" : NodeId,
        },
    }
    ui.CoreLink.CdhtNetworkTools.DispatchCommands(command)
}


func (ui *TerminalUI) pingTool(){
    var NodeId string
    fmt.Print("Node ID: ")
    fmt.Scanln(&NodeId)

    command := CDHTNetworkTools.ToolCommand{
        Type: CDHTNetworkTools.COMMAND_TYPE_PING,
	    OperationID: "1232",
	    Body: map[string]string{
            "NODE_ID" : NodeId,
        },
    }
    ui.CoreLink.CdhtNetworkTools.DispatchCommands(command)
}

func (ui *TerminalUI) printCDHTlog(){
    commands := []CDHTNetworkTools.ToolCommand {}
    for len(ui.CoreLink.CdhtNetworkTools.ResultChannel) > 0 {
        command := <- ui.CoreLink.CdhtNetworkTools.ResultChannel
        fmt.Println( command.ToString() )
        commands = append(commands, command)
    }

    for _, command := range commands {
        ui.CoreLink.CdhtNetworkTools.ResultChannel <- command
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


func printLog(logs []ReportModule.Log){
    for _, log := range logs {
        fmt.Println(log.ToString())
    }
}
