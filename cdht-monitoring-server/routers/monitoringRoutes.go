package routers

import (
	"github.com/gin-gonic/gin"
    "monitoring-server/controllers"
)

func MonitoringRoute(route *gin.Engine){
	/* 
	Monitoring related routes
	*/
	monitoringRouttes := route.Group("/monitoring")
	monitoringController := new(controllers.MonitoringController)
	monitoringRouttes.GET("/" , monitoringController.Monitor)
}