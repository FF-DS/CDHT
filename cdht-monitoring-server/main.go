package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
	"monitoring-server/routers"
    "os"
)


func init() {
    if err := godotenv.Load(); err != nil {
        log.Printf("No .env file found")
    }
}

func main() {
    // Init gin router
    router := gin.Default()
	routers.NodesRoute(router) 

    // Handle error response when a route is not defined
    router.NoRoute(func(c *gin.Context) {
        c.JSON(404, gin.H{"message": "Not found"})
    })

    // Init our server
    router.Run( ":" + os.Getenv("PORT") )
}
