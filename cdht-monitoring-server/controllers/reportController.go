package controllers

import (
    "github.com/gin-gonic/gin"
    . "monitoring-server/services"
    . "monitoring-server/util"
	. "monitoring-server/core"
    "go.mongodb.org/mongo-driver/bson"
    "context"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
    "log"
)


// type NodesController struct{}

type ReportController struct{}

func (report ReportController) Report(c *gin.Context){
    var success  = map[string]string{}
    success["message"] = "this is from report"
    c.JSON(200 , success)
}

/* 
this function will accept some filter params and will 
return the number of specified reports after applying the required filter
 */
func (report *ReportController) GetPackets(c *gin.Context){
    // TODO
    /* 
        inquire mongodb for the specified number of requests and return them sorted by time

        This function must filter value by :- date  , message type , node_id | filter values will be
        encoded in the request body
    */

    collection := ConnectDB("packets")

    cursor , err := collection.Find(context.TODO() , bson.M{})

    if err != nil{
        message := "some error message to edit later"

        GetError(err , message , c)
        return
    }

    defer cursor.Close(context.TODO())
    var packets []NormalPacket

    for cursor.Next(context.TODO()){
        var packet NormalPacket

        cursor.Decode(&packet)
    }

    if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
 
    c.JSON(200, packets )

}

func (report *ReportController) GetTestPackets(c *gin.Context){
    // TODO
    /* 
        inquire mongodb for the specified number of requests and return them sorted by time
    */

    collection := ConnectDB("test_packets")

    cursor , err := collection.Find(context.TODO() , bson.M{})

    if err != nil{
        message := "some error message to edit later"

        GetError(err , message , c)
        return
    }

    defer cursor.Close(context.TODO())
    var test_packets []NormalPacket

    for cursor.Next(context.TODO()){
        var test_packet NormalPacket

        cursor.Decode(&test_packet)
    }

    if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
 
    c.JSON(200, test_packets )

}

/* 
this function takes a packekt and send it to the mongo instance

*/
func (report *ReportController) PostPacket(c *gin.Context){
    collection := ConnectDB("packets")

    var packet NormalPacket

    if err := c.ShouldBindJSON(&packet); err == nil {
        packet.ID = primitive.NewObjectID()
        packet.CreatedDate = time.Now()

        result, err := collection.InsertOne(context.TODO(), packet)
    
        if err != nil {
            message := "some error message to edit later"

            GetError(err, message , c)
        }
    
        c.JSON(200, gin.H{"data": result})
    } else {
        c.JSON(401, gin.H{"error": err.Error()})
    }
    
}

func (report ReportController) PostTestPacket(c *gin.Context){
    collection := ConnectDB("test_packets")

    var test_packet NormalPacket

    if err := c.ShouldBindJSON(&test_packet); err == nil {
        test_packet.ID = primitive.NewObjectID()
        test_packet.CreatedDate = time.Now()

        result, err := collection.InsertOne(context.TODO(), test_packet)
    
        if err != nil {
            message := "some error message to edit later"

            GetError(err, message , c)
        }
    
        c.JSON(200, gin.H{"data": result})
    } else {
        c.JSON(401, gin.H{"error": err.Error()})
    }
}

func (report ReportController) ClearNormalPacketsCollection(c *gin.Context){

}

func (report ReportController) ClearTestPacketsCollection(c *gin.Context){
    
}