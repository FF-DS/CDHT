package routers

import (
    "github.com/gin-gonic/gin"
    "monitoring-server/controllers"
)

func DashbaordRoutes(route *gin.Engine){

	dashboardRoute := route.Group("/dashboard")
	dashboardController := new(controllers.DashboardController)
	
	dashboardRoute.GET("/", dashboardController.GetNodes)
	dashboardRoute.POST("/", dashboardController.RegisterNode)
}
