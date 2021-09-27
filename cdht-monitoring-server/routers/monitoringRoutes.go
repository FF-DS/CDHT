package routers

import (
	"monitoring-server/controllers"

	"github.com/gin-gonic/gin"
)

func MonitoringRoute(route *gin.Engine){
	/* 
	Monitoring related routes
	*/
	monitoringRoutes := route.Group("/monitoring")
	monitoringController := new(controllers.MonitoringController)
	monitoringRoutes.POST("/stats" , monitoringController.GetStatisticsForNode)
}