package controllers

import (
	"context"
	"log"
	"monitoring-server/core"
	"monitoring-server/services"
	"monitoring-server/util"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)





func (h *NodesController) GetNodes(c *gin.Context) {
    collection := services.ConnectDB("nodes")
    
    now := time.Now()
    queryTime := now.Add(-10 * time.Minute) // ten minutes ago

    cur, err := collection.Find(context.TODO(), bson.M{"created_date": bson.M{
        "$gte": primitive.NewDateTimeFromTime( queryTime ),
    }})

	if err != nil {
        message := "some error message to edit later"
		util.GetError(err, message , c)
		return
	}


	defer cur.Close(context.TODO())
    var connectedNodes []core.Node


    for cur.Next(context.TODO()) {
		var connectedNode core.Node
		err := cur.Decode(&connectedNode)

		if err != nil {
			log.Fatal(err)
		}

		connectedNodes = append(connectedNodes, connectedNode)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
 
    c.JSON(200, connectedNodes )
}


func (h *NodesController) RegisterNode(c *gin.Context) {
    collection := services.ConnectDB("nodes")

    var node core.Node
    if err := c.ShouldBindJSON(&node); err == nil {
        node.ID = primitive.NewObjectID()
        node.Created_date = time.Now()

        if node.Node_id == ""{
            c.JSON(422, gin.H{"error": "Node id required"})
            return
        }

        result, err := collection.InsertOne(context.TODO(), node)
    
        if err != nil {
            message := "some error message to edit later"
            util.GetError(err, message , c)
        }
    
        c.JSON(200, gin.H{"message": result})
    } else {
        c.JSON(401, gin.H{"error": err.Error()})
    }

}


func (h *NodesController) ClearNodeData(c *gin.Context) {
    collection := services.ConnectDB("nodes")
    
    now := time.Now()
    oldDataTime := now.Add(-10 * time.Minute) 

    _, err := collection.DeleteMany(context.TODO(),  bson.M{"created_date": bson.M{
        "$lt": primitive.NewDateTimeFromTime( oldDataTime ),
    }})


    if err != nil {
        log.Fatal(err)
        c.JSON(401, gin.H{"error": err.Error()})
    }

    c.JSON(200, gin.H{"message": "clean up performed\n"})
}