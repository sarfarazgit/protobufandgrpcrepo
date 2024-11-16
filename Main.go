package main

import (
	"fmt"
	"protobufandgrpcrepo/gincontextservices"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("Testing..")

	ginEngine := gin.Default()
	ginEngine.GET("/datarecieve/:message", gincontextservices.RecieveURLMessage)
	//ginEngine.POST("/post/:id", gincontextservices.SendMsg)
	ginEngine.Run()

}
