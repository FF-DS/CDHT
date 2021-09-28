package controllers

import (
	"context"
	"monitoring-server/core"
	"monitoring-server/services"
	"monitoring-server/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	// . "monitoring-server/util"
)

// type NodesController struct{}

func (tst TestController) RunPingTest(c *gin.Context){

    /* 
    this function will take the node_id for the ping test from the request body
    and inject it into the command body and set  the request type before queuing on 
    the command collection
    */

    commandCollection := services.ConnectDB(COMMANDS_COLLECTION_NAME)

    type RequestBody struct{
		Body interface{} `bson:"body" json:"body"`
	}

    var request RequestBody
	var command core.ToolCommand

	if err := c.ShouldBindJSON(&request) ; err == nil{

		command.OperationID = primitive.NewObjectID()
        command.Type = core.COMMAND_TYPE_PING
        command.Body = request.Body
	
        if result, err := commandCollection.InsertOne(context.TODO(), command) ; err == nil{
			c.JSON(http.StatusOK, gin.H{"message": "Ping queued successfully" , "data" : gin.H{ "result": result , "operation_id" : command.OperationID}})
		}else{
			message := "an error occured while trying to insert a ping command"
			util.GetError(err, message , c)
		}

    } else {
		message := "an error occured while trying to bind a request object from the request body"
        util.GetError(err , message , c)
    }

}

func (tst TestController) RunDNSLookUpTest(c *gin.Context){
    /* 
    this function will take the node_id for the DNS Lookup test from the request body
    and inject it into the command body and set  the request type before queuing on 
    the command collection
    */

    commandCollection := services.ConnectDB(COMMANDS_COLLECTION_NAME)

    type RequestBody struct{
		Body interface{} `bson:"body" json:"body"`
	}

    var request RequestBody
	var command core.ToolCommand

	if err := c.ShouldBindJSON(&request) ; err == nil{

		command.OperationID = primitive.NewObjectID()
        command.Type = core.COMMAND_TYPE_LOOK_UP
        command.Body = request.Body
	
        if result, err := commandCollection.InsertOne(context.TODO(), command) ; err == nil{
			c.JSON(http.StatusOK, gin.H{"message": "DNS look up command queued successfully" , "data" : gin.H{ "result": result , "operation_id" : command.OperationID}})
		}else{
			message := "an error occured while trying to insert a DNS look up command"
			util.GetError(err, message , c)
		}

    } else {
		message := "an error occured while trying to bind a request object from the request body"
        util.GetError(err , message , c)
    }
}

func (tst TestController) RunHopCountTest(c *gin.Context){
    /* 
    this function will take the source and dest node_id for the hop_count test from the request body
    and inject it into the command body and set  the request type before queuing on 
    the command collection
    */

    commandCollection := services.ConnectDB(COMMANDS_COLLECTION_NAME)

    type RequestBody struct{
		Body interface{} `bson:"body" json:"body"`
	}

    var request RequestBody
	var command core.ToolCommand

	if err := c.ShouldBindJSON(&request) ; err == nil{

		command.OperationID = primitive.NewObjectID()
        command.Type = core.COMMAND_TYPE_HOP_COUNT
        command.Body = request.Body
	
        if result, err := commandCollection.InsertOne(context.TODO(), command) ; err == nil{
			c.JSON(http.StatusOK, gin.H{"message": "Hop count command queued successfully" , "data" : gin.H{ "result": result , "operation_id" : command.OperationID}})
		}else{
			message := "an error occured while trying to insert a Hop count command"
			util.GetError(err, message , c)
		}

    } else {
		message := "an error occured while trying to bind a request object from the request body"
        util.GetError(err , message , c)
    }
}

func (tst TestController) FilterTestResults(c *gin.Context){
    /* 
    this function will filter results of command from the command collection based on the 
    given operation ID from the request body
    */

    commandResultsCollection := services.ConnectDB(RESULTS_COLLECTION_NAME)

	type RequestBody struct{
        Limit string `bson:"limit" json:"limit"`
		OperationId string `json:"operation_id" bson:"operation_id"`
	}

	
	var request RequestBody
	var command core.ToolCommand

    findOptions := options.Find()
    findOptions.SetSort(bson.D{primitive.E{Key:"created_date",Value:  -1}})
	limit , _ := strconv.ParseInt(request.Limit, 10, 64)
    findOptions.SetLimit(limit)
	
	if err := c.ShouldBindJSON(&request) ; err == nil{
		operation_id , _ := primitive.ObjectIDFromHex( request.OperationId);
		filter := bson.D{primitive.E{Key:"operation_id" , Value : operation_id}}
		if err := commandResultsCollection.FindOne(context.TODO() , filter).Decode(&command) ; err == nil{
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
