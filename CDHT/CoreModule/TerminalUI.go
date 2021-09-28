package CoreModule

import (
    "cdht/Applications/CDHTNetworkTools"
    "cdht/ReportModule"
    // "cdht/Config"

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
                ui.testCDHTtool(params)
            case "config":
                ui.uiConfigManager(params)
            case "replica":
                ui.printReplica(params)
            case "node":
                ui.CoreLink.RoutingTableInfo.PrintUpdatedNodeInfo()
            case "help":
                help()
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
    _, nodeId := ui.getNodeId( "Enter Node ID: " )

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


func (ui *TerminalUI) testCDHTtool(commands []string) {
    if  ui.CoreLink.CdhtNetworkTools ==  nil {
        return
    }
    if len(commands) < 2 {
        return
    }
    switch commands[1] {
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
    _, startNodeId := ui.getNodeId( "Start Node ID: " )
    _, endNodeId := ui.getNodeId( "End Node ID: " )

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
    _, NodeId := ui.getNodeId( "Node ID: " )

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
    _, NodeId := ui.getNodeId( "Node ID: " )

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


func (ui *TerminalUI) uiConfigManager(commands []string) {
    if  ui.CoreLink.Config ==  nil {
        return
    }
    
    if len(commands) < 2 {
        ui.printConfig()
        return
    }

    switch commands[1] {
        case "load":
            ui.loadFromFile()
        default:
            ui.printConfig()
    }
} 

func (ui *TerminalUI) printConfig(){
    ui.CoreLink.Config.PrintConfiguration()
}

func (ui *TerminalUI) loadFromFile(){
    ui.CoreLink.ConfigMngr.UpdateFromFile()
}


func (ui *TerminalUI) printReplica(params []string){
    if len(params) == 2 && params[1] == "remote" { // && ui.CoreLink.Config.Application_Mode == Config.MODE_REPLICA_NODE
        ui.CoreLink.RoutingTableInfo.PrintRemoteReplicaInfo()
    }else {
        ui.CoreLink.RoutingTableInfo.PrintCurrentReplicaInfo()
    }

}


// # ------------------  [Helper]  ------------------ #

func help(){
    fmt.Println("=========================================== [Help] ===========================================")
    fmt.Println("    [+]close :  close the application")
    fmt.Println("    [+]route :  show the current successor, predecessor & finger tables")
    fmt.Println("           ==>  route number:  show the tables the amount of provided (number) times with 2 second delay")
    fmt.Println("    [+]lookup:  conducts a lookup by providing node id")
    fmt.Println("           ==>  node_id: it should be a valid node id")
    fmt.Println("    [+]log:  print out logs that haven't yet sent to server")
    fmt.Println("           ==>  log route: routing table logs")
    fmt.Println("           ==>  log node: node forwarding activity logs")
    fmt.Println("           ==>  log network: network tools log data")
    fmt.Println("           ==>  log config: configuration changes log")
    fmt.Println("    [+]tool:  if the CDHT tool is running on the system this command used to do the following things")
    fmt.Println("           ==>  tool hop: do hop count")
    fmt.Println("           ==>  tool lookup: do lookup count")
    fmt.Println("           ==>  tool ping: do node ping ")
    fmt.Println("           ==>  tool log: check CDHT log info")
    fmt.Println("    [+]config:  loads the current configuration profile")
    fmt.Println("           ==>  config load: loads the current configuration profile from a file")
    fmt.Println("    [+]node: current node profile information")
    fmt.Println("    [+]replica: this command will list the replica's connected to the current node")
    fmt.Println("           ==>  replica remote: this will list all the routing information of the remote Main node.")
    fmt.Println("==============================================================================================")

}

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


func (ui *TerminalUI) getNodeId(message string) (*big.Int, string) {
    var nodeID string;
    fmt.Print(message)
    fmt.Scanln(&nodeID)

    NodeID, ok := new(big.Int).SetString(nodeID, 10)

    // inRange := modulo := base.Exp( base, m, nil)

    for !ok { 
		fmt.Println("[ERROR]: Not a valid node id ")
        fmt.Print(message)
        fmt.Scanln(&nodeID)
        NodeID, ok = new(big.Int).SetString(nodeID, 10)
	}
    return NodeID, nodeID
}