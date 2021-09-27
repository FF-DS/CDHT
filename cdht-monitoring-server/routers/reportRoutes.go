package routers

import (
	"monitoring-server/controllers"

	"github.com/gin-gonic/gin"
)

func ReportRoute(route *gin.Engine){
	/* 
	Report related routes
	*/
	reportRoutes := route.Group("/report")
	rpeortController := new(controllers.ReportController)
	reportRoutes.POST("/all" , rpeortController.GetReportEntries)
	reportRoutes.POST("/filtered" , rpeortController.GetFilteredReportEntries)
}