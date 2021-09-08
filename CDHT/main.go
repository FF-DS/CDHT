package main

import (
	"github.com/schollz/progressbar/v2"

    "cdht/Applications/TestApplications"
    "cdht/Applications/CDHTNetworkTools"
    "cdht/RoutingModule"
    "cdht/ReportModule"
    "cdht/NetworkTools"
    "cdht/API"

    "strconv"
    "strings"
    "math/big"
    "bufio"
    "fmt"
    
    "time"
    "sync"
    "os"
)



func main() {
    var wg sync.WaitGroup
    
    wg.Add(1)
    go runFirstNode(&wg);
    // go runSecondNode(&wg);
    // go runTestApps()


    wg.Wait()
}


// -------------------- apps ---------------------- # //
func runFirstNode(wg *sync.WaitGroup) {
    // api
    ports := testChangeAppServer()
    apiComm := API.ApiCommunication{
        ChannelSize: 100000,
        PORT : ports,
    }
    apiComm.Init()

    // report
    logMngr := ReportModule.Logger{}
    logMngr.Init()

    // Network Tool
    netTools := NetworkTools.NetworkTool{ Logger: &logMngr}
    netTools.Init(100000)

    // routing
    firstNode := RoutingModule.RoutingTable{ Applications: apiComm.Application, Logger: &logMngr, NetworkTools: netTools.NetworkToolPackets}
    testChangeNodeInfo( &firstNode )
    
    // init node
    firstNode.CreateRing()    
    
    // init api communication
    apiComm.NodeRoutingTable = &firstNode
    apiComm.StartAppServer()
    
    // init network tools
    netTools.RoutingTable = &firstNode
    netTools.RunTools()

    // cdht network tools
    cdhtTools := runCDHTNetworkTools(&firstNode, &apiComm)    
    
    // logo
    fmt.Println(logo)
    
    // just ps bars
    progressBar(100)

    fmt.Println("[Init]: + Initalizing routing tables.... ")
    progressBar(100)

    fmt.Println("[Init]: + Initalizing network tools.... ")
    progressBar(100)


    // user interface
    userUI(&apiComm, &firstNode, &cdhtTools)
    wg.Done()
}


func runSecondNode(wg *sync.WaitGroup) {
    // api
    ports := testChangeAppServer()
    apiComm := API.ApiCommunication{
        ChannelSize: 100000,
        PORT : ports,
    }
    apiComm.Init()

    // report
    logMngr := ReportModule.Logger{}
    logMngr.Init()


    // Network Tool
    netTools := NetworkTools.NetworkTool{ Logger: &logMngr}
    netTools.Init(100000)

    // routing
    secondNode := RoutingModule.RoutingTable{ Applications: apiComm.Application, Logger: &logMngr,  NetworkTools: netTools.NetworkToolPackets }
    testChangeNodeInfo( &secondNode )

    // init node
    secondNode.RunNode()
    
    // init api communication
    apiComm.NodeRoutingTable = &secondNode
    apiComm.StartAppServer()
    
    // init network tools
    netTools.RoutingTable = &secondNode
    netTools.RunTools()
    

    // cdht network tools
    var cdhtTools CDHTNetworkTools.CDHTNetworkTool
    cdhtTools = runCDHTNetworkTools(&secondNode, &apiComm)    

    // logo
    fmt.Println(logo)

    // just ps bars
    progressBar(100)

    fmt.Println("[Init]: + Initalizing routing tables.... ")
    progressBar(100)

    fmt.Println("[Init]: + Initalizing network tools.... ")
    progressBar(100)


    // user interface
    userUI(&apiComm, &secondNode, &cdhtTools)
    wg.Done()
}


func runTestApps(){
    go TestApplications.RunTestTCPApp();
    go TestApplications.RunTestUDPApp();
}


