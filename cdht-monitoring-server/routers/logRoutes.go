package routers

import (
	"github.com/gin-gonic/gin"
    "monitoring-server/controllers"
)

func LogRoute(route *gin.Engine){
	/* 
	Logs related routes
	*/

	logsRoutes := route.Group("/logs")
	logController := new(controllers.LogController)
	logsRoutes.GET("/" , logController.GetLogs)
	logsRoutes.POST("/" , logController.AddLog)
}