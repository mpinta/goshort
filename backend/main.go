package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"goshort/backend/config"
	"goshort/backend/data"
	"goshort/backend/handler"
)

func main() {
	c := config.GetConfig()

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.GET(c.Server.Status, handler.Status)
	r.POST(c.Server.Shorten, handler.Shorten)
	r.NoRoute(handler.Find)

	err := data.Recreate()
	if err != nil {
		panic(err)
	}

	err = r.Run(c.Server.Port)
	if err != nil {
		panic(err)
	}
}