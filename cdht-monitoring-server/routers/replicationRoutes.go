package routers

import (
	"monitoring-server/controllers"

	"github.com/gin-gonic/gin"
)

func ReplicationRoute(route *gin.Engine){
	/* 
	Replication related routes
	*/
	replicationRoutes := route.Group("/replication")
	replicationController := new(controllers.ReplicationController)
	replicationRoutes.GET("/" , replicationController.GetReplicasForNode)
	replicationRoutes.POST("/" , replicationController.AddReplicaForNode)
	replicationRoutes.DELETE("/delete-replica" , replicationController.DeleteReplicaFromNode)
	replicationRoutes.GET("/clear" , replicationController.ClearReplicasOfNode)

}