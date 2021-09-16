package routers

import (
	"github.com/gin-gonic/gin"
    "monitoring-server/controllers"
)

func TestRoute(route *gin.Engine){
	/* 
	Test related routes
	*/
	testRoutes := route.Group("/test")
	testController := new(controllers.TestController)
	testRoutes.GET("/" , testController.Test)
}