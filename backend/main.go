package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mpinta/goshort/backend/config"
	"github.com/mpinta/goshort/backend/data"
	"github.com/mpinta/goshort/backend/exception"
	"github.com/mpinta/goshort/backend/handler"
)

func main() {
	cfg := config.GetConfig()

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.GET(cfg.Server.RootEndpoint + cfg.Server.StatusEndpoint, handler.Status)
	r.POST(cfg.Server.RootEndpoint + cfg.Server.ShortenEndpoint, handler.Shorten)
	r.GET(cfg.Server.RootEndpoint + cfg.Server.FindEndpoint + "/:url", handler.Find)

	err := data.Recreate()
	if err != nil {
		exception.FatalInternal(err)
	}

	err = r.Run(cfg.Server.Port)
	if err != nil {
		exception.FatalInternal(err)
	}
}
