package routers

import (
	"monitoring-server/controllers"

	"github.com/gin-gonic/gin"
)

func CommandDispatcherRoute(route *gin.Engine){
	/* 
	Command Dispatcher related routes
	*/
	commandDispatcherRoutes := route.Group("/command-dispatcher")
	commandDispatcherController := new(controllers.CommandDispatcherController)
	commandDispatcherRoutes.GET("/all" , commandDispatcherController.GetPendingCommands)
	commandDispatcherRoutes.GET("/single-command" , commandDispatcherController.GetCommand)
	commandDispatcherRoutes.POST("/" , commandDispatcherController.AddPendingCommand)
	commandDispatcherRoutes.POST("/batch" , commandDispatcherController.AddPendingCommandsByBatch)
	commandDispatcherRoutes.GET("/result" , commandDispatcherController.GetCommandResultReports)
	commandDispatcherRoutes.POST("/result" , commandDispatcherController.AddCommandResponseReport)
	commandDispatcherRoutes.GET("/clear-commands" , commandDispatcherController.ClearCommandsCollection)
	commandDispatcherRoutes.GET("/clear-results" , commandDispatcherController.ClearCommandResultsCollection)
}