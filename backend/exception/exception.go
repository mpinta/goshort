package exception

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Http(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"Error": message})
}

func Internal(err error) {
	log.Println(err)
}

func FatalInternal(err error) {
	log.Fatalln(err)
}
