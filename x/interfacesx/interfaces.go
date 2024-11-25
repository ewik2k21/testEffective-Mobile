package interfacesx

import "github.com/gin-gonic/gin"

type RouteDefinition struct {
	Path    string
	Method  string
	Handler gin.HandlerFunc
}
