package main

import (
	"testEffective-Mobile/cmd"
	_ "testEffective-Mobile/docs"
)

// @title TestEffective-MobileAPI
// @version 1.0
// @description API server for testEffective-Mobile
// @host localhost:8080
// @BasePath /
func main() {
	cmd.Execute()
}
