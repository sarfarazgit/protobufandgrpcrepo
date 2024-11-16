package gincontextservices

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecieveURLMessage(context *gin.Context) {
	msgvalue := context.Param("message")
	context.JSON(http.StatusOK, gin.H{
		"message": "Message recieving from URL..." + msgvalue,
	})
}
