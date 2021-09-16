package routers

import (
	"github.com/gin-gonic/gin"
    "monitoring-server/controllers"
)

func ReportRoute(route *gin.Engine){
	/* 
	Report related routes
	*/
	reportRoutes := route.Group("/report")
	rpeortController := new(controllers.ReportController)
	reportRoutes.GET("/normal" , rpeortController.GetPackets)
	reportRoutes.GET("/test" , rpeortController.GetTestPackets)
	reportRoutes.POST("/normal-packet" , rpeortController.PostPacket)
	reportRoutes.POST("/test-packet" , rpeortController.PostTestPacket)
	reportRoutes.GET("/clear-normal" , rpeortController.ClearNormalPacketsCollection)
	reportRoutes.GET("/clear-test" , rpeortController.ClearTestPacketsCollection)
}