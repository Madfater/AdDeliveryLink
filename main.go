package main

import (
	"dcardAssignment/models"
	"dcardAssignment/routers"

	"github.com/gin-gonic/gin"
)

func main(){
	
	defer models.CloseDB()

	r:=gin.Default()

	routers.RouterInit(r)

	r.Run(":8080")
}