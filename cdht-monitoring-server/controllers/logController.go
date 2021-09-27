package controllers

import (
	"context"
	"monitoring-server/core"
	"monitoring-server/services"
	"monitoring-server/util"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson/primitive"
	// "time"
	"net/http"
)

func (lg LogController) GetLogs(c *gin.Context){
	
	logCollection := services.ConnectDB(LOG_COLLECTION_NAME)

	type RequestBody struct{
		Limit int64 `bson:"limit" json:"limit"`
	}
	
	var request RequestBody

	findOptions := options.Find()
	findOptions.SetSort(bson.D{primitive.E{Key:"created_date",Value:  -1}})

	if err := c.ShouldBindJSON(&request) ; err == nil{
		findOptions.SetLimit(request.Limit)

		
		if cursor , err := logCollection.Find(context.TODO() , bson.M{} , findOptions) ; err == nil{

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

func (lg LogController) GetLog(c *gin.Context){
	logCollection := services.ConnectDB(LOG_COLLECTION_NAME)

	type RequestBody struct{
		LogId primitive.ObjectID `json:"log_id" bson:"log_id"`
	}

	var request RequestBody
	
	var logEntry core.LogEntry
	
	if err := c.ShouldBindJSON(&request) ; err == nil{
		filter := bson.D{primitive.E{Key:"_id" , Value :request.LogId}}
		if err := logCollection.FindOne(context.TODO() , filter).Decode(&logEntry) ; err == nil{
			c.JSON(http.StatusOK , gin.H{"message" : "successfuly retrived the specified log" , "data" : logEntry})
		}else{
			message := "an error occured wile trying to retrive and decode a specified log"
			util.GetError(err , message , c)
		}
	}else{
		message := "an error occured while trying to bind request object "
        util.GetError(err , message , c)
	}
}

func (lg LogController) AddLog(c *gin.Context){
	logCollection := services.ConnectDB(LOG_COLLECTION_NAME)

	var logEntries []interface{}

	if err := c.ShouldBindJSON(&logEntries) ; err == nil{
	
        if result, err := logCollection.InsertMany(context.TODO(), logEntries) ; err == nil {
			c.JSON(http.StatusOK, gin.H{"message": "Log added successfully" , "data" : result})
		}else{
            message := "an error occured while inserting the batch of logs specified"
			util.GetError(err, message , c)
        }
    } else {
        message := "an error occured while trying to bind request object"
        util.GetError(err , message , c)
    }
}