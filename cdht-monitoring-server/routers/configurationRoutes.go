package routers

import (
	"monitoring-server/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigurationRoute(route *gin.Engine){
	/* 
	Configuration related routes
	*/
	configurationRoutes := route.Group("/configuration")
	configurationController := new(controllers.ConfigurationController)
	configurationRoutes.GET("/" , configurationController.GetConfigurationProfiles)
	configurationRoutes.GET("/config-reports" , configurationController.GetConfigurationReportEntries)
	configurationRoutes.GET("/current" , configurationController.GetCurrentConfigurationProfile)
	configurationRoutes.POST("/current" , configurationController.SetCurrentConfigurationProfile)
	configurationRoutes.POST("/add-config-profile" , configurationController.AddConfigurationProfile)
	configurationRoutes.POST("/set-jump-space" , configurationController.SetNodeSpaceBalancing)
	configurationRoutes.POST("/set-replication-level" , configurationController.SetReplicationLevel)
	configurationRoutes.DELETE("/delete-config-profile" , configurationController.DeleteConfigurationProfile)
	configurationRoutes.GET("/clear" , configurationController.ClearConfigurationProfilesCollection)

}