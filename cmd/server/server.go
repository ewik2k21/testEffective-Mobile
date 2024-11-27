package server

import (
	"net/http"

	"testEffective-Mobile/x/interfacesx"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/net/context"
)

type GinServer interface {
	Start(ctx context.Context, httpAddress string) error
	ShutDown(ctx context.Context) error
	RegisterRoute(method, path string, handler gin.HandlerFunc)
	RegisterGroupRoute(path string, routes []interfacesx.RouteDefinition, middleWare ...gin.HandlerFunc)
}

type GinServerBuilder struct {
}

type ginServer struct {
	engine *gin.Engine
	server *http.Server
}

func NewGinServerBuilder() *GinServerBuilder {
	return &GinServerBuilder{}
}

func (b *GinServerBuilder) Build() GinServer {
	engine := gin.Default()
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return &ginServer{engine: engine}
}

func (gs *ginServer) Start(ctx context.Context, httpAddress string) error {
	gs.server = &http.Server{
		Addr:    httpAddress,
		Handler: gs.engine,
	}

	go func() {
		if err := gs.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Listening %s \n", err)
		}
	}()

	logrus.Infof("Http server is running on port %s", httpAddress)
	return nil
}

func (gs *ginServer) ShutDown(ctx context.Context) error {
	if err := gs.server.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server shutdown %s", err)
	}

	logrus.Info("Server is exiting")

	return nil
}

// for register single route
func (gs *ginServer) RegisterRoute(method, path string, handler gin.HandlerFunc) {
	switch method {
	case "GET":
		gs.engine.GET(path, handler)
	case "POST":
		gs.engine.POST(path, handler)
	case "PUT":
		gs.engine.PUT(path, handler)
	case "DELETE":
		gs.engine.DELETE(path, handler)
	case "PATCH":
		gs.engine.PATCH(path, handler)

	default:
		logrus.Errorf("Invalid https method")
	}
}

// for register groupe route
func (gs *ginServer) RegisterGroupRoute(path string, routes []interfacesx.RouteDefinition, middleWare ...gin.HandlerFunc) {
	group := gs.engine.Group(path)
	group.Use(middleWare...)
	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Path, route.Handler)
		case "POST":
			group.POST(route.Path, route.Handler)
		case "PUT":
			group.PUT(route.Path, route.Handler)
		case "DELETE":
			group.DELETE(route.Path, route.Handler)
		case "PATCH":
			group.PATCH(route.Path, route.Handler)

		default:
			logrus.Errorf("Invalid https method")
		}
	}
}