func runCDHTNetworkTools(routingTable *RoutingModule.RoutingTable, api *API.ApiCommunication) CDHTNetworkTools.CDHTNetworkTool {
    pingPort := getInput("Enter Ping tool port: ")
    cdhtTools := CDHTNetworkTools.CDHTNetworkTool{
        AppServerIP: routingTable.NodeInfo().IP_address,
        AppServerPort: api.PORT,
        PingToolListeningPort: pingPort[0],
        ReadCommandDelay: 10000,
        ChannelSize: 10000,
        NodeId: routingTable.NodeInfo().Node_id,
        NodeAddress: routingTable.NodeInfo().IP_address + ":" + routingTable.NodeInfo().Port,
    }
    cdhtTools.Init()

    return cdhtTools
}



//  # ----------------------- UI  ----------------------- # //

func getInput(inputStr string) []string {
    fmt.Print(inputStr)
    scanner := bufio.NewScanner(os.Stdin)
    var strInput string
    if scanner.Scan() {
        strInput = scanner.Text()
    }

    return strings.Split(strInput, " ")
}

func userUI(api *API.ApiCommunication, route *RoutingModule.RoutingTable, cdhtTools *CDHTNetworkTools.CDHTNetworkTool){
    for {
        params := getInput("> ")

        switch params[0] {
            case "close" : 
                return
            case "route" :
                printRoutes(route, params)
            case "lookup" :
                lookUpUI(route, params)
            case "log" :
                logDump(route.Logger, params)
            case "tool":
                testCDHTtool(cdhtTools, params[1])
        }

    }
}



// # ------------------- Route Tool -------------------  # //
func printRoutes(route *RoutingModule.RoutingTable, params []string) {
    if len(params) == 1 {
        route.PrintRoutingInfo()
    }

    if len(params) == 2 {
        count, err := strconv.Atoi(params[1]) 
        if err != nil {
            count = 3
        }
        for i := 0; i < count; i++ {
            time.Sleep(time.Second * 2)
            route.PrintRoutingInfo()
        }
    }
}
// # ------------------- [END] Route Tool -------------------  # //


// # ------------------- LookUp Tool -------------------  # //
func lookUpUI(routingM *RoutingModule.RoutingTable, params []string) {
    var nodeId string
    fmt.Print("Enter Node Id: ")
    fmt.Scanln(&nodeId)

    NodeId, ok := new(big.Int).SetString(nodeId, 10)
    if !ok { fmt.Println("SetString: error node id") }

    succ := routingM.LookUp(NodeId)
    succ.PrintNodeInfo()
    fmt.Println("============ [Trace] ============")

    logged := make(map[string]string)
    for _, node := range succ.NodeTraversalLogs {
        if _, logged := logged[node.Node_id.String()]; logged {
            continue
        }

        node.PrintNodeInfo()
        logged[node.Node_id.String()] = node.Node_id.String()
    }
}
// # ------------------- [END] LookUp Tool -------------------  # //


// # ------------------- Log Tool -------------------  # //
func logDump(logger *ReportModule.Logger, params []string) {
    if len(params) < 2 {
        return 
    }

    switch params[1] {
        case "route":
            printLog( logger.RouteLogs() )
        case "node":
            printLog( logger.NodeLogs() )
        case "network":
            printLog( logger.NetworkToolLogs() )
        case "config":
            printLog( logger.ConfigToolLogs() )
    }

}

func printLog(logs []ReportModule.Log){
    for _, log := range logs {
        fmt.Println(log.ToString())
    }
}

func progressBar(amount time.Duration){
    bar := progressbar.New(10)
    for i := 0; i < 10; i++ {
        bar.Add(1)
        time.Sleep(amount * time.Millisecond)
    }
    fmt.Print("\n")
}
// # ------------------- [END] Log Tool -------------------  # //


// # ------------------- CDHT Network Tool -------------------  # //
func testCDHTtool(cdhtTools *CDHTNetworkTools.CDHTNetworkTool, command string) {
    if cdhtTools ==  nil {
        return
    }
    switch command {
        case "hop":
            hopCountTool(cdhtTools)
        case "lookup":
            lookUpTool(cdhtTools)
        case "ping":
            pingTool(cdhtTools)
        case "log":
            pintCDHTlog(cdhtTools)
    }
} 

