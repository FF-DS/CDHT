package routers

import (
	"monitoring-server/controllers"

	"github.com/gin-gonic/gin"
)

func LogRoute(route *gin.Engine){
	/* 
	Logs related routes
	*/

	logsRoutes := route.Group("/logs")
	logController := new(controllers.LogController)
	logsRoutes.GET("/" , logController.GetLogs)
	logsRoutes.GET("/single-log" , logController.GetLog)
	logsRoutes.POST("/" , logController.AddLog)
}