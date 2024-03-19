package main

import (
	"dcardAssignment/src/models"
	"dcardAssignment/src/routers"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	defer models.CloseDB()

	gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")
	
	r := gin.New()
	r.Use(gin.RecoveryWithWriter(io.MultiWriter(f)))
	r.Use(gin.LoggerWithWriter(io.MultiWriter(f)))

	routers.RouterInit(r)

	r.Run(":8080")
}
