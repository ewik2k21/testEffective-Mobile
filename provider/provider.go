package provider

import (
	"testEffective-Mobile/cmd/server"
	"testEffective-Mobile/internal/handler"
	"testEffective-Mobile/internal/repository"
	"testEffective-Mobile/internal/routes"
	"testEffective-Mobile/internal/services"

	"gorm.io/gorm"
)

func NewProvider(db *gorm.DB, server server.GinServer) {
	musicLibraryRepo := repository.NewMusicLibraryRepository(db)
	musicLibraryService := services.NewMusicLibraryService(musicLibraryRepo)
	musicLibraryHandler := handler.NewMusicLibraryHandler(musicLibraryService)
	routes.RegistterRoutes(server, musicLibraryHandler)

}
