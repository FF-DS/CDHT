package services

import (
	"context"
	// "encoding/json"
	// "fmt"
	"log"
	"net/http"
	"os"
    "time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
    "github.com/gin-gonic/gin"
)




func ConnectDB(collection_name string) *mongo.Collection {
	clientOptions := options.Client().ApplyURI( os.Getenv("DATABASE_CONN_URL") )

	// Connect to MongoDB
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Connected to MongoDB!")

	
	database_name := os.Getenv("DATABASE_NAME")

	if string( os.Getenv("APP_STATE") ) != "Prod" {
		database_name = os.Getenv("TEST_DATABASE_NAME") 
	}

	collection := client.Database( database_name ).Collection(collection_name)
	return collection
}



func DropCollection(collection_name string) bool {

	// Set client options
	clientOptions := options.Client().ApplyURI( os.Getenv("DATABASE_CONN_URL") )

	// Connect to MongoDB
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Connected to MongoDB!")

	database_name := os.Getenv("DATABASE_NAME")

	if string( os.Getenv("APP_STATE") ) != "Prod" {
		database_name = os.Getenv("TEST_DATABASE_NAME") 
	}

	collection := client.Database( database_name ).Collection(collection_name)
	
	if err = collection.Drop(ctx); err != nil {
		log.Fatal(err)
		return false
	}
	
	return true
}



type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}




func GetError(err error, message string, c *gin.Context) {

	log.Println(err)

    c.JSON(http.StatusInternalServerError , gin.H{"message": message , "error" : err })
}