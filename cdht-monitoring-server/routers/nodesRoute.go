package routers

import (
    "github.com/gin-gonic/gin"
    "monitoring-server/controllers"
)

func NodesRoute(route *gin.Engine){

	nodesRoutes := route.Group("/nodes")
	nodesController := new(controllers.NodesController)
	
	nodesRoutes.GET("/", nodesController.GetNodes)
	nodesRoutes.POST("/", nodesController.RegisterNode)
	nodesRoutes.GET("/clear", nodesController.ClearNodeData)

	/* 
	Configuration related routes
	*/
	configurationRoutes := route.Group("/configuration")
	configurationController := new(controllers.ConfigurationController)
	configurationRoutes.GET("/" , configurationController.GetConfigurationProfiles)
	configurationRoutes.GET("/current" , configurationController.GetCurrentConfigurationProfile)
	configurationRoutes.POST("/current" , configurationController.SetCurrentConfigurationProfile)
	configurationRoutes.POST("/add-config-profile" , configurationController.AddConfigurationProfile)
	configurationRoutes.POST("/set-jump-space" , configurationController.SetNodeSpaceBalancing)
	configurationRoutes.DELETE("/delete-config-profile" , configurationController.DeleteConfigurationProfile)
	configurationRoutes.GET("/clear" , configurationController.ClearConfigurationProfilesCollection)


	/* 
	Test related routes
	*/
	testRoutes := route.Group("/test")
	testController := new(controllers.TestController)
	testRoutes.GET("/" , testController.Test)

	/* 
	Report related routes
	*/
	reportRoutes := route.Group("/report")
	rpeortController := new(controllers.ReportController)
	reportRoutes.GET("/normal" , rpeortController.GetPackets)
	reportRoutes.GET("/test" , rpeortController.GetTestPackets)
	reportRoutes.POST("/normal-packet" , rpeortController.PostPacket)
	reportRoutes.POST("/test-packet" , rpeortController.PostTestPacket)
	reportRoutes.GET("/clear-normal" , rpeortController.ClearNormalPacketsCollection)
	reportRoutes.GET("/clear-test" , rpeortController.ClearTestPacketsCollection)


	/* 
	Monitoring related routes
	*/
	monitoringRouttes := route.Group("/monitoring")
	monitoringController := new(controllers.MonitoringController)
	monitoringRouttes.GET("/" , monitoringController.Monitor)

	/* 
	Command Dispatcher related routes
	*/
	commandDispatcherRoutes := route.Group("/command-dispatcher")
	commandDispatcherController := new(controllers.CommandDispatcherController)
	commandDispatcherRoutes.GET("/" , commandDispatcherController.GetPendingCommands)
	commandDispatcherRoutes.POST("/" , commandDispatcherController.AddPendingCommand)
	commandDispatcherRoutes.POST("/batch" , commandDispatcherController.AddPendingCommandsByBatch)
	commandDispatcherRoutes.GET("/result" , commandDispatcherController.GetCommandResultReports)
	commandDispatcherRoutes.POST("/result" , commandDispatcherController.AddCommandResponseReport)
	commandDispatcherRoutes.GET("/clear-commands" , commandDispatcherController.ClearCommandsCollection)
	commandDispatcherRoutes.GET("/clear-results" , commandDispatcherController.ClearCommandResultsCollection)
	
	/* 
	Replication related routes
	*/
	replicationRoutes := route.Group("/replication")
	replicationController := new(controllers.ReplicationController)
	replicationRoutes.GET("/" , replicationController.GetReplicasForNode)
	replicationRoutes.POST("/" , replicationController.AddReplicaForNode)
	replicationRoutes.DELETE("/delete-replica" , replicationController.DeleteReplicaFromNode)
	replicationRoutes.GET("/clear" , replicationController.ClearReplicasOfNode)

	/* 
	Logs related routes
	*/

	logsRoutes := route.Group("/logs")
	logController := new(controllers.LogController)
	logsRoutes.GET("/" , logController.GetLogs)
	logsRoutes.POST("/" , logController.AddLog)
}
