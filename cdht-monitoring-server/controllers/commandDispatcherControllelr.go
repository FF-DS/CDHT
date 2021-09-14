package controllers

import (
	. "monitoring-server/services"
	"monitoring-server/core"
    "github.com/gin-gonic/gin"
	"context"
    "log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
    // "go.mongodb.org/mongo-driver/bson/primitive"
    // "time"
	"net/http"
)

type CommandDispatcherController struct{}

/* 
This function will return a specified number of pending commands from the command collections
*/
func (self CommandDispatcherController) GetPendingCommands(c *gin.Context){

	findOptions := options.Find()
	findOptions.SetLimit(20)

	commandCollection := ConnectDB("commands")

	cursor , err := commandCollection.Find(context.TODO() , bson.M{} , findOptions )

	if err != nil{
		GetError(err , c)
		return
	}

	defer cursor.Close(context.TODO())
    var commands []core.Command


    for cursor.Next(context.TODO()) {
		var command core.Command
		err := cursor.Decode(&command)

		if err != nil {
			log.Fatal(err)
		}

		commands = append(commands, command)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
 
    c.JSON(http.StatusOK, gin.H{"message": "Pending commands retrived successfully", "data": commands} )
}

/* 
This function will add a command in to the command collection
*/
func (self CommandDispatcherController) AddPendingCommand(c *gin.Context){
	commandCollection := ConnectDB("commands")

	var command core.Command

	if err := c.ShouldBindJSON(&command) ; err == nil{
	
        result, err := commandCollection.InsertOne(context.TODO(), command)
    
        if err != nil {
            GetError(err, c)
        }
    
        c.JSON(http.StatusOK, gin.H{"message": "Commands queued successfully" , "data" : result})
    } else {
        c.JSON(401, gin.H{"error": err.Error()})
    }
}


/* 
This function will add a batch of commands in to the command collection
*/
func (self CommandDispatcherController) AddPendingCommandsByBatch(c *gin.Context){
	commandCollection := ConnectDB("commands")

	var commands []interface{}

	if err := c.ShouldBindJSON(&commands) ; err == nil{
	
        result, err := commandCollection.InsertMany(context.TODO(), commands)
    
        if err != nil {
            GetError(err, c)
        }
    
        c.JSON(http.StatusOK, gin.H{"message": "Commands queued successfully" , "data" : result})
    } else {
        c.JSON(401, gin.H{"error": err.Error()})
    }
}

/* 
This function will return results for a specific command
*/
func (self CommandDispatcherController) GetCommandResultReports(c *gin.Context){

	findOptions := options.Find()
	findOptions.SetLimit(20)

	commandResultsCollection := ConnectDB("command-results")

	cursor , err := commandResultsCollection.Find(context.TODO() , bson.M{} , findOptions )

	if err != nil{
		GetError(err , c)
		return
	}

	defer cursor.Close(context.TODO())
    var commandResults []core.CommandResult


    for cursor.Next(context.TODO()) {
		var commandRes core.CommandResult
		err := cursor.Decode(&commandRes)

		if err != nil {
			log.Fatal(err)
		}

		commandResults = append(commandResults, commandRes)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
 
    c.JSON(http.StatusOK, gin.H{"message": "Command Results retrived successfully", "data": commandResults} )

}

/* 
This function will add an intermideate result into the command_response collection
*/
func (self CommandDispatcherController) AddCommandResponseReport(c *gin.Context){
	commandResultsCollection := ConnectDB("command-results")

	var commandRes core.CommandResult

	if err := c.ShouldBindJSON(&commandRes) ; err == nil{
	
        result, err := commandResultsCollection.InsertOne(context.TODO(), commandRes)
    
        if err != nil {
            GetError(err, c)
        }
    
        c.JSON(http.StatusOK, gin.H{"message": "Command Result added successfully" , "data" : result})
    } else {
        c.JSON(401, gin.H{"error": err.Error()})
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