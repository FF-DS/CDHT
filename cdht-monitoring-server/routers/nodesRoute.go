package routers

import (
    "github.com/gin-gonic/gin"
    "monitoring-server/controllers"
)

func NodesRoute(route *gin.Engine){

	nodesRoutes := route.Group("/nodes")
	nodesController := new(controllers.NodesController)
	
	nodesRoutes.GET("/", nodesController.GetNodes)
	nodesRoutes.POST("/", nodesController.RegisterNode)
	nodesRoutes.GET("/clear", nodesController.ClearNodeData)
}
