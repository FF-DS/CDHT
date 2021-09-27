package controllers

import (
	"context"
	"fmt"
	"monitoring-server/core"
	"monitoring-server/services"
	"monitoring-server/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "log"
	// "time"
	// "fmt"
	// "strconv"
)

/*
This function will return a specified number of pending commands from the command collections
*/
func (cmd CommandDispatcherController) GetPendingCommands(c *gin.Context){

	commandCollection := services.ConnectDB(COMMANDS_COLLECTION_NAME)

	type RequestBody struct{
		Limit int64 `bson:"limit" json:"limit"`
	}
	
	var request RequestBody

	findOptions := options.Find()
	findOptions.SetSort(bson.D{primitive.E{Key:"created_date",Value:  -1}})
	
	if err := c.ShouldBindJSON(&request) ; err == nil{

		findOptions.SetLimit(request.Limit)
		filter := bson.D{{}}

		if cursor , err := commandCollection.Find(context.TODO() , filter,  findOptions); err == nil{
			
			defer cursor.Close(context.TODO())
			var commands []core.ToolCommand
	
			for cursor.Next(context.TODO()) {
				var command core.ToolCommand
				if err := cursor.Decode(&command) ; err == nil{
					commands = append(commands, command)
				}else {
					message := "an error occured while trying to decode pending command"
					util.GetError(err , message , c)
				}			
			}
	
	
			c.JSON(http.StatusOK, gin.H{"message": "Pending commands retrived successfully", "data": commands} )
		
		}else{
			message := "an error occured while trying to fetch pending commands"
			util.GetError(err , message , c)
		}
	}else{
		message := "an error occured while trying to bind request object"
        util.GetError(err , message , c)
	}

}

func (cmd CommandDispatcherController) GetCommand(c *gin.Context){

	commandCollection := services.ConnectDB(COMMANDS_COLLECTION_NAME)

	type RequestBody struct{
		CommandId primitive.ObjectID `json:"command_id" bson:"command_id"`
	}

	
	var request RequestBody
	
	var command core.ToolCommand
	
	if err := c.ShouldBindJSON(&request) ; err == nil{
		filter := bson.D{primitive.E{Key:"_id" , Value :request.CommandId}}
		if err := commandCollection.FindOne(context.TODO() , filter).Decode(&command) ; err == nil{
			c.JSON(http.StatusOK , gin.H{"message" : "successfuly retrived the specified command" , "data" : command})
		}else{
			message := "an error occured wile trying to retrive and decode a specified command"
			util.GetError(err , message , c)
		}
	}else{
		message := "an error occured while trying to bind request object "
        util.GetError(err , message , c)
	}
}

/* 
This function will add a command in to the command collection
*/
func (cmd CommandDispatcherController) AddPendingCommand(c *gin.Context){
	commandCollection := services.ConnectDB(COMMANDS_COLLECTION_NAME)

	var command core.ToolCommand

	if err := c.ShouldBindJSON(&command) ; err == nil{

		command.OperationID = primitive.NewObjectID()
	
        if result, err := commandCollection.InsertOne(context.TODO(), command) ; err == nil{
			c.JSON(http.StatusOK, gin.H{"message": "Commands queued successfully" , "data" : gin.H{ "result": result , "operation_id" : command.OperationID}})
		}else{
			message := "an error occured while trying to insert a command"
			util.GetError(err, message , c)
		}

    } else {
		message := "an error occured while trying to bind a command object from the request"
        util.GetError(err , message , c)
    }
}

/* 
This function will add a batch of commands in to the command collection
*/
func (cmd CommandDispatcherController) AddPendingCommandsByBatch(c *gin.Context){
	commandCollection := services.ConnectDB(COMMANDS_COLLECTION_NAME)

	var commands []interface{}

	if err := c.ShouldBindJSON(&commands) ; err == nil{
	
        if result, err := commandCollection.InsertMany(context.TODO(), commands) ; err == nil{
			c.JSON(http.StatusOK, gin.H{"message": "Commands queued successfully" , "data" : result})
		}else{
			message := "an error occured while trying to insert many command"
			util.GetError(err, message , c)
		}
            
    } else {
        message := "an error occured while trying to bing a command object from the request"
        util.GetError(err , message , c)
    }
}

/* 
This function will return results for a specific command
*/
func (cmd CommandDispatcherController) GetCommandResultReports(c *gin.Context){

	/* 
	maybe add some sort of filtering if needed to only show a trace of reports for a specific command
	or refactor out to a different method
	*/
	
	commandResultsCollection := services.ConnectDB(RESULTS_COLLECTION_NAME)

	type RequestBody struct{
		Limit int64 `bson:"limit" json:"limit"`
	}

	var request RequestBody

	findOptions := options.Find()

	if err := c.ShouldBindJSON(&request) ; err == nil{

		findOptions.SetLimit(request.Limit)
	
		if cursor , err := commandResultsCollection.Find(context.TODO() , bson.M{} , findOptions ) ; err == nil {
			defer cursor.Close(context.TODO())
			var commandResults []core.ToolCommand
	
	
			for cursor.Next(context.TODO()) {
				var commandRes core.ToolCommand
				if err := cursor.Decode(&commandRes) ; err == nil {
					commandResults = append(commandResults, commandRes)
				}else{
					message := "an error occured while trying to decode command result"
					util.GetError(err , message , c)
				}
	
			}
	
			c.JSON(http.StatusOK, gin.H{"message": "Command Results retrived successfully", "data": commandResults} )
	
		}else{
			message := "an error occured while trying to fetch command results"
			util.GetError(err , message , c)
		}
	}else{
		message := "an error occured while trying to bind request object "
        util.GetError(err , message , c)
	}
}

/* 
This function will add an intermideate result into the command_response collection
*/
func (cmd CommandDispatcherController) AddCommandResponseReport(c *gin.Context){
	commandResultsCollection := services.ConnectDB(RESULTS_COLLECTION_NAME)

	var commandRes core.ToolCommand

	if err := c.ShouldBindJSON(&commandRes) ; err == nil{

		fmt.Println(commandRes.OperationID)

		if message := validateCommandResult(commandRes , c) ; message != "VALID"{
			return 
		}
	
        if result, err := commandResultsCollection.InsertOne(context.TODO(), commandRes) ; err == nil{
			c.JSON(http.StatusOK, gin.H{"message": "Command Result added successfully" , "data" : result})
		}else{
			message := "an error occured while trying to insert a command result"
			util.GetError(err, message , c)
		}
    
    } else {
		message := "an error occured while trying to bind a command result object from the request"
        util.GetError(err , message , c)
    }
}

func validateCommandResult(commandObject core.ToolCommand , c *gin.Context) string{
	if commandObject.OperationID ==  primitive.NilObjectID {
		message := "OperationId must be set before result is sent to the collection"
        util.SendWarning( message , c)
		return message
	}

	return "VALID"
}


/* 
This function will clear the commands collection , thread lightly!!!!!!!!!!!!
*/
func (cmd CommandDispatcherController) ClearCommandsCollection(c *gin.Context){

}


/* 
This function will clear the command_results collection , thread lightly!!!!!!!!!!!!
*/
func (cmd CommandDispatcherController) ClearCommandResultsCollection(c *gin.Context){

}