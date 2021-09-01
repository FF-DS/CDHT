package main

import (
    "cdht/RoutingModule"
    "cdht/API"
    "cdht/NetworkModule"
    "cdht/Util"
    "time"
    "fmt"
    "math/big"
    "net"
    "strconv"
    "sync"
    "cdht/ReportModule"
    "os"
    "bufio"
    "strings"
	"github.com/schollz/progressbar/v2"
)



func main() {
    var wg sync.WaitGroup
    
    wg.Add(1)
    go runFirstNode(&wg);
    // go runSecondNode(&wg);
    // go runTestApp(&wg);

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

    // routing
    firstNode := RoutingModule.RoutingTable{ Applications: apiComm.Application, Logger: &logMngr }
    testChangeNodeInfo( &firstNode )

    // logo
    fmt.Println(logo)

    firstNode.CreateRing()
    apiComm.NodeRoutingTable = &firstNode
    
    apiComm.StartAppServer()
    progressBar(5)

    fmt.Println("[Init]: + Initalizing routing tables.... ")
    progressBar(10)


    userUI(&apiComm, &firstNode)
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


    // routing
    secondNode := RoutingModule.RoutingTable{ Applications: apiComm.Application, Logger: &logMngr }
    testChangeNodeInfo( &secondNode )

    // logo
    fmt.Println(logo)

    secondNode.RunNode()
    apiComm.NodeRoutingTable = &secondNode
    
    apiComm.StartAppServer()
    progressBar(5)

    fmt.Println("[Init]: + Initalizing routing tables.... ")
    progressBar(10)

    userUI(&apiComm, &secondNode)
    wg.Done()
}


func runTestApp(wg *sync.WaitGroup){
    reqObj1 := Util.RequestObject{ Type: "TEST_TYPE_1", RequestID: "REQ_1", AppName: "TEST_APP", AppID: 1}
    appPort1 := testAppAddress("Test App 1", &reqObj1)


    reqObj2 := Util.RequestObject{ Type: "TEST_TYPE_2", RequestID: "REQ_2", AppName: "TEST_APP", AppID: 1}
    appPort2 := testAppAddress("Test App 2", &reqObj2)

    go testApp(reqObj1, appPort1, "Test App 1")
    go testApp(reqObj2, appPort2, "Test App 2")

    time.Sleep(time.Minute * 300)
    wg.Done()
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


func userUI(api *API.ApiCommunication, route *RoutingModule.RoutingTable){
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
        }

    }
}


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


func lookUpUI(routingM *RoutingModule.RoutingTable, params []string) {
    var nodeId string
    fmt.Print("Enter Node Id: ")
    fmt.Scanln(&nodeId)

    NodeId, ok := new(big.Int).SetString(nodeId, 10)
    if !ok { fmt.Println("SetString: error node id") }

    succ := routingM.LookUp(NodeId)
    succ.PrintNodeInfo()
    fmt.Println("----------- [Trace] ------------")

    logged := make(map[string]string)
    for _, node := range succ.NodeTraversalLogs {
        if _, logged := logged[node.Node_id.String()]; logged {
            continue
        }

        node.PrintNodeInfo()
        logged[node.Node_id.String()] = node.Node_id.String()
    }
}


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

func progressBar(secs time.Duration){
    diff := secs/10
    bar := progressbar.New(10)
    for i := 0; i < 10; i++ {
        bar.Add(1)
        time.Sleep(diff * time.Second)
    }
    fmt.Print("\n")
}
//  # ----------------------- TEST  APPLICATION ----------------------- # //

func testApp(reqObj Util.RequestObject, port string, appName string) {
    appConn := NetworkModule.NewNetworkManager("127.0.0.1", port)

    res := appConn.Connect("TCP", func(connection interface{}){
        if connection, ok := connection.(net.Conn); ok { 
            netChannel := NetworkModule.NetworkChannel{ Connection: connection, ChannelSize: 100000 }
            netChannel.Init()
    
            // register app name
            netChannel.SendToSocket(Util.RequestObject{
                AppName: "TEST_APP",
            })        


            fmt.Println("["+appName+"]:+ Connected.")

            for {
                select {
                    case netReqObj := <- netChannel.ReqChannel:
				    	fmt.Printf("[PACKET-RECEIVED][%s][Node-ID][%s]: Packet: %s \n", appName, reqObj.SenderNodeId.String(), netReqObj)
				    default:
                        time.Sleep(time.Millisecond*500)
					    status := netChannel.SendToSocket(reqObj)

                        if !status {
                            fmt.Println("[API] Unable to send")
                        }
			    }
		    }
        }else {
            fmt.Println("["+appName+"]:+ Unable to Connect.")
        }
    })

    if !res {
        fmt.Println("["+appName+"]:+ Unable to Create Soc.")
    }
}


func testAppAddress(appName string, reqObject *Util.RequestObject) string {
    var senderNodeId, recvNodeId, appServerPort string
    var ok bool
    
    fmt.Println(" ------------------- ["+appName+"] ------------------- ")
    
    fmt.Print("      [+]: Enter Sender(Connecting) Application Server Port: ")
    fmt.Scanln(&appServerPort)
    
    fmt.Print("      [+]: Enter Sender(Connecting) Node Id: ")
    fmt.Scanln(&senderNodeId)
    
    fmt.Print("      [+]: Enter Reciever Node Id: ")
    fmt.Scanln(&recvNodeId)
    
    fmt.Println(" ------------------------------------------------------ ")

    reqObject.SenderNodeId, ok = new(big.Int).SetString(senderNodeId, 10)
    if !ok { fmt.Println("SetString: error node id") }

    reqObject.ReceiverNodeId, ok = new(big.Int).SetString(recvNodeId, 10)
    if !ok { fmt.Println("SetString: error m") }

    reqObject.RequestBody = "THIS IS REQ BODY FROM NODE: " + senderNodeId + " TO NODE " + recvNodeId

    return appServerPort
}
// ------------------------ end [test  app] -------------------- //




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