func hopCountTool(cdhtTools *CDHTNetworkTools.CDHTNetworkTool){
    var startNodeId, endNodeId string
    fmt.Print("Start Node Port: ")
    fmt.Scanln(&startNodeId)
    fmt.Print("End Node Port: ")
    fmt.Scanln(&endNodeId)

    command := CDHTNetworkTools.ToolCommand{
        Type: CDHTNetworkTools.COMMAND_TYPE_HOP_COUNT,
	    OperationID: "1232",
	    Body: map[string]string{
            "START_NODE_ID" : startNodeId,
            "END_NODE_ID" : endNodeId,
        },
    }
    cdhtTools.DispatchCommands(command)
}


func lookUpTool(cdhtTools *CDHTNetworkTools.CDHTNetworkTool){
    var NodeId string
    fmt.Print("Node Port: ")
    fmt.Scanln(&NodeId)

    command := CDHTNetworkTools.ToolCommand{
        Type: CDHTNetworkTools.COMMAND_TYPE_LOOK_UP,
	    OperationID: "1232",
	    Body: map[string]string{
            "NODE_ID" : NodeId,
        },
    }
    cdhtTools.DispatchCommands(command)
}


func pingTool(cdhtTools *CDHTNetworkTools.CDHTNetworkTool){
    var NodeId string
    fmt.Print("Node Port: ")
    fmt.Scanln(&NodeId)

    command := CDHTNetworkTools.ToolCommand{
        Type: CDHTNetworkTools.COMMAND_TYPE_PING,
	    OperationID: "1232",
	    Body: map[string]string{
            "NODE_ID" : NodeId,
        },
    }
    cdhtTools.DispatchCommands(command)
}

func pintCDHTlog(cdhtTools *CDHTNetworkTools.CDHTNetworkTool){
    for len(cdhtTools.ResultChannel) > 0 {
        command := <- cdhtTools.ResultChannel
        fmt.Println( command.ToString() )
    }
}
// # ------------------- [END] CDHT Network Tool -------------------  # //

// # ------------------- [END] UI TOOOLS -------------------  # //




// ------------------------ test node  infos -------------------- //
func testChangeAppServer() string {
    var port string
    fmt.Print("      [+]: Enter App server Port: ")
    fmt.Scanln(&port)
    return port
}

func testChangeNodeInfo(routingM *RoutingModule.RoutingTable) {
    var nodeId, m, remote_port, port string
    fmt.Print("      [+]: Enter Remote Node Port: ")
    fmt.Scanln(&remote_port)
    
    fmt.Print("      [+]: Enter Port: ")
    fmt.Scanln(&port)
    
    fmt.Print("      [+]: Enter Node Id: ")
    fmt.Scanln(&nodeId)
    
    fmt.Print("      [+]: Enter M: ")
    fmt.Scanln(&m)
    fmt.Println("\n")


    NodeId, ok := new(big.Int).SetString(nodeId, 10)
    if !ok { fmt.Println("SetString: error node id") }

    M, ok := new(big.Int).SetString(m, 10)
    if !ok { fmt.Println("SetString: error m") }

    i, _ := strconv.Atoi(m)

    routingM.RemoteNodeAddr = "127.0.0.1:"+remote_port
    routingM.NodePort = port
    routingM.JumpSpacing = 2
    routingM.FingerTableLength = i
    routingM.Node_id = NodeId
    routingM.M = M
    routingM.IP_address = "127.0.0.1"
}





var logo  string = `
                                                ▄████▄  ▓█████▄  ██░ ██ ▄▄▄█████▓
                                                ▒██▀ ▀█  ▒██▀ ██▌▓██░ ██▒▓  ██▒ ▓▒
                                                ▒▓█    ▄ ░██   █▌▒██▀▀██░▒ ▓██░ ▒░
                                                ▒▓▓▄ ▄██▒░▓█▄   ▌░▓█ ░██ ░ ▓██▓ ░ 
                                                ▒ ▓███▀ ░░▒████▓ ░▓█▒░██▓  ▒██▒ ░ 
                                                ░ ░▒ ▒  ░ ▒▒▓  ▒  ▒ ░░▒░▒  ▒ ░░   
                                                  ░  ▒    ░ ▒  ▒  ▒ ░▒░ ░    ░    
                                                ░         ░ ░  ░  ░  ░░ ░  ░      
                                                ░ ░         ░     ░  ░  ░         
                                                ░         ░
`