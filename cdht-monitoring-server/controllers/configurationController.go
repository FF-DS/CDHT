package controllers

import (
    "github.com/gin-gonic/gin"
)


// type NodesController struct{}

type ConfigurationController struct{}
func (config ConfigurationController) Config(c *gin.Context){
    var success  = map[string]string{}
    success["message"] = "this is from configuration"
    c.JSON(200 , success)
}

/*
This function will return all the configuration profiles created and stored in the collection
  */
func (config ConfigurationController) GetConfigurationProfiles(c *gin.Context){

}

/* 
This function will return the currently used configuration profile
*/
func (config ConfigurationController) GetCurrentConfigurationProfile(c *gin.Context){
    
}

/* 
This function will set a configuration profile selected by the user and notifies the dispatch server
to propagate the changes accordingly
*/
func (config ConfigurationController) SetCurrentConfigurationProfile(c *gin.Context){
    
}


/* 
This function will set the nodespace balancing of the system and notify the command dispatch server to 
to propagate the changes accordingly
*/
func (config ConfigurationController) SetNodeSpaceBalancing(c *gin.Context){
    
}

/* 
This function will store an newly created configuration profile by the user
*/
func (config ConfigurationController) AddConfigurationProfile(c *gin.Context){
    
}

/* 
This function will delete a configuration profile unless its a currently used one
if the currently used one is being requested to be deleted notify the user to set another
profile
*/
func (config ConfigurationController) DeleteConfigurationProfile(c *gin.Context){
    
}

/* 
This function will clear the configuration profiles database , thread lightly!!!!!!!!!!!!
*/
func (config ConfigurationController) ClearConfigurationProfilesCollection(c *gin.Context){
    
}
