package gincontextservices

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecieveURLMessage(context *gin.Context) {
	msgvalue := context.Param("message")
	context.JSON(http.StatusOK, gin.H{
		"message": "Message recieving from URL..." + msgvalue,
	})
}

func SendMsg(context *gin.Context) {

	msgvalue := context.Param("id")

	bookdata := Book{
		Id:   1001,
		Page: "20s",
		Name: "Golang",
	}

	jsonbookdata, error := json.Marshal(bookdata)

	if error != nil {
		fmt.Println("Error ovccured..", error)
		return
	}

	fmt.Println("Book id is", msgvalue, string(jsonbookdata))

	/* c.JSON(http.StatusFound, gin.H{
		"Id ":   ":" + Id,
		"Page":  ":" + Page,
		"Name ": ":" + Name,
	})*/
}

type Book struct {
	Id   int
	Page string
	Name string
}
