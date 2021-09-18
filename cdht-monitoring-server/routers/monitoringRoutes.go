package routers

import (
	"monitoring-server/controllers"

	"github.com/gin-gonic/gin"
)

func MonitoringRoute(route *gin.Engine){
	/* 
	Monitoring related routes
	*/
	monitoringRouttes := route.Group("/monitoring")
	monitoringController := new(controllers.MonitoringController)
	monitoringRouttes.GET("/stats" , monitoringController.GetStatisticsForNode)
}