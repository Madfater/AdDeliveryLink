package main

import (
	"dcardAssignment/routers"

	"github.com/gin-gonic/gin"
)

func main(){
	
	r:=gin.Default()

	routers.RouterInit(r)

	r.Run(":8080")
}