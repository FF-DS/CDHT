package controllers

import (
    "github.com/gin-gonic/gin"
    // . "monitoring-server/util"
)


// type NodesController struct{}

type MonitoringController struct{}

func (monitor MonitoringController) Monitor(c *gin.Context){
    var success  = map[string]string{}
    success["message"] = "this is from ein monitoring"
    c.JSON(200 , success)
}
