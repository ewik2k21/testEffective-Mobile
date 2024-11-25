package cmd

import (
	"context"
	"os"
	"time"

	"testEffective-Mobile/cmd/server"
	"testEffective-Mobile/config"
	"testEffective-Mobile/provider"

	"github.com/sirupsen/logrus"
)

func Execute() {
	builder := server.NewGinServerBuilder()
	server := builder.Build()

	ctx := context.Background()
	config.LoadEnviroment()

	db, err := config.SetUpDatabase()
	if err != nil {
		logrus.Fatalf("Error setting up database %s", err)
	}

	provider.NewProvider(db, server)

	go func() {
		if err := server.Start(ctx, os.Getenv(config.AppPort)); err != nil {
			logrus.Errorf("Error starting server: %s", err)
		}
	}()

	<-ctx.Done()
	logrus.Info("Server stopped")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.ShutDown(shutdownCtx); err != nil {
		logrus.Errorf("Error shutting down server %s", err)
	}
}