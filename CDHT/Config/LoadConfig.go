package Config

import (
	"github.com/tkanos/gonfig"
	"time"
	"net/http"
    "io/ioutil"
    "encoding/json"
	"cdht/ReportModule"
)


type Config struct {
	configuration *Configuration
	Logger *ReportModule.Logger
}


// ## ---------- init --------------- ##
func (config *Config) LoadConfig() *Configuration {
	configuration := Configuration{}

	// load from file
	gonfig.GetConf("./cdht-config.json", &configuration)
	config.configuration = &configuration
	config.configuration.CopyConfiguration( &Configuration{} )
	
	// load from server
	if serverConfig := config.downloadFromServer(); serverConfig != nil {
		config.configuration.CopyConfiguration( serverConfig )
	}

	config.configuration.ValidateConfig()
	return config.configuration
}



// ## ---------- update --------------- ##
func (config *Config) DownloadConfiguration() {
	for {
		time.Sleep(time.Second * config.configuration.CONFIGURATION_DOWNLOAD_DELAY)
		if serverConfig := config.downloadFromServer(); serverConfig != nil {
			if config.Logger != nil && (serverConfig.REPLICATION_COUNT != config.configuration.REPLICATION_COUNT || serverConfig.Jump_Spacing != config.configuration.Jump_Spacing){
				config.Logger.ConfigToolLog(ReportModule.Log{
					Type: ReportModule.LOG_TYPE_CONFIGURATION_SERVICE,
					OperationStatus: ReportModule.LOG_OPERATION_STATUS_SUCCESS,
					LogLocation: ReportModule.LOG_LOCATION_TYPE_INCOMMING,
					NodeId: config.configuration.GetNodeID(),
					LogBody:"Node id with "+ config.configuration.GetNodeID().String()+" configuration data is updated",
				})
			}
			config.configuration.CopyConfiguration( serverConfig )
		}
	}
}


func (config *Config) UpdateFromFile() *Configuration {
	configuration := Configuration{}
	gonfig.GetConf("./cdht-config.json", &configuration)
	config.configuration.CopyConfiguration( &configuration )
	config.configuration.ValidateConfig()

	if config.Logger != nil && (configuration.REPLICATION_COUNT != config.configuration.REPLICATION_COUNT || configuration.Jump_Spacing != config.configuration.Jump_Spacing){
		config.Logger.ConfigToolLog(ReportModule.Log{
			Type: ReportModule.LOG_TYPE_CONFIGURATION_SERVICE,
			OperationStatus: ReportModule.LOG_OPERATION_STATUS_SUCCESS,
			LogLocation: ReportModule.LOG_LOCATION_TYPE_INCOMMING,
			NodeId: config.configuration.GetNodeID(),
			LogBody:"Node id with "+ config.configuration.GetNodeID().String()+" configuration data is updated",
		})
	}

	return config.configuration
}



// ## -------------- helper -------------- ##

func (config *Config) downloadFromServer()  *Configuration {
	resp, err := http.Get( config.configuration.CONFIGURATION_DOWNLOAD_URL )
    if err != nil {
		return nil
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
		return nil
    }

    var remoteConfig Configuration
    err = json.Unmarshal(body, &remoteConfig)
	return &remoteConfig
}