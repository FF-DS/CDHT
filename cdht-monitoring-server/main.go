package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
	"monitoring-server/routers"
	. "monitoring-server/middlewares"
    "os"
)

const (
    node_name string = "mann"
)


func init() {
    if err := godotenv.Load(); err != nil {
        log.Printf("No .env file found")
    }
}

func main() {
    // Init gin router
    router := gin.Default()

    initMiddlewares(router)

    initRouters(router)

    // Init our server
    router.Run( ":" + os.Getenv("PORT") )
}

func initMiddlewares(router *gin.Engine){
    router.Use(CORSMiddleware());
}

func initRouters(router *gin.Engine){
    routers.NodesRoute(router)
    routers.CommandDispatcherRoute(router) 
    routers.ConfigurationRoute(router) 
    routers.LogRoute(router) 
    routers.MonitoringRoute(router) 
    routers.ReplicationRoute(router) 
    routers.ReportRoute(router) 
    routers.TestRoute(router) 
    
    // Handle error response when a route is not defined
    router.NoRoute(func(c *gin.Context) {
        c.JSON(404, gin.H{"message": "The specified route is not registered"})
    })
}

