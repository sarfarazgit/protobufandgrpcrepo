package main

import (
	"fmt"
	"protobufandgrpcrepo/gincontextservices"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("Testing..")

	ginEngine := gin.Default()
	ginEngine.GET("/datarecieve", gincontextservices.RecieveURLMessage)
	ginEngine.Run()

}
