package controllers

import (
	"context"
	"monitoring-server/core"
	"monitoring-server/services"
	"monitoring-server/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson/primitive"
	// "time"
	"net/http"
)

func (mon MonitoringController) GetStatisticsForNode(c *gin.Context){
    /* 
    filter report logs for a specified node and send to the caller , 
    maybe add some sort of filter

	maybe select only appropriate log types so that only relevant information 
	is presented to the caller
    */

    logCollection := services.ConnectDB(LOG_COLLECTION_NAME)

	type RequestBody struct{
        /* 
        the data type of node_id must be primitive.ObjectId in production mode
        */
        NodeId string  `bson:"node_id" json:"node_id"`
		Limit string `bson:"limit" json:"limit"`
	}
	
	var request RequestBody

	findOptions := options.Find()
	findOptions.SetSort(bson.D{primitive.E{Key:"created_date",Value:  -1}})

	if err := c.ShouldBindJSON(&request) ; err == nil{
		limit , _ := strconv.ParseInt(request.Limit, 10, 64)
		findOptions.SetLimit(limit)

		
		if cursor , err := logCollection.Find(context.TODO() , bson.D{primitive.E{Key: "node_id" , Value: request.NodeId} , primitive.E{Key: "type" , Value: core.LOG_TYPE_NODE_INFORMATION}} , findOptions) ; err == nil{

			defer cursor.Close(context.TODO())
			var logs []core.LogEntry
		
		
			for cursor.Next(context.TODO()) {
				var logEntry core.LogEntry
				if err := cursor.Decode(&logEntry) ; err == nil{

					logs = append(logs, logEntry)
				}else{
					message := "an error occured while trying to decode log entry"
					util.GetError(err , message , c)
				}
			}

			c.JSON(http.StatusOK, gin.H{"message": "Logs retrived succesfully", "data": logs} )

		}else{
			message := "an error occured while trying to fetch log entries"
			util.GetError(err , message , c)
		}
	
	}else{
		message := "an error occured while trying to bind request object"
        util.GetError(err , message , c)
	}

}
