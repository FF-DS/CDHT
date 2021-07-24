package controllers

import (
    "github.com/gin-gonic/gin"
	"monitoring-server/services"
	"monitoring-server/core"
    "context"
)


type DashboardController struct{}
collection := services.ConnectDB()



func (h *DashboardController) GetNodes(c *gin.Context) {
    
    cur, err := collection("nodes").Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
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

    var node core.Node = Node(  c.Param("ip_address"), c.Param("node_id") )


	result, err := collection.InsertOne(context.TODO(), node)

	if err != nil {
		helper.GetError(err, w)
		return
	}


    c.JSON(200, gin.H{"message": result})
}