package routers

import (
	"github.com/gin-gonic/gin"
    "monitoring-server/controllers"
)

func ConfigurationRoute(route *gin.Engine){
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

}