package controllers

import (
	"github.com/gin-gonic/gin"
)

type CommandDispatcherController struct{}

/* 
This function will return a specified number of pending commands from the command collections
*/
func (command CommandDispatcherController) GetPendingCommands(c *gin.Context){

}


/* 
This function will add a command in to the command collection
*/
func (command CommandDispatcherController) AddPendingCommand(c *gin.Context){

}

/* 
This function will return results for a specific command
*/
func (command CommandDispatcherController) GetCommandResultReports(c *gin.Context){

}

/* 
This function will add an intermideate result into the command_response collection
*/
func (command CommandDispatcherController) AddCommandResponseReport(c *gin.Context){

}


/* 
This function will clear the commands collection , thread lightly!!!!!!!!!!!!
*/
func (command CommandDispatcherController) ClearCommandsCollection(c *gin.Context){

}


/* 
This function will clear the command_results collection , thread lightly!!!!!!!!!!!!
*/
func (command CommandDispatcherController) ClearCommandResultsCollection(c *gin.Context){

}