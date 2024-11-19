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

/* Protobuf packages installation:

1. sudo yum install -y protobuf-compiler

2. [In VS CODE]

-> go get google.golang.org/protobuf@latest
-> go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

[Proto file compilation command]

$ protoc --go_out=. --go_opt=paths=source_relative  models/Person.proto

*/
