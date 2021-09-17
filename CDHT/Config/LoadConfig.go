package Config

import (
	"github.com/tkanos/gonfig"
	"time"
	"net/http"
    "io/ioutil"
    "encoding/json"
)


type Config struct {
	configuration *Configuration
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
			config.configuration.CopyConfiguration( serverConfig )
		}
	}
}


func (config *Config) UpdateFromFile() *Configuration {
	configuration := Configuration{}
	gonfig.GetConf("./cdht-config.json", &configuration)
	config.configuration.CopyConfiguration( &configuration )
	config.configuration.ValidateConfig()

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