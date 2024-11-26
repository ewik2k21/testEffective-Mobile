package routes

import (
	"testEffective-Mobile/cmd/server"
	"testEffective-Mobile/internal/handler"
	"testEffective-Mobile/x/interfacesx"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RegistterRoutes(server server.GinServer, musicLibraryHandler *handler.MusicLibraryHandler) {
	server.RegisterGroupRoute("/music_library", []interfacesx.RouteDefinition{
		{Method: "POST", Path: "/add_song", Handler: musicLibraryHandler.AddSong},
		{Method: "PATCH", Path: "/update_song/:songid", Handler: musicLibraryHandler.UpdateSong},
		{Method: "DELETE", Path: "/delete_song/:songid", Handler: musicLibraryHandler.DeleteSong},
		{Method: "GET", Path: "/get_all_data/", Handler: musicLibraryHandler.GetAllMusicLibraryData},
		{Method: "GET", Path: "/get_text_song/:songid", Handler: musicLibraryHandler.GetSongText},
	}, func(ctx *gin.Context) {
		logrus.Infof("Request on %s", ctx.Request.URL.Path)
	})
}
