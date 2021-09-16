package util

import (
	"log"
	"net/http"
    "github.com/gin-gonic/gin"
)

func GetError(err error, message string, c *gin.Context) {

	log.Println(err.Error())

    c.JSON(http.StatusInternalServerError , gin.H{"message": message , "error" : err.Error() })
}