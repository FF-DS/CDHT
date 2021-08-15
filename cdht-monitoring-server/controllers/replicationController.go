package controllers

import (
	"github.com/gin-gonic/gin"
)


type ReplicationController struct{}

/* 
This function will get all the existing replicas of a given node
*/
func (replica ReplicationController) GetReplicasForNode(c *gin.Context){
}

/* 
This function will add a given replica into a new or existing replica set
*/
func (replica ReplicationController) AddReplicaForNode(c *gin.Context){
}

/* 
This function deletes a given replica from the set of replicas of a node
*/
func (replica ReplicationController) DeleteReplicaFromNode(c *gin.Context){
}

/* 
This function will clear the replicas for a given node
*/
func (replica ReplicationController) ClearReplicasOfNode(c *gin.Context){
}
