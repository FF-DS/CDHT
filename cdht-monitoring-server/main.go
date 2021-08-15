package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
	"monitoring-server/routers"
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
    router.Use(CORSMiddleware());
	routers.NodesRoute(router) 

    // Handle error response when a route is not defined
    router.NoRoute(func(c *gin.Context) {
        c.JSON(404, gin.H{"message": "Not found"})
    })

    // Init our server
    router.Run( ":" + os.Getenv("PORT") )
}


func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}