package controllers

import (
	"github.com/gin-gonic/gin"
	// . "monitoring-server/util"
)


type ReplicationController struct{}

/* 
This function will get all the existing replicas of a given node
*/
func (replica ReplicationController) GetReplicasForNode(c *gin.Context){
	/* 
	get the parent node id, then filter the replicas collection using 
	that id to get the children replica nodes for the given node
	*/
}

/* 
This function will add a given replica into a new or existing replica set
*/
func (replica ReplicationController) AddReplicaForNode(c *gin.Context){
	/* 
	take the node id of the parent node and a replica objec/construct one 
	and insert into the collection
	*/
}

/* 
This function deletes a given replica from the set of replicas of a node
*/
func (replica ReplicationController) DeleteReplicaFromNode(c *gin.Context){
	/* 
	just take the replica node id to delete and well, delete it i guess
	*/
}

/* 
This function will clear the replicas for a given node
*/
func (replica ReplicationController) ClearReplicasOfNode(c *gin.Context){
}
