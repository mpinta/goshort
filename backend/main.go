package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"goshort/backend/config"
	"goshort/backend/data"
	"goshort/backend/exception"
	"goshort/backend/handler"
)

func main() {
	cfg := config.GetConfig()

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.GET(cfg.Server.StatusEndpoint, handler.Status)
	r.POST(cfg.Server.ShortenEndpoint, handler.Shorten)
	r.NoRoute(handler.Find)

	err := data.Recreate()
	if err != nil {
		exception.FatalInternal(err)
	}

	err = r.Run(cfg.Server.Port)
	if err != nil {
		exception.FatalInternal(err)
	}
}
