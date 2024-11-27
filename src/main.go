package main

import (
	"github.com/Madfater/AdDeliveryLink/routers"
	"github.com/Madfater/AdDeliveryLink/utils"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

// @title Advertisement API
// @version 1.0
// @description a simple API for advertisement

// @BasePath /v1/api
func main() {
	gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")

	r := gin.New()
	r.Use(gin.RecoveryWithWriter(io.MultiWriter(f)))
	r.Use(gin.LoggerWithWriter(io.MultiWriter(f)))

	routers.RouterInit(r)

	err := r.Run(":8080")
	utils.HandleError(err, "Fail to run server")
}
