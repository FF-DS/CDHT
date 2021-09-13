package CoreModule

import (
	"github.com/schollz/progressbar/v2"
	"time"
	"fmt"
)

func (core *Core) UpdateApiCommunicationConfig() {
	if core.ApiCommunication == nil {
		return
	}

	core.ApiCommunication.ChannelSize = core.Config.Application_Channel_Size
}

func (core *Core) UpdateLogManagerConfig() {
	if core.LogManager == nil {
		return
	}

	core.LogManager.Configure( core.Config.GetLogConfiguration() )
}

func (core *Core) UpdateRoutingTableInfoConfig() {
	if core.RoutingTableInfo == nil {
		return
	}

	core.RoutingTableInfo.RoutingUpdateDelay = core.Config.Routing_Update_Delay
	core.RoutingTableInfo.SuccessorsTableLength = core.Config.Successors_Table_Length
	core.RoutingTableInfo.FingerTableLength = core.Config.Finger_Table_Length
	core.RoutingTableInfo.JumpSpacing = core.Config.Jump_Spacing
}

func (core *Core) UpdateCdhtNetworkToolsConfig() {
	if core.CdhtNetworkTools == nil {
		return
	}

	core.CdhtNetworkTools.ReadCommandDelay = core.Config.CDHT_API_Communication_Delay
	core.CdhtNetworkTools.URLGetCommandFromServer = core.Config.CDHT_URL_Get_Command_From_Server
	core.CdhtNetworkTools.URLSendCommandResult = core.Config.CDHT_URL_Send_Command_Result
}

func (core *Core) UpdateCdhtTestApplicationConfig() {
	if core.CdhtTestApplication == nil {
		return
	}
	core.CdhtTestApplication.AppName = core.Config.TEST_APP_AppName
	core.CdhtTestApplication.PacketDelay = core.Config.TEST_APP_Packet_Delay
}





func (core *Core) UpdateApplicationConfiguration(){
	for {
		time.Sleep(time.Second * core.Config.CONFIGURATION_DOWNLOAD_DELAY)
		core.UpdateApiCommunicationConfig()
		core.UpdateLogManagerConfig()
		core.UpdateRoutingTableInfoConfig()
		core.UpdateCdhtNetworkToolsConfig()
		core.UpdateCdhtTestApplicationConfig()
	}
}

// # ------------------  [Helper]  ------------------ #

func progressBar(amount time.Duration){
    bar := progressbar.New(10)
    for i := 0; i < 10; i++ {
        bar.Add(1)
        time.Sleep(amount * time.Millisecond)
    }
    fmt.Print("\n")
}