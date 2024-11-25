package routes

import (
	"testEffective-Mobile/cmd/server"
	"testEffective-Mobile/internal/handler"
	"testEffective-Mobile/x/interfacesx"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RegistterRoutes(server server.GinServer, musicLibraryHandler *handler.MusicLibraryHandler) {
	server.RegisterGroupRoute("/path", []interfacesx.RouteDefinition{
		{},
	}, func(ctx *gin.Context) {
		logrus.Infof("Request on %s", ctx.Request.URL.Path)
	})
}
