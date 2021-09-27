package routers

import (
	"monitoring-server/controllers"

	"github.com/gin-gonic/gin"
)

func TestRoute(route *gin.Engine){
	/* 
	Test related routes
	*/
	testRoutes := route.Group("/test")
	testController := new(controllers.TestController)
	testRoutes.POST("/ping" , testController.RunPingTest)
	testRoutes.POST("/dns-lookup" , testController.RunDNSLookUpTest)
	testRoutes.POST("/hop-count" , testController.RunHopCountTest)
	testRoutes.GET("/filter-results" , testController.FilterTestResults)
}