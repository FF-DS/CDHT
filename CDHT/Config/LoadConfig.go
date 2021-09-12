package Config

import (
	"github.com/tkanos/gonfig"
)


type Config struct {
	configuration *Configuration
}

func (config *Config) LoadConfig() Configuration {
	configuration := Configuration{}
	gonfig.GetConf("./cdht-config.json", &configuration)
	config.configuration = &configuration

	return configuration
}


func (config *Config) ValidateConfiguration() {
	config.checkNodeConfiguration()
}

func (config *Config) checkNodeConfiguration() {

}