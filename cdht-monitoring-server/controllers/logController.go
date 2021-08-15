package controllers

import(
	. "monitoring-server/services"
	"monitoring-server/core"
    "github.com/gin-gonic/gin"
	"context"
    "log"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
	"net/http"
)

type LogController struct{}

func (self LogController) GetLogs(c *gin.Context){

	logCollection := ConnectDB("logs")

	cursor , err := logCollection.Find(context.TODO() , bson.M{})

	if err != nil{
		GetError(err , c)
		return
	}

	defer cursor.Close(context.TODO())
    var logs []core.LogEntry


    for cursor.Next(context.TODO()) {
		var logEntry core.LogEntry
		err := cursor.Decode(&logEntry)

		if err != nil {
			log.Fatal(err)
		}

		logs = append(logs, logEntry)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
 
    c.JSON(http.StatusOK, gin.H{"message": "Logs retrived succesfully", "data": logs} )
}

func (self LogController) AddLog(c *gin.Context){
	logCollection := ConnectDB("logs")

	var logEntry core.LogEntry

	if err := c.ShouldBindJSON(&logEntry) ; err == nil{
		logEntry.ID = primitive.NewObjectID()
        logEntry.CreatedDate = time.Now()

        result, err := logCollection.InsertOne(context.TODO(), logEntry)
    
        if err != nil {
            GetError(err, c)
        }
    
        c.JSON(http.StatusOK, gin.H{"message": "Log added successfully" , "data" : result})
    } else {
        c.JSON(401, gin.H{"error": err.Error()})
    }
}