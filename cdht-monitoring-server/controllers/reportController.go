package controllers

import (
	"context"
	"fmt"
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
type FilterRequestBody struct{
    Limit int64 `bson:"limit" json:"limit"`
    NodeId string `bson:"node_id" json:"node_id"`
    OperationStatus string `bson:"operation_status" json:"operation_status"`
    LogLocation string `bson:"log_location" json:"log_location"`
    StartDate string `bson:"start_date" json:"start_date"`
    EndDate string `bson:"end_date" json:"end_date"`
}

func (report ReportController) GetReportEntries(c *gin.Context){
    /* 
    this function will return everything as it is called with out any filters

    but appropriate filter must be applied as all log types are not relevant for the front end
    */
    logCollection := services.ConnectDB(LOG_COLLECTION_NAME)

	type RequestBody struct{
		Limit string `bson:"limit" json:"limit"`
	}
	
	var request RequestBody

	findOptions := options.Find()
	findOptions.SetSort(bson.D{primitive.E{Key:"created_date",Value:  -1}})

	if err := c.ShouldBindJSON(&request) ; err == nil{
		limit , _ := strconv.ParseInt(request.Limit, 10, 64)
		findOptions.SetLimit(limit)

		
		if cursor , err := logCollection.Find(context.TODO() , bson.M{} , findOptions) ; err == nil{

			defer cursor.Close(context.TODO())
			var logs []core.LogEntry
		
		
			for cursor.Next(context.TODO()) {
				var logEntry core.LogEntry
				if err := cursor.Decode(&logEntry) ; err == nil{

					logs = append(logs, logEntry)
				}else{
					message := "an error occured while trying to decode report entry"
					util.GetError(err , message , c)
				}
			}

			c.JSON(http.StatusOK, gin.H{"message": "Reports retrived succesfully", "data": logs} )

		}else{
			message := "an error occured while trying to fetch report entries"
			util.GetError(err , message , c)
		}
	
	}else{
		message := "an error occured while trying to bind request object"
        util.GetError(err , message , c)
	}
}

func (report ReportController) GetFilteredReportEntries(c *gin.Context){
    /* 
    this function takes various filters like nodeId , messageType , date , report count
    and return the result to the user
    */

    logCollection := services.ConnectDB(LOG_COLLECTION_NAME)

	var request FilterRequestBody

	findOptions := options.Find()
	findOptions.SetSort(bson.D{primitive.E{Key:"created_date",Value:  -1}})

	if err := c.ShouldBindJSON(&request) ; err == nil{

       if message :=  validateFilter(request , c); message != "VALID" {
           return
       }

		findOptions.SetLimit(request.Limit)
         filter := bson.D{ primitive.E{Key: "node_id", Value : request.NodeId } ,
          primitive.E{Key: "operation_status", Value : request.OperationStatus } ,
           primitive.E{Key: "log_location", Value : request.LogLocation } , 
            primitive.E{Key: "created_date", Value : bson.M{ "$gt" : request.StartDate}} ,
             primitive.E{Key: "created_date", Value : bson.M{ "$lt" : request.EndDate}}  }


		fmt.Println("The filter is : ")
        fmt.Println(filter)

		
		if cursor , err := logCollection.Find(context.TODO() , filter, findOptions) ; err == nil{

			defer cursor.Close(context.TODO())
			var logs []core.LogEntry
		
		
			for cursor.Next(context.TODO()) {
				var logEntry core.LogEntry
				if err := cursor.Decode(&logEntry) ; err == nil{

					logs = append(logs, logEntry)
				}else{
					message := "an error occured while trying to decode report entry"
					util.GetError(err , message , c)
				}
			}

			c.JSON(http.StatusOK, gin.H{"message": "Reports retrived succesfully", "data": logs} )

		}else{
			message := "an error occured while trying to fetch report entries"
			util.GetError(err , message , c)
		}
	
	}else{
		message := "an error occured while trying to bind request object"
        util.GetError(err , message , c)
	}

}

func validateFilter(requestObject FilterRequestBody , c *gin.Context) string{
    if requestObject.NodeId == "" {
        message := "Please select a specific node before trying to filter reports"
        util.SendWarning( message , c)
        return message 
    }else if requestObject.OperationStatus == "" {
        message := "Please select operation status before trying to filter reports"
        util.SendWarning( message , c)
        return message
    }else if requestObject.LogLocation == "" {
        message := "Please select report location before trying to filter reports"
        util.SendWarning( message , c)
        return message
    }else if requestObject.StartDate == "" {
        message := "Please select a start date before trying to filter reports"
        util.SendWarning( message , c)
        return message
    }else if requestObject.EndDate == "" {
        message := "Please select an end date before trying to filter reports"
        util.SendWarning( message , c)
        return message
    }

    return "VALID"
}