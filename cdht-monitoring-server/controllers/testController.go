package controllers

import (
    "github.com/gin-gonic/gin"
)


// type NodesController struct{}

type TestController struct{}

func (test TestController) Test(c *gin.Context){
    var success  = map[string]string{}
    success["message"] = "this is from test"
    c.JSON(200 , success)
}
