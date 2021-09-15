package controllers

import (
	. "monitoring-server/services"
	"monitoring-server/core"
    "github.com/gin-gonic/gin"
	"context"
    // "log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
    // "go.mongodb.org/mongo-driver/bson/primitive"
    // "time"
	"net/http"
	"strconv"
)

const (
	COMMANDS_COLLECTION_NAME string = "commands"
	RESULTS_COLLECTION_NAME string = "command-results"
	RESULT_LIMIT int64 = 20
)

type CommandDispatcherController struct{
}

/* 
This function will return a specified number of pending commands from the command collections
*/
func (self CommandDispatcherController) GetPendingCommands(c *gin.Context){

	findOptions := options.Find()
	findOptions.SetLimit(RESULT_LIMIT)

	commandCollection := ConnectDB(COMMANDS_COLLECTION_NAME)

	if cursor , err := commandCollection.Find(context.TODO(), findOptions); err == nil{
		
		defer cursor.Close(context.TODO())
		var commands []core.Command

		for cursor.Next(context.TODO()) {
			var command core.Command
			if err := cursor.Decode(&command) ; err == nil{
				commands = append(commands, command)
			}else {
				message := "an error occured while trying to decode pending command"
				GetError(err , message , c)
			}			
		}

		c.JSON(http.StatusOK, gin.H{"message": "Pending commands retrived successfully", "data": commands} )
	
	}else{
		message := "an error occured while trying to fetch pending commands"
		GetError(err , message , c)
	}
}

func (self CommandDispatcherController) GetCommand(c *gin.Context){
	commandCollection := ConnectDB(COMMANDS_COLLECTION_NAME)

	commandId , _ := strconv.Atoi(c.Query("command_id"))

	filter := bson.D{{"_id" , commandId}}

	var command core.Command

	if err := commandCollection.FindOne(context.TODO() , filter).Decode(&command) ; err == nil{
		c.JSON(http.StatusOK , gin.H{"message" : "successfuly retrived the specified command" , "data" : command})
	}else{
		message := "an error occured wile trying to retrive and decode a specified command"
		GetError(err , message , c)
	}
}

/* 
This function will add a command in to the command collection
*/
func (self CommandDispatcherController) AddPendingCommand(c *gin.Context){
	commandCollection := ConnectDB(COMMANDS_COLLECTION_NAME)

	var command core.Command

	if err := c.ShouldBindJSON(&command) ; err == nil{
	
        if result, err := commandCollection.InsertOne(context.TODO(), command) ; err == nil{
			c.JSON(http.StatusOK, gin.H{"message": "Commands queued successfully" , "data" : result})
		}else{
			message := "an error occured while trying to insert a command"
			GetError(err, message , c)
		}

    } else {
		message := "an error occured while trying to bind a command object from the request"
        GetError(err , message , c)
    }
}

/* 
This function will add a batch of commands in to the command collection
*/
func (self CommandDispatcherController) AddPendingCommandsByBatch(c *gin.Context){
	commandCollection := ConnectDB(COMMANDS_COLLECTION_NAME)

	var commands []interface{}

	if err := c.ShouldBindJSON(&commands) ; err == nil{
	
        if result, err := commandCollection.InsertMany(context.TODO(), commands) ; err == nil{
			c.JSON(http.StatusOK, gin.H{"message": "Commands queued successfully" , "data" : result})
		}else{
			message := "an error occured while trying to insert many command"
			GetError(err, message , c)
		}
            
    } else {
        message := "an error occured while trying to bing a command object from the request"
        GetError(err , message , c)
    }
}

/* 
This function will return results for a specific command
*/
func (self CommandDispatcherController) GetCommandResultReports(c *gin.Context){
	
	commandResultsCollection := ConnectDB(RESULTS_COLLECTION_NAME)

	findOptions := options.Find()
	findOptions.SetLimit(RESULT_LIMIT)


	if cursor , err := commandResultsCollection.Find(context.TODO() , bson.M{} , findOptions ) ; err == nil {
		defer cursor.Close(context.TODO())
		var commandResults []core.CommandResult


		for cursor.Next(context.TODO()) {
			var commandRes core.CommandResult
			if err := cursor.Decode(&commandRes) ; err == nil {
				commandResults = append(commandResults, commandRes)
			}else{
				message := "an error occured while trying to decode command result"
				GetError(err , message , c)
			}

		}

		c.JSON(http.StatusOK, gin.H{"message": "Command Results retrived successfully", "data": commandResults} )

	}else{
		message := "an error occured while trying to fetch command results"
		GetError(err , message , c)
	}

}

/* 
This function will add an intermideate result into the command_response collection
*/
func (self CommandDispatcherController) AddCommandResponseReport(c *gin.Context){
	commandResultsCollection := ConnectDB("command-results")

	var commandRes core.CommandResult

	if err := c.ShouldBindJSON(&commandRes) ; err == nil{
	
        if result, err := commandResultsCollection.InsertOne(context.TODO(), commandRes) ; err == nil{
			c.JSON(http.StatusOK, gin.H{"message": "Command Result added successfully" , "data" : result})
		}else{
			message := "an error occured while trying to insert a command result"
			GetError(err, message , c)
		}
    
    } else {
		message := "an error occured while trying to bind a command result object from the request"
        GetError(err , message , c)
    }
}


/* 
This function will clear the commands collection , thread lightly!!!!!!!!!!!!
*/
func (self CommandDispatcherController) ClearCommandsCollection(c *gin.Context){

}


/* 
This function will clear the command_results collection , thread lightly!!!!!!!!!!!!
*/
func (self CommandDispatcherController) ClearCommandResultsCollection(c *gin.Context){

}