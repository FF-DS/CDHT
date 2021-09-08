package CDHTNetworkTools

import (
    "net/http"
    "io/ioutil"
    "encoding/json"
    "bytes"

	"math/big"
	"time"
)


type CDHTNetworkTool struct {
	AppServerIP string
	AppServerPort string
	PingToolListeningPort string
	ReadCommandDelay time.Duration
	ChannelSize int
	NodeId *big.Int
	NodeAddress string

	ResultChannel chan ToolCommand
	hopCountTool *HopCountTool
	lookUpTool *LookUpTool
	pingTool *PingTool
}



// init env & network apps
func (netTool *CDHTNetworkTool) Init() {
	netTool.ResultChannel = make(chan ToolCommand, netTool.ChannelSize)

	netTool.hopCountTool = &HopCountTool{
		AppServerIP: netTool.AppServerIP, AppServerPort: netTool.AppServerPort, 
		ChannelSize: netTool.ChannelSize, NodeId: netTool.NodeId, 
		NodeAddress: netTool.NodeAddress, ResultChannel: netTool.ResultChannel,
	}
	netTool.hopCountTool.Init()

	netTool.lookUpTool = &LookUpTool{
		AppServerIP: netTool.AppServerIP, AppServerPort: netTool.AppServerPort, 
		ChannelSize: netTool.ChannelSize, NodeId: netTool.NodeId, 
		NodeAddress: netTool.NodeAddress, ResultChannel: netTool.ResultChannel,
	}
	netTool.lookUpTool.Init()

	netTool.pingTool = &PingTool{
		AppServerIP: netTool.AppServerIP, AppServerPort: netTool.AppServerPort, 
		ChannelSize: netTool.ChannelSize, NodeId: netTool.NodeId, ToolListeningPort: netTool.PingToolListeningPort,
		NodeAddress: netTool.NodeAddress, ResultChannel: netTool.ResultChannel,
	}
	netTool.pingTool.Init()

	go netTool.runApiCommunicationTools();
}


func (netTool *CDHTNetworkTool) runApiCommunicationTools(){
	go func(){
		for {
			time.Sleep(time.Millisecond * netTool.ReadCommandDelay )
			netTool.fetchCommands()
		}
	}()

	go func(){
		for {
			time.Sleep(time.Millisecond * netTool.ReadCommandDelay )
			netTool.sendReportToServer()
		}
	}()
}



// dispatch commands to different tools
func (netTool *CDHTNetworkTool) DispatchCommands(command ToolCommand){
	if command.Type == COMMAND_TYPE_HOP_COUNT {
		netTool.hopCountTool.CommandChannel <- command
	}else if command.Type == COMMAND_TYPE_LOOK_UP {
		netTool.lookUpTool.CommandChannel <- command
	}else if command.Type == COMMAND_TYPE_PING {
		netTool.pingTool.CommandChannel <- command
	}
}


// -------------- [API]: Communicate with the server -------------- //

// fetch command from CDHT Monitoring Server
func (netTool *CDHTNetworkTool)  fetchCommands(){
	resp, err := http.Get(URL_GET_COMMAND_FROM_SERVER)

    if err != nil {
		return 
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
		return 
    }
    
    var commands []ToolCommand
    err = json.Unmarshal(body, &commands)
    
    if err != nil {
		return 
	}

    for _, command := range commands {
		netTool.DispatchCommands(command)
    }
}


// report result to CDHT Monitoring server
func (netTool *CDHTNetworkTool)  sendReportToServer(){
	command := <- netTool.ResultChannel
	postBody, err := json.Marshal( command )

	if err != nil {
		netTool.ResultChannel <- command
		return 
    }
	
    responseBody := bytes.NewBuffer(postBody)
    resp, err := http.Post(URL_SEND_COMMAND_RESULT, "application/json", responseBody)
    
    if err != nil {
		netTool.ResultChannel <- command
		return 
    }
    
   resp.Body.Close();
}