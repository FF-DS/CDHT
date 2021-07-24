package controllers

import (
    "github.com/gin-gonic/gin"
	"monitoring-server/services"
	"monitoring-server/core"
    "context"
    "log"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)


type DashboardController struct{}


func (h *DashboardController) GetNodes(c *gin.Context) {
    collection := services.ConnectDB("nodes")
    cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		services.GetError(err, c)
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


    c.JSON(200, gin.H{"connected_nodes": connectedNodes })
}




func (h *DashboardController) RegisterNode(c *gin.Context) {
    collection := services.ConnectDB("nodes")

    var node core.Node
    if err := c.ShouldBindJSON(&node); err == nil {
        node.ID = primitive.NewObjectID()
        node.CreatedDate = time. Now()


        result, err := collection.InsertOne(context.TODO(), node)
    
        if err != nil {
            services.GetError(err, c)
            return
        }
    
        c.JSON(200, gin.H{"message": result})
    } else {
        c.JSON(401, gin.H{"error": err.Error()})
    }

}